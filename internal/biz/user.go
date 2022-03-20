package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type ProfileRepo interface {
	GetProfile(ctx context.Context, username string) (*Profile, error)
	FollowUser(ctx context.Context, username string) (*Profile, error)
	UnfollowUser(ctx context.Context, username string) (*Profile, error)
}

type UserUsecase struct {
	ur   UserRepo
	pr   ProfileRepo
	jwtc *conf.JWT
	log  *log.Helper
}

type Profile struct {
	Username  string
	Bio       string
	Image     string
	Following bool
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func verifyPassword(hashed, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	if err != nil {
		return false
	}
	return true
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, jwtc *conf.JWT, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, jwtc: jwtc, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GenerateToken(username string) string {
	return auth.GenerateToken(uc.jwtc.Token, username)
}

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    uc.GenerateToken(username),
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(hashPassword(password), u.PasswordHash) {
		return nil, errors.New("Login Error")
	}

	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Token:    "xxx",
		Bio:      u.Bio,
		Image:    u.Image,
	}, nil
}

func (uc *UserUsecase) GetCurrentUser(ctx context.Context) (*User, error) {
	return nil, nil
}
