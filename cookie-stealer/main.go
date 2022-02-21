package main

import (
	"github.com/ansg191/northstars-backend/cookie-stealer/handler"
	pb "github.com/ansg191/northstars-backend/cookie-stealer/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("cookie-stealer"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterCookieStealerHandler(srv.Server(), new(handler.CookieStealer))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
