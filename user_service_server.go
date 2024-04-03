package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/gowthamand7/proto-example/ytservice"
)

type UserServiceServer struct {
	*DB
	ytservice.UnimplementedUserServiceServer
}

func (s *UserServiceServer) CreateUser(ctx context.Context, in *ytservice.CreateUserRequest) (*ytservice.User, error) {
	// Check if a user with the same name already exists
	if _, err := s.getUserByName(in.Name); err == nil {
		return &ytservice.User{}, fmt.Errorf("user already exists with the same name: %s", in.Name)
	}

	// Generate a new UUID for the user's ID
	userId := uuid.New().String()

	// Create a new user with the provided name and generated ID
	newUser := &ytservice.User{
		Name: in.Name,
		Id:   userId,
	}

	// Add the new user to the database
	s.DB.Users[newUser.Id] = *newUser

	// Print a message to indicate that the user has been created
	fmt.Fprintf(os.Stdout, "user created with id: %s, name: %s\n", newUser.Id, newUser.Name)
	fmt.Fprintf(os.Stdout, "%+v \n", s.DB.Users)

	// Return the newly created user
	return newUser, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, in *ytservice.GetUserRequest) (*ytservice.User, error) {

	user, ok := s.DB.Users[in.Id]

	if !ok {
		return &ytservice.User{}, fmt.Errorf("user not found %s", in.Id)
	}

	return &user, nil
}

func (s *UserServiceServer) ListUser(ctx context.Context, in *ytservice.ListUserRequest) (*ytservice.Users, error) {

	u := ytservice.Users{
		Users: make(map[string]*ytservice.User, 0),
	}

	for _, user := range *&s.Users {
		u.Users[user.Id] = &ytservice.User{
			Name: user.Name,
			Id:   user.Id,
		}
	}
	return &u, nil
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, in *ytservice.User) (*ytservice.User, error) {

	var user *ytservice.User
	var err error
	if user, err = s.GetUser(ctx, &ytservice.GetUserRequest{Id: in.Id}); err != nil {
		return &ytservice.User{}, fmt.Errorf("user not found %s", in.Id)
	}

	user.Name = in.Name
	s.Users[user.Id] = *user
	return user, nil
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, in *ytservice.GetUserRequest) (*ytservice.User, error) {
	var user *ytservice.User
	var err error
	if user, err = s.GetUser(ctx, &ytservice.GetUserRequest{Id: in.Id}); err != nil {
		return &ytservice.User{}, fmt.Errorf("user not found %s", in.Id)
	}

	delete(s.Users, user.Id)
	return user, nil
}

func (s *UserServiceServer) getUserByName(name string) (*ytservice.User, error) {

	for _, u := range *&s.Users {
		if u.Name == name {
			return &u, nil
		}
	}
	return &ytservice.User{}, fmt.Errorf("user not found %s", name)
}
