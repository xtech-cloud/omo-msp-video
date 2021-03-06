// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/video/build.proto

package video

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Build service

type BuildService interface {
	// 将多个视频合并为一个视频
	Combine(ctx context.Context, in *CombineRequest, opts ...client.CallOption) (*BuildResponse, error)
	// 将一个视频切割为多个视频
	Split(ctx context.Context, in *SplitRequest, opts ...client.CallOption) (*BuildResponse, error)
	// 将图片和音频制作为一个视频
	Render(ctx context.Context, in *RenderRequest, opts ...client.CallOption) (*BuildResponse, error)
}

type buildService struct {
	c    client.Client
	name string
}

func NewBuildService(name string, c client.Client) BuildService {
	return &buildService{
		c:    c,
		name: name,
	}
}

func (c *buildService) Combine(ctx context.Context, in *CombineRequest, opts ...client.CallOption) (*BuildResponse, error) {
	req := c.c.NewRequest(c.name, "Build.Combine", in)
	out := new(BuildResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildService) Split(ctx context.Context, in *SplitRequest, opts ...client.CallOption) (*BuildResponse, error) {
	req := c.c.NewRequest(c.name, "Build.Split", in)
	out := new(BuildResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildService) Render(ctx context.Context, in *RenderRequest, opts ...client.CallOption) (*BuildResponse, error) {
	req := c.c.NewRequest(c.name, "Build.Render", in)
	out := new(BuildResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Build service

type BuildHandler interface {
	// 将多个视频合并为一个视频
	Combine(context.Context, *CombineRequest, *BuildResponse) error
	// 将一个视频切割为多个视频
	Split(context.Context, *SplitRequest, *BuildResponse) error
	// 将图片和音频制作为一个视频
	Render(context.Context, *RenderRequest, *BuildResponse) error
}

func RegisterBuildHandler(s server.Server, hdlr BuildHandler, opts ...server.HandlerOption) error {
	type build interface {
		Combine(ctx context.Context, in *CombineRequest, out *BuildResponse) error
		Split(ctx context.Context, in *SplitRequest, out *BuildResponse) error
		Render(ctx context.Context, in *RenderRequest, out *BuildResponse) error
	}
	type Build struct {
		build
	}
	h := &buildHandler{hdlr}
	return s.Handle(s.NewHandler(&Build{h}, opts...))
}

type buildHandler struct {
	BuildHandler
}

func (h *buildHandler) Combine(ctx context.Context, in *CombineRequest, out *BuildResponse) error {
	return h.BuildHandler.Combine(ctx, in, out)
}

func (h *buildHandler) Split(ctx context.Context, in *SplitRequest, out *BuildResponse) error {
	return h.BuildHandler.Split(ctx, in, out)
}

func (h *buildHandler) Render(ctx context.Context, in *RenderRequest, out *BuildResponse) error {
	return h.BuildHandler.Render(ctx, in, out)
}
