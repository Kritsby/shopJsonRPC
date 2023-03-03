package service

import (
	"dev/lamoda_test/internal/entity"
	"dev/lamoda_test/internal/repository/postgresql"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type Stocker interface {
	Reserve(products []int) error
	ReserveRelease(products []int) error
	GetAmount(stock int) ([]entity.Products, error)
}

type Service struct {
	Stocker
}

func NewService(repo *postgresql.Repository) *Service {
	return &Service{
		Stocker: NewStockService(repo.PostgresStocker),
	}
}
