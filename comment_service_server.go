package main

import (
	"context"

	"github.com/gowthamand7/proto-example/ytservice"
)

type CommentServiceServer struct {
	*DB
	ytservice.UnimplementedCommentServiceServer
}

func (s *CommentServiceServer) CreateComment(ctx context.Context, in *ytservice.CreateCommentRequest) (*ytservice.Comment, error) {

	return &ytservice.Comment{}, nil
}

func (s *CommentServiceServer) GetComment(ctx context.Context, in *ytservice.GetCommentRequest) (*ytservice.Comment, error) {
	return &ytservice.Comment{}, nil
}

func (s *CommentServiceServer) ListComments(ctx context.Context, in *ytservice.ListCommentsRequest) (*ytservice.Comments, error) {
	return &ytservice.Comments{}, nil
}

func (s *CommentServiceServer) UpdateComment(ctx context.Context, in *ytservice.UpdateCommentRequest) (*ytservice.Comment, error) {
	return &ytservice.Comment{}, nil
}

func (s *CommentServiceServer) DeleteComment(ctx context.Context, in *ytservice.GetCommentRequest) (*ytservice.Comment, error) {
	return &ytservice.Comment{}, nil
}
