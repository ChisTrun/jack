package server

import (
	"fmt"
	"log"
	"net"
	pb0 "server/api"
	"server/internal/server/jack"
	"server/package/config"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func Serve(cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb0.RegisterJackServer(grpcServer, jack.NewServer())

	log.Printf("server is running on: %v:%v", cfg.Server.Host, cfg.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
