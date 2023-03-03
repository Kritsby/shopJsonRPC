package postgresql

import (
	"context"
	"dev/lamoda_test/internal/entity"
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

	for rows.Next() {
		var result entity.Products

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

	for rows.Next() {
		var result entity.Products

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

	return nil
}

func (s *StockPsql) GetAmount(stockId int) ([]entity.Products, error) {
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

	var result []entity.Products
	for rows.Next() {
		var r entity.Products

		err = rows.Scan(&r.Storage, &r.Product, &r.Amount)
		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}
