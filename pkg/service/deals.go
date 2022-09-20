package service

import (
	"example.com/server/pkg/repository"
	"example.com/server/structs"
)

type DealsService struct {
	repo repository.Deal
}

func NewDealsService(repo repository.Deal) *DealsService {
	return &DealsService{repo: repo}
}

func (s *DealsService) CreateDeal(user structs.User) (int, error) {
	return s.repo.CreateDeal(user)
}

func (s *DealsService) CreateUnverifiedDeal(deal structs.UnclaimedDeal, user structs.User) (int, error) {
	return s.repo.CreateUnverifiedDeal(deal, user)
}

func (s *DealsService) GetDealById(id int) (structs.Deal, error) {
	return s.repo.GetDealById(id)
}

func (s *DealsService) GetDealsByUserId(id int) ([]structs.Deal, error) {
	return s.repo.GetDealsByUserId(id)
}

func (s *DealsService) VerifyLastDeal(id int) error {
	return s.repo.VerifyLastDeal(id)
}
