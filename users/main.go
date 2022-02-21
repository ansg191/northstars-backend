package main

import (
	"github.com/ansg191/northstars-backend/users/handler"
	pb "github.com/ansg191/northstars-backend/users/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("users"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUsersHandler(srv.Server(), new(handler.Users))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
