package main

import (
	cookiestealer "github.com/ansg191/northstars-backend/cookie-stealer/proto"
	database "github.com/ansg191/northstars-backend/database/proto"
	"github.com/ansg191/northstars-backend/users/handler"
	pb "github.com/ansg191/northstars-backend/users/proto"
	"github.com/micro/micro/v3/proto/auth"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("users"),
		service.Version("latest"),
	)

	userHandler := new(handler.Users)
	userHandler.Cookies = cookiestealer.NewCookieStealerService("cookie-stealer", srv.Client())
	userHandler.Accounts = auth.NewAccountsService("auth", srv.Client())
	userHandler.DB = database.NewDatabaseService("database", srv.Client())

	// Register handler
	pb.RegisterUsersHandler(srv.Server(), userHandler)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
