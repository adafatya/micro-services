package main

import (
	"fmt"
	"net"

	"github.com/adafatya/micro-services/auth-service/internal/config"
	"github.com/adafatya/micro-services/auth-service/pkg/util"
)

func main() {
	db := config.NewDatabase()

	s := config.NewGRPCServer()

	config.Bootstrap(&config.BootstrapConfig{
		DB:         db,
		GRPCServer: s,
	})

	host := util.GetEnv("SERVER_HOST", "")
	port := util.GetEnv("SERVER_PORT", "")
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		panic(err)
	}

	fmt.Println("starting server ...")

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
