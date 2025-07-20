package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type CreateProduct struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func (data *CreateProduct) Validate() error {
	return nil
}

type ProductModel struct {
	Db *pgxpool.Pool
}

func (model *ProductModel) Get() ([]Product, error) {
	sql := "SELECT * FROM products"

	rows, err := model.Db.Query(context.Background(), sql)
	if err != nil {
		return []Product{}, err
	}

	products := []Product{}
	for rows.Next() {
		product, err := pgx.RowToStructByName[Product](rows)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (model *ProductModel) Create(data CreateProduct) (int64, error) {
	if err := data.Validate(); err != nil {
		return -1, err
	}

	sql := "INSERT INTO products(title, price, stock) VALUES($1, $2, $3) RETURNING id"
	args := []any{data.Title, data.Price, data.Stock}

	var insertId int64
	if err := model.Db.QueryRow(context.Background(), sql, args...).Scan(&insertId); err != nil {
		return -1, err
	}

	return insertId, nil
}
