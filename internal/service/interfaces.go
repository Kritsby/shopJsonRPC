package service

import (
	"dev/lamoda_test/internal/model"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type Stocker interface {
	Reserve(products []int) error
	ReserveRelease(products []int) error
	GetAmount(stock int) ([]model.Products, error)
}
