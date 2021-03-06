package main

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"io"
	"jaeger-demo/service_b/pb"
	"jaeger-demo/service_b/service"
	cpb "jaeger-demo/service_c/pb"
	"net"
)

func initTracer(serviceName string) (opentracing.Tracer, io.Closer) {
	sender, err := jaeger.NewUDPTransport("127.0.0.1:6831", 0)
	if err != nil {
		panic(err)
	}
	tracer, closer := jaeger.NewTracer(
		serviceName,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(sender),
	)
	return tracer, closer
}

func main() {
	tracer, closer := initTracer("service_b_grpc")
	defer closer.Close()

	listener, err := net.Listen("tcp", ":5002")
	if err != nil {
		panic(err)

	}
	conn, err := grpc.Dial("127.0.0.1:5003", grpc.WithInsecure(), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracer)))

	if err != nil {
		panic(err)
	}
	client := cpb.NewAvatarClient(conn)

	userService := service.NewUser(client)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))
	pb.RegisterUserServer(s, userService)

	fmt.Println("start grpc server")
	if err := s.Serve(listener); err != nil {
		panic(err)
	}

}
