package repository

import (
	"dev/lamoda_test/internal/model"
)

type Repository interface {
	Reserve(products []int) error
	ReserveRelease(products []int) error
	GetAmount(stock int) ([]model.Products, error)
}
