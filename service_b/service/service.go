package service

import (
	"context"
	"jaeger-demo/service_b/pb"
)

type User struct {
}

func (u *User) GetAvatar(ctx context.Context, request *pb.GetAvatarRequest) (*pb.GetAvatarReply, error) {
	username := request.GetUsername()
	return &pb.GetAvatarReply{Avatar: "https://www.baidu.com&avatar=" + username}, nil
}

func NewUser() *User {
	return &User{}
}
