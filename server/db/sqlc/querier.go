// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	GetUser(ctx context.Context, id int64) (int64, error)
}

var _ Querier = (*Queries)(nil)
