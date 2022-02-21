// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/users.proto

package users

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Users service

func NewUsersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Users service

type UsersService interface {
	NewUser(ctx context.Context, in *NewUserRequest, opts ...client.CallOption) (*NewUserResponse, error)
}

type usersService struct {
	c    client.Client
	name string
}

func NewUsersService(name string, c client.Client) UsersService {
	return &usersService{
		c:    c,
		name: name,
	}
}

func (c *usersService) NewUser(ctx context.Context, in *NewUserRequest, opts ...client.CallOption) (*NewUserResponse, error) {
	req := c.c.NewRequest(c.name, "Users.NewUser", in)
	out := new(NewUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	NewUser(context.Context, *NewUserRequest, *NewUserResponse) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) error {
	type users interface {
		NewUser(ctx context.Context, in *NewUserRequest, out *NewUserResponse) error
	}
	type Users struct {
		users
	}
	h := &usersHandler{hdlr}
	return s.Handle(s.NewHandler(&Users{h}, opts...))
}

type usersHandler struct {
	UsersHandler
}

func (h *usersHandler) NewUser(ctx context.Context, in *NewUserRequest, out *NewUserResponse) error {
	return h.UsersHandler.NewUser(ctx, in, out)
}