package usecase

import (
	"fmt"

	pb "github.com/loak155/microservices/proto/user"
	"golang.org/x/crypto/bcrypt"

	"github.com/loak155/microservices/services/auth/client"
	"github.com/loak155/microservices/services/auth/utils"
)

type IAuthUsecase interface {
	Signup(username string, email string, password string) (int, error)
	Signin(email string, password string) (string, error)
	GenerateToken(user_id int) (string, error)
	ValidateToken(token string) (bool, error)
	RefreshToken(token string) (string, error)
}

type authUsecase struct {
	uc         client.IUserGRPCClient
	jwtManager *utils.JwtManager
}

func NewAuthUsecase(uc client.IUserGRPCClient, jwtManager utils.JwtManager) IAuthUsecase {
	return &authUsecase{uc, &jwtManager}
}

func (uu *authUsecase) Signup(username string, email string, password string) (int, error) {
	req := pb.CreateUserRequest{
		User: &pb.User{
			Username: username,
			Email:    email,
			Password: password,
		},
	}
	res, err := uu.uc.CreateUser(&req)
	if err != nil {
		return 0, err
	}
	return int(res.User.Id), nil
}

func (uu *authUsecase) Signin(email string, password string) (string, error) {
	req := pb.GetUserByEmailRequest{Email: email}
	res, err := uu.uc.GetUserByEmail(&req)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.User.Password), []byte(password))
	if err != nil {
		return "", err
	}
	token, err := uu.jwtManager.Generate(int(res.User.Id))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uu *authUsecase) GenerateToken(user_id int) (string, error) {
	req := pb.GetUserRequest{Id: int32(user_id)}
	res, err := uu.uc.GetUser(&req)
	if err != nil {
		return "", err
	}
	// TODO: resのパスワードが一致するか確認する
	fmt.Println(res)
	token, err := uu.jwtManager.Generate(user_id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uu *authUsecase) ValidateToken(token string) (bool, error) {
	_, err := uu.jwtManager.ValidateToken(token)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (uu *authUsecase) RefreshToken(token string) (string, error) {
	claims, err := uu.jwtManager.ValidateToken(token)
	if err != nil {
		return "", err
	}
	refreshToken, err := uu.jwtManager.Generate(claims.UserId)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
