-- name: FindCategoryById :one
SELECT category_id, category_name, description, picture
	FROM categorieS WHERE category_id=$1;

-- name: FindAllCategory :many
SELECT category_id, category_name, description, picture
	FROM categories;

-- name: CreateCategory :one
INSERT INTO categories(
	 category_name, description, picture)
	VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateCategory :one
UPDATE categories
	SET category_name=$2, description=$3, picture=$4
	WHERE category_id=$1    
    RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
	WHERE category_id=$1
    RETURNING *;