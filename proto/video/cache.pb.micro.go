// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/video/cache.proto

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

// Client API for Cache service

type CacheService interface {
	// 将数据转存为文件
	Save(ctx context.Context, opts ...client.CallOption) (Cache_SaveService, error)
}

type cacheService struct {
	c    client.Client
	name string
}

func NewCacheService(name string, c client.Client) CacheService {
	return &cacheService{
		c:    c,
		name: name,
	}
}

func (c *cacheService) Save(ctx context.Context, opts ...client.CallOption) (Cache_SaveService, error) {
	req := c.c.NewRequest(c.name, "Cache.Save", &SaveRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &cacheServiceSave{stream}, nil
}

type Cache_SaveService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*SaveRequest) error
}

type cacheServiceSave struct {
	stream client.Stream
}

func (x *cacheServiceSave) Close() error {
	return x.stream.Close()
}

func (x *cacheServiceSave) Context() context.Context {
	return x.stream.Context()
}

func (x *cacheServiceSave) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *cacheServiceSave) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *cacheServiceSave) Send(m *SaveRequest) error {
	return x.stream.Send(m)
}

// Server API for Cache service

type CacheHandler interface {
	// 将数据转存为文件
	Save(context.Context, Cache_SaveStream) error
}

func RegisterCacheHandler(s server.Server, hdlr CacheHandler, opts ...server.HandlerOption) error {
	type cache interface {
		Save(ctx context.Context, stream server.Stream) error
	}
	type Cache struct {
		cache
	}
	h := &cacheHandler{hdlr}
	return s.Handle(s.NewHandler(&Cache{h}, opts...))
}

type cacheHandler struct {
	CacheHandler
}

func (h *cacheHandler) Save(ctx context.Context, stream server.Stream) error {
	return h.CacheHandler.Save(ctx, &cacheSaveStream{stream})
}

type Cache_SaveStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*SaveRequest, error)
}

type cacheSaveStream struct {
	stream server.Stream
}

func (x *cacheSaveStream) Close() error {
	return x.stream.Close()
}

func (x *cacheSaveStream) Context() context.Context {
	return x.stream.Context()
}

func (x *cacheSaveStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *cacheSaveStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *cacheSaveStream) Recv() (*SaveRequest, error) {
	m := new(SaveRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
