package service

import (
	"dev/lamoda_test/internal/entity"
	"dev/lamoda_test/internal/repository/postgresql"
)

type StockService struct {
	repo postgresql.PostgresStocker
}

func NewStockService(repo postgresql.PostgresStocker) *StockService {
	return &StockService{
		repo: repo,
	}
}

func (s *StockService) Reserve(products []int) error {
	err := s.repo.Reserve(products)
	if err != nil {
		return err
	}

	return nil
}
func (s *StockService) ReserveRelease(products []int) error {
	err := s.repo.ReserveRelease(products)
	if err != nil {
		return err
	}
	return nil
}

func (s *StockService) GetAmount(stock int) ([]entity.Products, error) {
	result, err := s.repo.GetAmount(stock)
	if err != nil {
		return nil, err
	}
	return result, nil
}
