package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InvoiceItem struct {
	Id        int64     `json:"id"`
	InvoiceId int64     `json:"invoice_id"`
	ProductId int64     `json:"product_id"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	Qty       int       `json:"qty"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type CreateInvoiceItem struct {
	InvoiceId int64   `json:"invoice_id"`
	ProductId int64   `json:"product_id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Qty       int     `json:"qty"`
}

func (data *CreateInvoiceItem) Validate() error {
	if data.InvoiceId <= 0 {
		return errors.New("invoice_id is required")
	}
	if data.ProductId <= 0 {
		return errors.New("product_id is required")
	}
	if data.Title == "" {
		return errors.New("title is required")
	}
	if data.Price == 0.0 {
		return errors.New("price is required")
	}
	if data.Qty < 1 {
		return errors.New("qty is required")
	}
	return nil
}

type InvoiceItemModel struct {
	Db *pgxpool.Pool
}

func (model *InvoiceItemModel) Get() ([]InvoiceItem, error) {
	sql := "SELECT * FROM invoice_items"

	rows, err := model.Db.Query(context.Background(), sql)
	if err != nil {
		return []InvoiceItem{}, err
	}

	items := []InvoiceItem{}
	for rows.Next() {
		item, err := pgx.RowToStructByName[InvoiceItem](rows)
		if err != nil {
			return items, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (model *InvoiceItemModel) Create(data CreateInvoiceItem) (int64, error) {
	if err := data.Validate(); err != nil {
		return -1, err
	}

	sql := "INSERT INTO invoice_items(invoice_id, title, price, qty) VALUES($1, $2, $3, $4) RETURNING id"
	args := []any{data.InvoiceId, data.Title, data.Price, data.Qty}

	var insertId int64
	if err := model.Db.QueryRow(context.Background(), sql, args...).Scan(&insertId); err != nil {
		return -1, err
	}

	return insertId, nil
}
