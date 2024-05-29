// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package pb

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginResp, error)
	Signup(ctx context.Context, in *SignUpReq, opts ...client.CallOption) (*SignUpResp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginResp, error) {
	req := c.c.NewRequest(c.name, "UserService.Login", in)
	out := new(LoginResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Signup(ctx context.Context, in *SignUpReq, opts ...client.CallOption) (*SignUpResp, error) {
	req := c.c.NewRequest(c.name, "UserService.Signup", in)
	out := new(SignUpResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Login(context.Context, *LoginReq, *LoginResp) error
	Signup(context.Context, *SignUpReq, *SignUpResp) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Login(ctx context.Context, in *LoginReq, out *LoginResp) error
		Signup(ctx context.Context, in *SignUpReq, out *SignUpResp) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Login(ctx context.Context, in *LoginReq, out *LoginResp) error {
	return h.UserServiceHandler.Login(ctx, in, out)
}

func (h *userServiceHandler) Signup(ctx context.Context, in *SignUpReq, out *SignUpResp) error {
	return h.UserServiceHandler.Signup(ctx, in, out)
}
