package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/gowthamand7/proto-example/ytservice"
)

type VideoServiceServer struct {
	*DB
	ytservice.UnimplementedVideoServiceServer
}

func (s *VideoServiceServer) GetVideo(ctx context.Context, in *ytservice.VideoRequest) (*ytservice.Video, error) {

	v, ok := s.DB.Videos[in.Id]

	if !ok {
		return &ytservice.Video{}, fmt.Errorf("video not found, video id : %s", in.Id)
	}

	return &v, nil
}

func (s *VideoServiceServer) CreateVideo(ctx context.Context, in *ytservice.CreateVideoRequest) (*ytservice.Video, error) {

	user, ok := s.DB.Users[in.UserId]
	if !ok {
		return &ytservice.Video{}, fmt.Errorf("user not found, user %s", in.UserId)
	}
	video := &ytservice.Video{
		Id:          uuid.New().String(),
		User:        &user,
		Title:       in.Title,
		Description: in.Description,
		Comments:    make(map[string]*ytservice.Comment),
	}

	s.DB.Videos[video.Id] = *video
	return video, nil
}

func (s *VideoServiceServer) UpdateVideo(ctx context.Context, in *ytservice.UpdateVideoRequest) (*ytservice.Video, error) {

	v, ok := s.DB.Videos[in.Id]

	if !ok {
		return &ytservice.Video{}, fmt.Errorf("video not found, video id : %s", in.Id)
	}

	v.Title = in.Title
	v.Description = in.Description

	s.DB.Videos[in.Id] = v

	return &v, nil
}

func (s *VideoServiceServer) DeleteVideo(ctx context.Context, in *ytservice.VideoRequest) (*ytservice.Video, error) {

	v, ok := s.DB.Videos[in.Id]

	if !ok {
		return &ytservice.Video{}, fmt.Errorf("video not found, video id : %s", in.Id)
	}
	delete(s.Videos, v.Id)

	return &v, nil
}

func (s *VideoServiceServer) ListVideo(ctx context.Context, in *ytservice.ListVideoRequest) (*ytservice.Videos, error) {

	user, ok := s.DB.Users[in.UserId]
	if !ok {
		return &ytservice.Videos{}, fmt.Errorf("user not found, user %s", in.UserId)
	}

	videos := &ytservice.Videos{
		Videos: make(map[string]*ytservice.Video, 0),
	}

	for k, v := range s.Videos {
		if v.User.Id == user.Id {
			videos.Videos[k] = &ytservice.Video{
				Id:          v.Id,
				Title:       v.Title,
				Description: v.Description,
				User:        &user,
			}
		}
	}

	return videos, nil
}
