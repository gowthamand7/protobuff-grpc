package main

import (
	"log"
	"net"

	"github.com/gowthamand7/proto-example/ytservice"
	"google.golang.org/grpc"
)

type DB struct {
	Users    map[string]ytservice.User
	Videos   map[string]ytservice.Video
	Comments map[string]ytservice.Comment
}

func main() {

	var db DB = DB{
		Users:    make(map[string]ytservice.User),
		Videos:   make(map[string]ytservice.Video),
		Comments: make(map[string]ytservice.Comment),
	}

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	cs := &CommentServiceServer{
		DB: &db,
	}
	ytservice.RegisterCommentServiceServer(s, cs)

	us := &UserServiceServer{
		DB: &db,
	}
	ytservice.RegisterUserServiceServer(s, us)

	vs := &VideoServiceServer{
		DB: &db,
	}
	ytservice.RegisterVideoServiceServer(s, vs)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
