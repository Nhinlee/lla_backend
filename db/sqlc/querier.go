// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	CreateLearningItem(ctx context.Context, arg CreateLearningItemParams) (LearningItem, error)
}

var _ Querier = (*Queries)(nil)
