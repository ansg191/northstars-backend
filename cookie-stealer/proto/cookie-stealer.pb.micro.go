// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cookie-stealer.proto

package cookiestealer

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

// Api Endpoints for CookieStealer service

func NewCookieStealerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CookieStealer service

type CookieStealerService interface {
	StealTeamUnifyCookies(ctx context.Context, in *StealTeamUnifyCookiesRequest, opts ...client.CallOption) (*StealTeamUnifyCookiesResponse, error)
}

type cookieStealerService struct {
	c    client.Client
	name string
}

func NewCookieStealerService(name string, c client.Client) CookieStealerService {
	return &cookieStealerService{
		c:    c,
		name: name,
	}
}

func (c *cookieStealerService) StealTeamUnifyCookies(ctx context.Context, in *StealTeamUnifyCookiesRequest, opts ...client.CallOption) (*StealTeamUnifyCookiesResponse, error) {
	req := c.c.NewRequest(c.name, "CookieStealer.StealTeamUnifyCookies", in)
	out := new(StealTeamUnifyCookiesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CookieStealer service

type CookieStealerHandler interface {
	StealTeamUnifyCookies(context.Context, *StealTeamUnifyCookiesRequest, *StealTeamUnifyCookiesResponse) error
}

func RegisterCookieStealerHandler(s server.Server, hdlr CookieStealerHandler, opts ...server.HandlerOption) error {
	type cookieStealer interface {
		StealTeamUnifyCookies(ctx context.Context, in *StealTeamUnifyCookiesRequest, out *StealTeamUnifyCookiesResponse) error
	}
	type CookieStealer struct {
		cookieStealer
	}
	h := &cookieStealerHandler{hdlr}
	return s.Handle(s.NewHandler(&CookieStealer{h}, opts...))
}

type cookieStealerHandler struct {
	CookieStealerHandler
}

func (h *cookieStealerHandler) StealTeamUnifyCookies(ctx context.Context, in *StealTeamUnifyCookiesRequest, out *StealTeamUnifyCookiesResponse) error {
	return h.CookieStealerHandler.StealTeamUnifyCookies(ctx, in, out)
}