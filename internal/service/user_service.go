package service

import (
	"github.com/itmrchow/todolist-proto/protobuf/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	accountConn     *grpc.ClientConn
	accountLocation string
	accountOptions  []grpc.DialOption
}

func NewUserService(location string, options ...grpc.DialOption) (svc *UserService, err error) {

	svc = &UserService{
		accountLocation: location,
		accountOptions:  options,
	}

	svc.accountOptions = append(svc.accountOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))

	svc.accountConn, err = grpc.NewClient(svc.accountLocation, svc.accountOptions...)
	if err != nil {
		return nil, err
	}

	return
}

func (u *UserService) NewClient() (client user.UserServiceClient, err error) {

	if u.accountConn == nil {
		u.accountConn, err = grpc.NewClient(u.accountLocation, u.accountOptions...)
		if err != nil {
			return nil, err
		}
	}

	// todo: reconnect

	client = user.NewUserServiceClient(u.accountConn)

	return
}
