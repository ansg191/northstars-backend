package main

import (
	"github.com/ansg191/northstars-backend/touchpad/handler"
	pb "github.com/ansg191/northstars-backend/touchpad/proto"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("touchpad"),
		service.Version("latest"),
	)

	touchpadHandler := new(handler.Touchpad)

	// Register handler
	pb.RegisterTouchpadHandler(srv.Server(), touchpadHandler)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
