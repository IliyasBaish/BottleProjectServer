package service

import (
	"example.com/server/pkg/repository"
	"example.com/server/structs"
)

type Authorization interface {
	CreateUser(user structs.User) (int, error)
	GetUser(username, password string) (structs.User, error)
	GetUserById(id int) (structs.User, error)
	AddCoinsToWallet(user structs.User, coins float32) error
	GetUsers() ([]structs.User, error)
	GetToken(username, password string) (string, error)
	VerifyDeal(station string, password string) (error, int)
}

type Deal interface {
	CreateDeal(user structs.User) (int, error)
	GetDealById(id int) (structs.Deal, error)
	GetDealsByUserId(id int) ([]structs.Deal, error)
	VerifyLastDeal(id int) error
	CreateUnverifiedDeal(deal structs.UnclaimedDeal, user structs.User) (int, error)
}

type Service struct {
	Authorization
	Deal
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Deal:          NewDealsService(repos.Deal),
	}
}
