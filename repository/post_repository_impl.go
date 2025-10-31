package repository

import (
	"belajar-rest-api-golang/helper"
	"belajar-rest-api-golang/model/domain"
	"context"
	"database/sql"
	"errors"
)

type PostRepositoryImpl struct {
}

func NewPostRepository() PostRepository {
	return &PostRepositoryImpl{}
}

func (repository *PostRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "INSERT INTO posts(title, slug, content, image_url, author_id) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, post.Title, post.Slug, post.Content, post.ImageURL, post.AuthorId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	query := "SELECT id, title, slug, content, image_url, author_id, created_at FROM posts WHERE id = ?"
	row := tx.QueryRowContext(ctx, query, id)

	err = row.Scan(&post.Id, &post.Title, &post.Slug, &post.Content, &post.ImageURL, &post.AuthorId, &post.CreatedAt)
	helper.PanicIfError(err)

	return post
}

func (repository *PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "update posts set title = ?, content = ?, image_url = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, post.Title, post.Content, post.ImageURL, post.Id)
	helper.PanicIfError(err)

	return post
}

func (repository *PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, post domain.Post) {
	SQL := "delete from posts where id = ?"
	_, err := tx.ExecContext(ctx, SQL, post.Id)
	helper.PanicIfError(err)

}

func (repository *PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error) {
	SQL := "Select id, title, content, image_url, author_id, created_at from posts where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	post := domain.Post{}
	if rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.ImageURL, &post.AuthorId, &post.CreatedAt)
		helper.PanicIfError(err)
		return post, nil
	} else {
		return post, errors.New("post is not found")
	}
}

func (repository *PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Post {
	SQL := "Select id, title, content, image_url, author_id, created_at from posts"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		post := domain.Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.ImageURL, &post.AuthorId, &post.CreatedAt)
		helper.PanicIfError(err)
		posts = append(posts, post)
	}
	return posts
}

func (repository *PostRepositoryImpl) FindAuthorIdByPostId(ctx context.Context, tx *sql.Tx, postId int) (int, error) {
	SQL := "SELECT author_id FROM posts WHERE id = ?"
	row := tx.QueryRowContext(ctx, SQL, postId)

	var authorId int
	err := row.Scan(&authorId)
	if err != nil {
		return 0, err
	}

	return authorId, nil
}
