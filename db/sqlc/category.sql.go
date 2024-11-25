// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories(
	 category_name, description, picture)
	VALUES ($1, $2, $3) RETURNING category_id, category_name, description, picture
`

type CreateCategoryParams struct {
	CategoryName string  `json:"category_name"`
	Description  *string `json:"description"`
	Picture      []byte  `json:"picture"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (*Category, error) {
	row := q.db.QueryRow(ctx, createCategory, arg.CategoryName, arg.Description, arg.Picture)
	var i Category
	err := row.Scan(
		&i.CategoryID,
		&i.CategoryName,
		&i.Description,
		&i.Picture,
	)
	return &i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
	WHERE category_id=$1
    RETURNING category_id, category_name, description, picture
`

func (q *Queries) DeleteCategory(ctx context.Context, categoryID int32) error {
	_, err := q.db.Exec(ctx, deleteCategory, categoryID)
	return err
}

const findAllCategory = `-- name: FindAllCategory :many
SELECT category_id, category_name, description, picture
	FROM categories
`

func (q *Queries) FindAllCategory(ctx context.Context) ([]*Category, error) {
	rows, err := q.db.Query(ctx, findAllCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.Description,
			&i.Picture,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findCategoryById = `-- name: FindCategoryById :one
SELECT category_id, category_name, description, picture
	FROM categorieS WHERE category_id=$1
`

func (q *Queries) FindCategoryById(ctx context.Context, categoryID int32) (*Category, error) {
	row := q.db.QueryRow(ctx, findCategoryById, categoryID)
	var i Category
	err := row.Scan(
		&i.CategoryID,
		&i.CategoryName,
		&i.Description,
		&i.Picture,
	)
	return &i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories
	SET category_name=$2, description=$3, picture=$4
	WHERE category_id=$1    
    RETURNING category_id, category_name, description, picture
`

type UpdateCategoryParams struct {
	CategoryID   int32   `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Description  *string `json:"description"`
	Picture      []byte  `json:"picture"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (*Category, error) {
	row := q.db.QueryRow(ctx, updateCategory,
		arg.CategoryID,
		arg.CategoryName,
		arg.Description,
		arg.Picture,
	)
	var i Category
	err := row.Scan(
		&i.CategoryID,
		&i.CategoryName,
		&i.Description,
		&i.Picture,
	)
	return &i, err
}
