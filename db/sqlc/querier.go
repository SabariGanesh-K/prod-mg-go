// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	AddCompressedProductImageUrlsByID(ctx context.Context, arg AddCompressedProductImageUrlsByIDParams) (Products, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Products, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	GetProductByProductID(ctx context.Context, id string) (Products, error)
	GetProductsByUserID(ctx context.Context, arg GetProductsByUserIDParams) ([]Products, error)
	GetUserByID(ctx context.Context, userID string) (Users, error)
	//   is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
}

var _ Querier = (*Queries)(nil)
