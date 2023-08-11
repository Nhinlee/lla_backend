-- name: CreateLearningItem :one
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
) RETURNING *;

-- name: GetLearningItem :many
SELECT * FROM learning_item;

-- name: DeleteLearningItem :one
DELETE FROM learning_item WHERE id = $1 RETURNING *;
