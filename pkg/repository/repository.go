package repository

import (
	"example.com/server/structs"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user structs.User) (int, error)
	GetUser(username string, password string) (structs.User, error)
	GetUserById(id int) (structs.User, error)
	AddCoinsToWallet(user structs.User, coins float32) error
	GetUsers() ([]structs.User, error)
	VerifyStation(stationName string, password string) (error, int)
}

type Deal interface {
	CreateDeal(user structs.User) (int, error)
	GetDealById(id int) (structs.Deal, error)
	GetDealsByUserId(id int) ([]structs.Deal, error)
	VerifyLastDeal(station int) error
	CreateUnverifiedDeal(deal structs.UnclaimedDeal, user structs.User) (int, error)
}

type Repository struct {
	Authorization
	Deal
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgre(db),
		Deal:          NewDealsPG(db),
	}
}
