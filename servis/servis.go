package servis

import (
	"errors"
	"fmt"
	_const "game/const"
	"game/entity"
	"game/param"
	"game/pkg/hashPassword"
	"game/pkg/richerror"
	"game/repository/mysql"
)

type Repository interface {
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
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

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {
	//TODO- we should verify phone number by verification code

	ps, err := hashPassword.HashPassword(req.Password)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("i can't hashed password: %w", err)
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    ps,
	}

	createdUser, rErr := s.repo.Register(user)
	if rErr != nil {

		return param.RegisterResponse{}, fmt.Errorf("unxeopted error: %w", rErr)
	}

	resp := param.RegisterResponse{struct {
		ID          uint   `json:"id"`
		PhoneNumber string `json:"phone_number"`
		Name        string `json:"name"`
	}{ID: createdUser.ID,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name}}

	return resp, nil
}

// TODO - please implement me

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	const op = "servis.login"
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, richerror.
			New(op).
			WithError(err).
			WithMeta(map[string]interface{}{"phone req ": req.PhoneNumber})
	}

	isvalid := hashPassword.VerifyPassword(req.Password, user.Password)

	if !isvalid {
		return param.LoginResponse{}, errors.New("username or password is not correct")
	}

	//  TODO - you should work this function

	accesstoken, aErr := s.auth.CreateAccessToken(user, _const.AccessTokenSubject)

	if aErr != nil {
		return param.LoginResponse{}, aErr
	}

	refreshToken, rErr := s.auth.CreateRefreshToken(user, _const.RefreshTokenSubject)

	if rErr != nil {
		return param.LoginResponse{}, fmt.Errorf("enxepted error %w", rErr)

	}
	return param.LoginResponse{
		User: param.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Token: param.TokenResponse{
			AccessToken:  accesstoken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "servis.Profile"

	user, err := s.repo.GetUserByID(req.UserID)
	//log.Fatal("req :", req.UserID)
	if err != nil {

		return param.ProfileResponse{},
			richerror.New(op).
				WithError(err).
				WithMeta(map[string]interface{}{"req": req})
	}
	fmt.Println(user.Name)

	fmt.Println(user.Name)

	return param.ProfileResponse{Name: user.Name}, nil
}
