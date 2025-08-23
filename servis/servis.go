package servis

import (
	"errors"
	"fmt"
	_const "game/const"
	"game/entity"
	"game/pkg/hashPassword"
	"game/pkg/phonenumber"
	"game/repository/mysql"
)

type Repository interface {
	IsUniquePhoneNumber(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User, subject string) (string, error)
	CreateRefreshToken(user entity.User, subject string) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

type RegisterUser struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User struct {
		ID          uint   `json:"id"`
		PhoneNumber string `json:"phone_number"`
		Name        string `json:"name"`
	} `json:"user"`
}

func New(authgenerator AuthGenerator, repo *mysql.MySqlDb) Service {
	return Service{auth: authgenerator, repo: repo}
}

func (s Service) Register(req RegisterUser) (RegisterResponse, error) {
	//TODO- we should verify phone number by verification code

	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, errors.New("invalid phone number")
	}
	if isUniq, err := s.repo.IsUniquePhoneNumber(req.PhoneNumber); err != nil || !isUniq {
		if err != nil {
			return RegisterResponse{}, err
		}

		if !isUniq {
			return RegisterResponse{}, errors.New("phone number is not unique")
		}
	}
	//TODO - The password must be strong.
	if len(req.Password) < 7 {
		return RegisterResponse{}, errors.New("password must be at least 7 characters")
	}

	hashPass, hErr := hashPassword.HashPassword(req.Password)
	if hErr != nil {
		return RegisterResponse{}, hErr
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    hashPass,
	}

	createdUser, rErr := s.repo.Register(user)
	if rErr != nil {
		return RegisterResponse{}, fmt.Errorf("unxeopted error: %w", rErr)
	}
	//var resp RegisterResponse
	//resp.User.ID = createdUser.ID
	//resp.User.PhoneNumber = createdUser.PhoneNumber
	//resp.User.Name = createdUser.Name
	//return resp, nil

	resp := RegisterResponse{struct {
		ID          uint   `json:"id"`
		PhoneNumber string `json:"phone_number"`
		Name        string `json:"name"`
	}{ID: createdUser.ID,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name}}
	return resp, nil
}

// TODO - please implement me
type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {

	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, err
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password is not correct")
	}

	isvalid := hashPassword.VerifyPassword(req.Password, user.Password)

	if !isvalid {
		return LoginResponse{}, errors.New("username or password is not correct")
	}

	//  TODO - you should work this function

	accesstoken, aErr := s.auth.CreateAccessToken(user, _const.AccessTokenSubject)

	if aErr != nil {
		return LoginResponse{}, aErr
	}

	refreshToken, rErr := s.auth.CreateRefreshToken(user, _const.RefreshTokenSubject)

	if rErr != nil {
		return LoginResponse{}, fmt.Errorf("enxepted error %w", rErr)

	}
	return LoginResponse{AccessToken: accesstoken, RefreshToken: refreshToken}, err
}

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}

type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	user, err := s.repo.GetUserByID(req.UserID)
	//log.Fatal("req :", req.UserID)
	if err != nil {

		return ProfileResponse{}, fmt.Errorf("unexepted error %w", err)
	}
	fmt.Println(user.Name)

	fmt.Println(user.Name)
	return ProfileResponse{Name: user.Name}, nil
}
