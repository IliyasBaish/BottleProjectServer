package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"example.com/server/pkg/jwt_auth"
	"example.com/server/pkg/repository"
	"example.com/server/structs"
	jwt "github.com/dgrijalva/jwt-go/v4"
)

const salt = "fsdpfnbairng934wt2=tmvarg"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user structs.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(username, password string) (structs.User, error) {
	password = generatePasswordHash(password)
	return s.repo.GetUser(username, password)
}

func (s *AuthService) GetUserById(id int) (structs.User, error) {
	return s.repo.GetUserById(id)
}

func (s *AuthService) AddCoinsToWallet(user structs.User, coins float32) error {
	err := s.repo.AddCoinsToWallet(user, coins)
	return err
}

func (s *AuthService) GetUsers() ([]structs.User, error) {
	return s.repo.GetUsers()
}

func (s *AuthService) GetToken(username string, password string) (string, error) {
	pwd := generatePasswordHash(password)
	user, err := s.repo.GetUser(username, pwd)
	if err != nil || user.Username == "" {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt_auth.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 30)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: user.Username,
	})
	return token.SignedString("secret")
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) VerifyDeal(station string, password string) (error, int) {
	password = generatePasswordHash(password)
	err, stationId := s.repo.VerifyStation(station, password)
	if err != nil {
		return err, -1
	}

	return nil, stationId
}
