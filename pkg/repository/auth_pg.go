package repository

import (
	"errors"
	"fmt"

	"example.com/server/structs"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgre(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user structs.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ('%s', '%s') RETURNING id", usersTable, user.Username, user.Password)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username string, password string) (structs.User, error) {
	var user structs.User
	query := fmt.Sprintf("select * from %s where username='%s' AND password_hash='%s';", usersTable, username, password)
	row := r.db.QueryRow(query)
	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Wallet, &user.Role); err != nil {
		return structs.User{}, nil
	}

	return user, nil
}

func (r *AuthPostgres) GetUserById(id int) (structs.User, error) {
	var user structs.User
	query := fmt.Sprintf("select * from %s where id=%d", usersTable, id)
	row := r.db.QueryRow(query)
	if err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Wallet, &user.Role); err != nil {
		return structs.User{}, nil
	}
	//logrus.Fatalf(user.Username)

	return user, nil
}

func (r *AuthPostgres) AddCoinsToWallet(user structs.User, coins float32) error {
	query := fmt.Sprintf("update %s set wallet=%.2f where id=%d;", usersTable, user.Wallet+coins, user.Id)
	row := r.db.QueryRow(query)
	if err := row.Err(); err != nil {
		return err
	}

	return nil
}

func (r *AuthPostgres) GetUsers() ([]structs.User, error) {
	users := []structs.User{}
	query := fmt.Sprintf("select id, username, wallet, role user from %s", usersTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := structs.User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Wallet, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *AuthPostgres) VerifyStation(stationName string, password string) (error, int) {
	var station structs.User
	query := fmt.Sprintf("select * from %s where username='%s' AND password_hash='%s';", usersTable, stationName, password)
	row := r.db.QueryRow(query)
	if err := row.Scan(&station.Id, &station.Username, &station.Password, &station.Wallet, &station.Role); err != nil {
		return err, -1
	}

	if station.Role != "station" {
		return errors.New("not station"), -1
	}

	return nil, station.Id
}
