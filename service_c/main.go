package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"jaeger-demo/service_c/pb"
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

type Model struct {
	db *sql.DB
}

func NewModel(db *sql.DB) *Model {
	return &Model{db: db}
}

type AvatarProfile struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func (m *Model) GetAvatarURL(ctx context.Context, username string) (string, error) {
	var url string
	span, _ := opentracing.StartSpanFromContext(ctx, "Some Func")
	span.SetTag("db", "demo.avatar")
	span.LogFields(log.String("event", "xxxx"))
	span.SetTag("error", true)
	defer span.Finish()
	stmt, err := m.db.Prepare(`select avatar from demo.avatar where username = ? limit 1`)
	if err != nil {
		return "", err
	}

	err = stmt.QueryRow(username).Scan(&url)
	if err != nil {
		span.LogFields(log.String("err", err.Error()))
		return "", err
	}

	return url, nil
}

type Server struct {
	model *Model
}

func (s *Server) GetAvatarURL(ctx context.Context, request *pb.GetAvatarURLRequest) (*pb.GetAvatarURLReply, error) {
	username := request.Username
	if username == "" {
		return nil, status.Error(codes.InvalidArgument, "username 不能为空")
	}
	url, err := s.model.GetAvatarURL(ctx, username)
	if err != nil {
		fmt.Println(err)
		return nil, status.Error(codes.Internal, "数据库出现错误")
	}
	return &pb.GetAvatarURLReply{Url: url}, nil
}

func NewServer(model *Model) *Server {
	return &Server{model: model}
}

func main() {
	tracer, closer := initTracer("service_c")
	opentracing.InitGlobalTracer(tracer)
	defer closer.Close()
	srv := grpc.NewServer(grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)),
	)
	model := NewModel(NewDB())
	// register grpc server
	pb.RegisterAvatarServer(srv, NewServer(model))

	lis, err := net.Listen("tcp", ":5003")
	checkErr(err)
	fmt.Println("start service_c grpc service")
	err = srv.Serve(lis)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo")
	checkErr(err)
	return db
}
