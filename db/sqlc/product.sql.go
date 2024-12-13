// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: product.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const addCompressedProductImageUrlsByID = `-- name: AddCompressedProductImageUrlsByID :one
UPDATE products
SET
compressed_product_images_urls = $1 
WHERE id = $2 
RETURNING id, user_id, product_name, product_description, product_price, product_urls, compressed_product_images_urls, created_at
`

type AddCompressedProductImageUrlsByIDParams struct {
	CompressedProductImagesUrls []string `json:"compressed_product_images_urls"`
	ID                          string   `json:"id"`
}

func (q *Queries) AddCompressedProductImageUrlsByID(ctx context.Context, arg AddCompressedProductImageUrlsByIDParams) (Products, error) {
	row := q.db.QueryRowContext(ctx, addCompressedProductImageUrlsByID, pq.Array(arg.CompressedProductImagesUrls), arg.ID)
	var i Products
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductName,
		&i.ProductDescription,
		&i.ProductPrice,
		pq.Array(&i.ProductUrls),
		pq.Array(&i.CompressedProductImagesUrls),
		&i.CreatedAt,
	)
	return i, err
}

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
    id,
    user_id,
    product_name,
    product_description,
    product_price,
    product_urls,
    compressed_product_images_urls
) VALUES (
    $1,$2,$3,$4,$5,$6,'{}'
) RETURNING id, user_id, product_name, product_description, product_price, product_urls, compressed_product_images_urls, created_at
`

type CreateProductParams struct {
	ID                 string   `json:"id"`
	UserID             string   `json:"user_id"`
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductPrice       string   `json:"product_price"`
	ProductUrls        []string `json:"product_urls"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Products, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ID,
		arg.UserID,
		arg.ProductName,
		arg.ProductDescription,
		arg.ProductPrice,
		pq.Array(arg.ProductUrls),
	)
	var i Products
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductName,
		&i.ProductDescription,
		&i.ProductPrice,
		pq.Array(&i.ProductUrls),
		pq.Array(&i.CompressedProductImagesUrls),
		&i.CreatedAt,
	)
	return i, err
}

const getProductByProductID = `-- name: GetProductByProductID :one
SELECT id, user_id, product_name, product_description, product_price, product_urls, compressed_product_images_urls, created_at FROM products
WHERE id = $1
`

func (q *Queries) GetProductByProductID(ctx context.Context, id string) (Products, error) {
	row := q.db.QueryRowContext(ctx, getProductByProductID, id)
	var i Products
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductName,
		&i.ProductDescription,
		&i.ProductPrice,
		pq.Array(&i.ProductUrls),
		pq.Array(&i.CompressedProductImagesUrls),
		&i.CreatedAt,
	)
	return i, err
}

const getProductsByUserID = `-- name: GetProductsByUserID :many
SELECT id, user_id, product_name, product_description, product_price, product_urls, compressed_product_images_urls, created_at FROM products
WHERE user_id = $1
  AND product_price >= COALESCE($2, product_price)
  AND product_price <= COALESCE($3, product_price)
  AND product_name ILIKE '%' || COALESCE($4, product_name) || '%'
`

type GetProductsByUserIDParams struct {
	UserID      string         `json:"user_id"`
	MinPrice    sql.NullString `json:"min_price"`
	MaxPrice    sql.NullString `json:"max_price"`
	ProductName sql.NullString `json:"product_name"`
}

func (q *Queries) GetProductsByUserID(ctx context.Context, arg GetProductsByUserIDParams) ([]Products, error) {
	rows, err := q.db.QueryContext(ctx, getProductsByUserID,
		arg.UserID,
		arg.MinPrice,
		arg.MaxPrice,
		arg.ProductName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Products
	for rows.Next() {
		var i Products
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductName,
			&i.ProductDescription,
			&i.ProductPrice,
			pq.Array(&i.ProductUrls),
			pq.Array(&i.CompressedProductImagesUrls),
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
