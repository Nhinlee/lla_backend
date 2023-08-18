-- name: CreateTopic :one
INSERT INTO topics (
    id,
    name,
    created_at,
    updated_at
) VALUES (
    $1,
    $2,
    now(),
    now()
) RETURNING *;

-- name: GetAllTopics :many
SELECT id, name FROM topics 
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: DeleteTopic :one
UPDATE topics SET deleted_at = now() WHERE id = $1 RETURNING *;

-- name: GetTopicsAndTotalLearningItems :many
SELECT t.id, t.name, COUNT(li.id) AS total_learning_items
FROM topics t LEFT JOIN learning_items li ON t.id = li.topic_id
WHERE t.deleted_at IS NULL AND li.deleted_at IS NULL
GROUP BY t.id
ORDER BY t.created_at DESC;