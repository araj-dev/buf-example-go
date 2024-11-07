package main

import (
	"buf.build/gen/go/araj-dev-ind/example-go/connectrpc/go/v1/v1connect"
	v1 "buf.build/gen/go/araj-dev-ind/example-go/protocolbuffers/go/v1"
	"connectrpc.com/connect"
	"context"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type ApiServer struct{}

func (s *ApiServer) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1.GetUserResponse{
		Id:       1,
		Username: "aa",
	})
	res.Header().Set("Version", "v1")
	return res, nil
}

func main() {
	api := &ApiServer{}
	mux := http.NewServeMux()
	path, handler := v1connect.NewUserServiceHandler(api)
	mux.Handle(path, handler)
	corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{})) // corsのハンドラを追加した

	http.ListenAndServe(
		"localhost:8080",
		corsHandler,
	)
}
