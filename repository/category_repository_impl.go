package repository

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(name, slug) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name, category.Slug)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	query := "SELECT id, name, slug, created_at FROM categories WHERE id = ?"
	row := tx.QueryRowContext(ctx, query, id)
	err = row.Scan(&category.Id, &category.Name, &category.Slug, &category.CreatedAt)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name, slug, created_at FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		rows.Scan(&category.Id, &category.Name, &category.Slug, &category.CreatedAt)
		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name, slug, created_at FROM categories WHERE id = ?"
	row := tx.QueryRowContext(ctx, SQL, categoryId)

	var category domain.Category
	err := row.Scan(&category.Id, &category.Name, &category.Slug, &category.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return category, nil
		}
		helper.PanicIfError(err)
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}
