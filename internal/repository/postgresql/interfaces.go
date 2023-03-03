package postgresql

import (
	"dev/lamoda_test/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresStocker interface {
	Reserve(products []int) error
	ReserveRelease(products []int) error
	GetAmount(stock int) ([]entity.Products, error)
}

type Tester interface {
}

type Repository struct {
	PostgresStocker
	Tester
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		PostgresStocker: NewStockPostgres(db)}
}
