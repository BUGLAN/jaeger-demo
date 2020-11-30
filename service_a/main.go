package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"io"
	"jaeger-demo/service_b/pb"
	"net/http"
)

func initTracer(serviceName string) (opentracing.Tracer, io.Closer) {
	sender, err := jaeger.NewUDPTransport("127.0.0.1:6831", 0)
	if err != nil {
		panic(err)
	}

	// TODO 使用jaeger.NewHTTPHeaderPropagator()
	// TODO 生成tracer id
	tracer, closer := jaeger.NewTracer(
		serviceName,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(sender),
	)
	return tracer, closer
}

func main() {
	tracer, closer := initTracer("service_a_gin")
	defer closer.Close()
	opentracing.InitGlobalTracer(tracer)

	engine := gin.Default()
	engine.Use(ginhttp.Middleware(tracer))
	route(engine)
	engine.Run(":5001")
}

func route(engine *gin.Engine) {
	fn := func(c *gin.Context) {
		span := opentracing.SpanFromContext(c.Request.Context())
		if span == nil {
			panic(errors.New("span is nil"))
		}
	}
	engine.GET("/", fn)
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
		return
	})
	engine.GET("/avatar", getAvatar)
}

func getUseClient() pb.UserClient {
	tracer := opentracing.GlobalTracer()
	conn, err := grpc.Dial("127.0.0.1:5002", grpc.WithInsecure(), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracer)))
	if err != nil {
		panic(err)
	}

	return pb.NewUserClient(conn)
}

func getAvatar(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	span, _ := opentracing.StartSpanFromContextWithTracer(c, tracer, "geta")
	span.SetTag("a", "b")
	defer span.Finish()
	client := getUseClient()
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "sorry, username is 必须的",
		})
		return
	}
	reply, err := client.GetAvatar(context.TODO(), &pb.GetAvatarRequest{Username: username})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"avatar": reply.Avatar,
	})
	return
}
