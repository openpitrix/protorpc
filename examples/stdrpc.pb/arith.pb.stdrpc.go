// Code generated by protoc-gen-stdrpc. DO NOT EDIT.
//
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-gen-stdrpc
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-plugin-common
//
// source: arith.proto

package service

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/golang/protobuf/proto"
)

var (
	_ = fmt.Sprint
	_ = io.Reader(nil)
	_ = log.Print
	_ = net.Addr(nil)
	_ = rpc.Call{}
	_ = time.Second

	_ = proto.String
)

type ArithService interface {
	Add(in *ArithRequest, out *ArithResponse) error
	Mul(in *ArithRequest, out *ArithResponse) error
	Div(in *ArithRequest, out *ArithResponse) error
	Error(in *ArithRequest, out *ArithResponse) error
}

// AcceptArithServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func AcceptArithServiceClient(lis net.Listener, x ArithService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("ArithService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// RegisterArithService publish the given ArithService implementation on the server.
func RegisterArithService(srv *rpc.Server, x ArithService) error {
	if err := srv.RegisterName("ArithService", x); err != nil {
		return err
	}
	return nil
}

// NewArithServiceServer returns a new ArithService Server.
func NewArithServiceServer(x ArithService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("ArithService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// ListenAndServeArithService listen announces on the local network address laddr
// and serves the given ArithService implementation.
func ListenAndServeArithService(network, addr string, x ArithService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("ArithService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeConn(conn)
	}
}

// ServeArithService serves the given ArithService implementation.
func ServeArithService(conn io.ReadWriteCloser, x ArithService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("ArithService", x); err != nil {
		log.Fatal(err)
	}
	srv.ServeConn(conn)
}

type ArithServiceClient struct {
	*rpc.Client
}

// NewArithServiceClient returns a ArithService stub to handle
// requests to the set of ArithService at the other end of the connection.
func NewArithServiceClient(conn io.ReadWriteCloser) *ArithServiceClient {
	c := rpc.NewClient(conn)
	return &ArithServiceClient{c}
}

func (c *ArithServiceClient) Add(in *ArithRequest) (out *ArithResponse, err error) {
	if in == nil {
		in = new(ArithRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ArithResponse)
	if err = c.Call("ArithService.Add", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *ArithServiceClient) AsyncAdd(in *ArithRequest, out *ArithResponse, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ArithRequest)
	}
	return c.Go(
		"ArithService.Add",
		in, out,
		done,
	)
}

func (c *ArithServiceClient) Mul(in *ArithRequest) (out *ArithResponse, err error) {
	if in == nil {
		in = new(ArithRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ArithResponse)
	if err = c.Call("ArithService.Mul", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *ArithServiceClient) AsyncMul(in *ArithRequest, out *ArithResponse, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ArithRequest)
	}
	return c.Go(
		"ArithService.Mul",
		in, out,
		done,
	)
}

func (c *ArithServiceClient) Div(in *ArithRequest) (out *ArithResponse, err error) {
	if in == nil {
		in = new(ArithRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ArithResponse)
	if err = c.Call("ArithService.Div", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *ArithServiceClient) AsyncDiv(in *ArithRequest, out *ArithResponse, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ArithRequest)
	}
	return c.Go(
		"ArithService.Div",
		in, out,
		done,
	)
}

func (c *ArithServiceClient) Error(in *ArithRequest) (out *ArithResponse, err error) {
	if in == nil {
		in = new(ArithRequest)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ArithResponse)
	if err = c.Call("ArithService.Error", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *ArithServiceClient) AsyncError(in *ArithRequest, out *ArithResponse, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ArithRequest)
	}
	return c.Go(
		"ArithService.Error",
		in, out,
		done,
	)
}

// DialArithService connects to an ArithService at the specified network address.
func DialArithService(network, addr string) (*ArithServiceClient, error) {
	c, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &ArithServiceClient{c}, nil
}

// DialArithServiceTimeout connects to an ArithService at the specified network address.
func DialArithServiceTimeout(network, addr string, timeout time.Duration) (*ArithServiceClient, error) {
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &ArithServiceClient{rpc.NewClient(conn)}, nil
}
