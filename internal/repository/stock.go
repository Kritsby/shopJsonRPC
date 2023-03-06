package repository

import (
	"context"
	"dev/lamoda_test/internal/model"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"strconv"
)

type StockPsql struct {
	db *pgxpool.Pool
}

func NewStockPostgres(db *pgxpool.Pool) *StockPsql {
	return &StockPsql{db: db}
}

func (s *StockPsql) Reserve(products []int) error {
	updateQuery := `
UPDATE
    shop.product_amount
SET
    amount = amount - 1
WHERE
    product_id = ANY(ARRAY[$1::INTEGER[]])
  	AND amount > 0
RETURNING storage_id, product_id, amount;`

	rows, err := s.db.Query(context.Background(), updateQuery, products)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var result model.Products

		err = rows.Scan(&result.Storage, &result.Product, &result.Amount)
		if err != nil {
			return err
		}

		storage := strconv.Itoa(result.Storage)
		product := strconv.Itoa(result.Product)
		amount := strconv.Itoa(result.Amount)

		log.Info().
			Str("Storage", storage).
			Str("Product", product).
			Str("Amount", amount).
			Msg("Changes")
	}

	if err = rows.Err(); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
	}

	return nil
}

func (s *StockPsql) ReserveRelease(products []int) error {
	updateQuery := `
UPDATE
    shop.product_amount
SET
    amount = amount + 1
WHERE
    product_id = ANY(ARRAY[$1::INTEGER[]])
  	AND amount > 0
RETURNING storage_id, product_id, amount;`

	rows, err := s.db.Query(context.Background(), updateQuery, products)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var result model.Products

		err = rows.Scan(&result.Storage, &result.Product, &result.Amount)
		if err != nil {
			return err
		}

		storage := strconv.Itoa(result.Storage)
		product := strconv.Itoa(result.Product)
		amount := strconv.Itoa(result.Amount)

		log.Info().
			Str("Storage", storage).
			Str("Product", product).
			Str("Amount", amount).
			Msg("Changes")
	}

	if err = rows.Err(); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
	}

	return nil
}

func (s *StockPsql) GetAmount(stockId int) ([]model.Products, error) {
	query := `
SELECT
    storage_id,
    product_id,
	amount
FROM
	shop.product_amount WHERE storage_id = $1`

	rows, err := s.db.Query(context.Background(), query, stockId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Products
	for rows.Next() {
		var r model.Products

		err = rows.Scan(&r.Storage, &r.Product, &r.Amount)
		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	if err = rows.Err(); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
	}

	return result, nil
}
