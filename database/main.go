package main

import (
	"github.com/ansg191/northstars-backend/database/handler"
	pb "github.com/ansg191/northstars-backend/database/proto"
	"github.com/ansg191/northstars-backend/database/utils"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("database"),
		service.Version("latest"),
	)

	dbHandler := new(handler.Database)

	var err error
	dbHandler.DB, err = utils.LoadDB()
	if err != nil {
		logger.Fatal(err)
	}

	// Register handler
	pb.RegisterDatabaseHandler(srv.Server(), dbHandler)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
