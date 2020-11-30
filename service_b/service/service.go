package service

import (
	"context"
	"jaeger-demo/service_b/pb"
	cpb "jaeger-demo/service_c/pb"
)

type User struct {
	Client cpb.AvatarClient
}

func (u *User) GetAvatar(ctx context.Context, request *pb.GetAvatarRequest) (*pb.GetAvatarReply, error) {
	username := request.GetUsername()
	reply, err := u.Client.GetAvatarURL(ctx, &cpb.GetAvatarURLRequest{Username: username})
	if err != nil {
		return nil, err
	}

	return &pb.GetAvatarReply{Avatar: reply.Url}, nil
}

func NewUser(client cpb.AvatarClient) *User {
	return &User{Client: client}
}
