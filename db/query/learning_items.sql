-- name: CreateLearningItem :one
INSERT INTO learning_items (
    id,
    image_link,
    english_word,
    vietnamese_word,
    english_sentences,
    created_at,
    updated_at,
    user_id,
    topic_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    now(),
    now(),
    $6,
    $7
) RETURNING *;

-- name: GetLearningItemById :one
SELECT * FROM learning_items 
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetAllLearningItems :many
SELECT * FROM learning_items 
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: DeleteLearningItem :one
UPDATE learning_items SET deleted_at = now() WHERE id = $1 RETURNING *;

-- name: HardDeleteLearningItem :one
DELETE FROM learning_items WHERE id = $1 RETURNING *;

-- name: UpdateLearningItem :exec
UPDATE learning_items SET
    image_link = $2,
    english_word = $3,
    vietnamese_word = $4,
    english_sentences = $5,
    updated_at = now(),
    completed_at = $6,
    topic_id = $7
WHERE id = $1;