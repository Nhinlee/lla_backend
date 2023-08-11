// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: learning_item.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const createLearningItem = `-- name: CreateLearningItem :one
INSERT INTO learning_item (
    id,
    image_link,
    english_word,
    vietnamese_word,
    english_sentences
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING id, image_link, english_word, vietnamese_word, english_sentences
`

type CreateLearningItemParams struct {
	ID               string   `json:"id"`
	ImageLink        string   `json:"image_link"`
	EnglishWord      string   `json:"english_word"`
	VietnameseWord   string   `json:"vietnamese_word"`
	EnglishSentences []string `json:"english_sentences"`
}

func (q *Queries) CreateLearningItem(ctx context.Context, arg CreateLearningItemParams) (LearningItem, error) {
	row := q.db.QueryRowContext(ctx, createLearningItem,
		arg.ID,
		arg.ImageLink,
		arg.EnglishWord,
		arg.VietnameseWord,
		pq.Array(arg.EnglishSentences),
	)
	var i LearningItem
	err := row.Scan(
		&i.ID,
		&i.ImageLink,
		&i.EnglishWord,
		&i.VietnameseWord,
		pq.Array(&i.EnglishSentences),
	)
	return i, err
}