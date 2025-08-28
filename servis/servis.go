package servis

import (
	"errors"
	"fmt"
	_const "game/const"
	"game/dto"
	"game/entity"
	"game/pkg/hashPassword"
	"game/pkg/richerror"
	"game/repository/mysql"
)

type Repository interface {
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

func New(authgenerator AuthGenerator, repo *mysql.MySqlDb) Service {
	return Service{auth: authgenerator, repo: repo}
}

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	//TODO- we should verify phone number by verification code

	ps, err := hashPassword.HashPassword(req.Password)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("i can't hashed password: %w", err)
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    ps,
	}

	createdUser, rErr := s.repo.Register(user)
	if rErr != nil {

		return dto.RegisterResponse{}, fmt.Errorf("unxeopted error: %w", rErr)
	}

	resp := dto.RegisterResponse{struct {
		ID          uint   `json:"id"`
		PhoneNumber string `json:"phone_number"`
		Name        string `json:"name"`
	}{ID: createdUser.ID,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name}}

	return resp, nil
}

// TODO - please implement me

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	User  dto.UserInfo  `json:"user"`
	Token TokenResponse `json:"token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	const op = "servis.login"
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, richerror.
			New(op).
			WithError(err).
			WithMeta(map[string]interface{}{"phone req ": req.PhoneNumber})
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
	return LoginResponse{
		User: dto.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Token: TokenResponse{
			AccessToken:  accesstoken,
			RefreshToken: refreshToken,
		},
	}, nil
}

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}

type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	const op = "servis.Profile"

	user, err := s.repo.GetUserByID(req.UserID)
	//log.Fatal("req :", req.UserID)
	if err != nil {

		return ProfileResponse{},
			richerror.New(op).
				WithError(err).
				WithMeta(map[string]interface{}{"req": req})
	}
	fmt.Println(user.Name)

	fmt.Println(user.Name)

	return ProfileResponse{Name: user.Name}, nil
}
