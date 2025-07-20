package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Invoice struct {
	Id            int64     `json:"id"`
	CustomerName  string    `json:"customer_name"`
	CustomerPhone string    `json:"customer_phone"`
	TotalAmount   float64   `json:"total_amount"`
	Discount      float64   `json:"discount"`
	FinalAmount   float64   `json:"final_amount"`
	CreatedOn     time.Time `json:"created_on"`
	UpdatedOn     time.Time `json:"updated_on"`
}

type CreateInvoice struct {
	CustomerName  string              `json:"customer_name"`
	CustomerPhone string              `json:"customer_phone"`
	Discount      float64             `json:"discount"`
	Items         []CreateInvoiceItem `json:"items"`
}

func (data *CreateInvoice) Validate() error {
	if data.CustomerName == "" {
		return errors.New("customer_name is required")
	}
	if data.CustomerPhone == "" {
		return errors.New("customer_phone is required")
	}
	if len(data.Items) == 0 {
		return errors.New("items are required")
	}
	return nil
}

type InvoiceModel struct {
	Db *pgxpool.Pool
}

func (model *InvoiceModel) Get() ([]Invoice, error) {
	sql := "SELECT * FROM invoices"

	rows, err := model.Db.Query(context.Background(), sql)
	if err != nil {
		return []Invoice{}, err
	}

	invoices := []Invoice{}
	for rows.Next() {
		invoice, err := pgx.RowToStructByName[Invoice](rows)
		if err != nil {
			return invoices, err
		}

		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (model *InvoiceModel) GetById(id int64) (Invoice, error) {
	sql := "SELECT id, customer_name, customer_phone, total_amount, discount, final_amount, created_on, updated_on FROM invoices WHERE id=$1"
	args := []any{id}

	var invoice Invoice
	if err := model.Db.QueryRow(context.Background(), sql, args...).Scan(
		&invoice.Id,
		&invoice.CustomerName,
		&invoice.CustomerPhone,
		&invoice.TotalAmount,
		&invoice.Discount,
		&invoice.FinalAmount,
		&invoice.CreatedOn,
		&invoice.UpdatedOn,
	); err != nil {
		return invoice, err
	}

	return invoice, nil
}

func (model *InvoiceModel) Create(data CreateInvoice) (int64, error) {
	if err := data.Validate(); err != nil {
		return -1, err
	}

	// Creating the invoice
	totalAmount := 0.0
	for _, item := range data.Items {
		totalAmount += item.Price * float64(item.Qty)
	}

	finalAmount := totalAmount - data.Discount

	tx, err := model.Db.Begin(context.Background())
	if err != nil {
		return -1, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	sql := "INSERT INTO invoices(customer_name, customer_phone, total_amount, discount, final_amount) VALUES($1, $2, $3, $4, $5) RETURNING id"
	args := []any{data.CustomerName, data.CustomerPhone, totalAmount, data.Discount, finalAmount}

	var invoiceId int64
	if err = tx.QueryRow(context.Background(), sql, args...).Scan(&invoiceId); err != nil {
		return -1, err
	}

	// Saving the items
	for _, item := range data.Items {
		item.InvoiceId = invoiceId
		if err = item.Validate(); err != nil {
			return invoiceId, err
		}

		sql := "INSERT INTO invoice_items(invoice_id, product_id, title, price, qty) VALUES($1, $2, $3, $4, $5)"
		args := []any{item.InvoiceId, item.ProductId, item.Title, item.Price, item.Qty}

		_, err = tx.Exec(context.Background(), sql, args...)
		if err != nil {
			return invoiceId, err
		}

		// Updating stock for the product
		sql = "UPDATE products SET stock=stock-$1 WHERE id=$2"
		args = []any{item.Qty, item.ProductId}

		_, err = tx.Exec(context.Background(), sql, args...)
		if err != nil {
			return invoiceId, err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return -1, err
	}

	return invoiceId, nil
}
