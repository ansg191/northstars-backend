package handler

import (
	"context"
	"github.com/ansg191/northstars-backend/database/utils"
	"github.com/micro/micro/v3/service/errors"
	"gorm.io/gorm"

	log "github.com/micro/micro/v3/service/logger"

	database "github.com/ansg191/northstars-backend/database/proto"
)

type Database struct {
	DB *gorm.DB
}

// CreateAccount is a single request handler called via client.Call or the generated client code
func (e *Database) CreateAccount(ctx context.Context, req *database.CreateAccountRequest, rsp *database.CreateAccountResponse) error {
	log.Info("Received Database.CreateAccount request")

	if req.Id == 0 || req.Email == "" {
		return errors.BadRequest("database.CreateAccount", "Email or ID not provided")
	}

	account := utils.Account{
		ID:           req.Id,
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumbers: nil,
	}

	if err := e.DB.WithContext(ctx).Create(&account).Error; err != nil {
		return err
	}

	return nil
}
