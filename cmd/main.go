package main

import (
	"context"
	"fmt"
	"github.com/RichardLQ/confs"
	"github.com/RichardLQ/fix-srv/auth"
	"github.com/RichardLQ/fix-srv/client"
	"github.com/RichardLQ/fix-srv/proto/stub"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strings"
)

var RpcSrv *grpc.Server

func main() {
	confs.NewStart().BinComb(&client.Global)
	go startRpc()
	startHttp()
}

func startRpc() {
	//启动rpc服务
	RpcSrv := grpc.NewServer()
	stub.RegisterFixServiceServer(RpcSrv, stub.UserService)
	lis, _ := net.Listen("tcp", client.Global.FixServiceConf.RpcIp+":"+
		client.Global.FixServiceConf.RpcPort)
	RpcSrv.Serve(lis)
}
func startHttp() {
	//启动http服务
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithInsecure()}
	err := stub.RegisterFixServiceHandlerFromEndpoint(context.Background(),
		gwmux, client.Global.FixServiceConf.RpcIp+":"+
			client.Global.FixServiceConf.RpcPort, opt)
	if err != nil {
		log.Fatal(fmt.Errorf("启动rpc失败:%+v", err))
	}
	httpServer := &http.Server{
		Addr: client.Global.FixServiceConf.HttpIp+":"+
			client.Global.FixServiceConf.HttpPort,
		//Handler: gwmux,
		Handler: grpcHandlerFunc(RpcSrv, gwmux),
	}
	httpServer.ListenAndServe()
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			a := auth.MyClaims{Token: r.Header.Get("token")}
			_, err := a.Decryption()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			allowCORS(otherHandler).ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	fmt.Println("preflight request for:", r.URL.Path)
	return
}
