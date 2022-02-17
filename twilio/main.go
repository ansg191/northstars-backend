package main

import (
	"twilio/handler"
	pb "twilio/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("twilio"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTwilioHandler(srv.Server(), new(handler.Twilio))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
