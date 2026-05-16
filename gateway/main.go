package main

import (
	"context"
	"net/http"
	pb "proto/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// 将 HTTP 请求注册到 mux，并指定 gRPC 服务地址
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		panic(err)
	}

	// 启动 HTTP 服务
	println("HTTP Gateway listening on :8080")
	http.ListenAndServe(":8080", mux)
}
