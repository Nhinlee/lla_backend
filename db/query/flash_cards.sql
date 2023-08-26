-- name: GetLearningItemsByTopicAndCompleted :many
SELECT * FROM learning_items 
WHERE topic_id = $1 AND deleted_at IS NULL
ORDER BY completed_at ASC
LIMIT $2;

-- name: UpdateCompletedAt :batchexec
UPDATE learning_items
SET completed_at = NOW()
WHERE id = $1;