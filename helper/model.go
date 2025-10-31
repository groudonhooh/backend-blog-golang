package helper

import (
	"belajar-rest-api-golang/model/domain"
	"belajar-rest-api-golang/model/web"
)

func ToPostResponse(post domain.Post) web.PostResponse {
	return web.PostResponse{
		Id:        post.Id,
		Title:     post.Title,
		Slug:      post.Slug,
		Content:   post.Content,
		ImageURL:  post.ImageURL,
		AuthorId:  post.AuthorId,
		CreatedAt: post.CreatedAt,
	}
}

func ToCommentResponse(comment domain.Comment) web.CommentResponse {
	return web.CommentResponse{
		Id:        comment.Id,
		PostId:    comment.PostId,
		Content:   comment.Content,
		AuthorId:  comment.AuthorId,
		CreatedAt: comment.CreatedAt,
	}
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToPostResponses(posts []domain.Post) []web.PostResponse {
	var postResponses []web.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, ToPostResponse(post))
	}
	return postResponses
}
