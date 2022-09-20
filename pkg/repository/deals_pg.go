package repository

import (
	"fmt"
	"strings"

	"example.com/server/structs"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type DealsPG struct {
	db *sqlx.DB
}

func NewDealsPG(db *sqlx.DB) *DealsPG {
	return &DealsPG{db: db}
}

func (r *DealsPG) CreateDeal(user structs.User) (int, error) {
	var id int
	deals := []structs.UnclaimedDeal{}
	query := fmt.Sprintf("select * from unclaimed_deals where user_id = %d and verified = true", user.Id)
	rows, err := r.db.Query(query)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		deal := structs.UnclaimedDeal{}
		err = rows.Scan(&deal.Id, &deal.UserId, &deal.Coins, &deal.Date, pq.Array(&deal.Bottles), &deal.BottleCount, &deal.Station, &deal.Time, &deal.Verified)
		if err != nil {
			return -1, err
		}
		deals = append(deals, deal)
	}
	var verifiedDeal structs.Deal
	for _, d := range deals {
		verifiedDeal.Bottles = append(verifiedDeal.Bottles, d.Bottles...)
	}
	verifiedDeal.UserId = deals[0].UserId
	verifiedDeal.BottleCount = len(deals)
	verifiedDeal.Coins = getBotllesCost(verifiedDeal.Bottles, r)

	bottles := "'{" + "\"" + strings.Join(verifiedDeal.Bottles, "\", \"") + "\"" + "}'"
	query = fmt.Sprintf("INSERT INTO %s(user_id, coins, bottles, bottle_count) values (%d, %.2f, %s, %d) RETURNING id", dealsTable, verifiedDeal.UserId, verifiedDeal.Coins, bottles, verifiedDeal.BottleCount)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	query = fmt.Sprintf("delete from unclaimed_deals where user_id = %d", user.Id)
	rows, err = r.db.Query(query)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *DealsPG) CreateUnverifiedDeal(deal structs.UnclaimedDeal, user structs.User) (int, error) {
	var id int

	cost := getBotllesCost(deal.Bottles, r)
	bottles := "'{" + "\"" + strings.Join(deal.Bottles, "\", \"") + "\"" + "}'"

	query := fmt.Sprintf("INSERT INTO %s(user_id, coins, bottles, bottle_count, station) values (%d, %.2f, %s, %d, %d) RETURNING id", unverifiedDeals, user.Id, cost, bottles, deal.BottleCount, deal.Station)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *DealsPG) GetDealById(id int) (structs.Deal, error) {
	var deal structs.Deal

	query := fmt.Sprintf("select * from deals where id=%d", id)
	row := r.db.QueryRow(query)
	if err := row.Scan(&deal.Id, &deal.UserId, &deal.Coins, &deal.Date, &deal.Bottles, &deal.BottleCount); err != nil {
		return structs.Deal{}, err
	}

	return deal, nil
}

func getBotllesCost(barcodes []string, r *DealsPG) float32 {
	var total_cost float32 = 0.00
	var query string
	for _, barcode := range barcodes {
		var bottle_cost float32
		query = fmt.Sprintf("SELECT cost FROM bottles where barcode='%s'", barcode)
		row := r.db.QueryRow(query)
		if err := row.Scan(&bottle_cost); err == nil {
			total_cost += bottle_cost
		}
	}
	return total_cost
}

func (r *DealsPG) GetDealsByUserId(id int) ([]structs.Deal, error) {
	var deals []structs.Deal
	query := fmt.Sprintf("select * from deals where user_id = %d", id)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		deal := structs.Deal{}
		err = rows.Scan(&deal.Id, &deal.UserId, &deal.Coins, &deal.Date, &deal.Bottles, &deal.BottleCount)
		if err != nil {
			return nil, err
		}
		deals = append(deals, deal)
	}
	return deals, nil
}

func (r *DealsPG) VerifyLastDeal(station int) error {
	var deal structs.UnclaimedDeal
	query := fmt.Sprintf("select * from unclaimed_deals where station = %d and verified != true order by time limit 1", station)
	row := r.db.QueryRow(query)
	if err := row.Scan(&deal.Id, &deal.UserId, &deal.Coins, &deal.Date, pq.Array(&deal.Bottles), &deal.BottleCount, &deal.Station, &deal.Time, &deal.Verified); err != nil {
		return err
	}
	query = fmt.Sprintf("update unclaimed_deals set verified = true where id = %d and verified = false", deal.Id)
	row = r.db.QueryRow(query)
	if err := row.Err(); err != nil {
		return err
	}

	return nil
}
