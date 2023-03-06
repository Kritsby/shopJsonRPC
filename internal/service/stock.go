package service

import (
	"dev/lamoda_test/internal/model"
	"dev/lamoda_test/internal/repository"
)

type Stock struct {
	repo repository.Repository
}

func NewStock(repo repository.Repository) *Stock {
	return &Stock{
		repo: repo,
	}
}

func (s *Stock) Reserve(products []int) error {
	err := s.repo.Reserve(products)
	if err != nil {
		return err
	}

	return nil
}
func (s *Stock) ReserveRelease(products []int) error {
	err := s.repo.ReserveRelease(products)
	if err != nil {
		return err
	}
	return nil
}

func (s *Stock) GetAmount(stock int) ([]model.Products, error) {
	result, err := s.repo.GetAmount(stock)
	if err != nil {
		return nil, err
	}
	return result, nil
}
