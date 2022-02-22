package handler

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	microErr "github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"

	database "github.com/ansg191/northstars-backend/database/proto"
	"github.com/ansg191/northstars-backend/database/utils"
)

type Database struct {
	DB *gorm.DB
}

// CreateAccount is a single request handler called via client.Call or the generated client code
func (e *Database) CreateAccount(ctx context.Context, req *database.CreateAccountRequest, rsp *database.CreateAccountResponse) error {
	log.Info("Received Database.CreateAccount request")

	if req.Id == 0 || req.Email == "" {
		return microErr.BadRequest("database.CreateAccount", "Email or ID not provided")
	}

	account := utils.Account{
		ID:           req.Id,
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		JoinDate:     req.JoinDate.AsTime(),
		PhoneNumbers: nil,
	}

	if err := e.DB.WithContext(ctx).Create(&account).Error; err != nil {
		return err
	}

	return nil
}

func (e *Database) GetAccount(ctx context.Context, req *database.GetAccountRequest, rsp *database.GetAccountResponse) error {
	var account utils.Account
	if id, ok := req.GetIdentifier().(*database.GetAccountRequest_Id); ok {
		if err := e.DB.WithContext(ctx).First(&account, id.Id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return microErr.NotFound("database.GetAccount", "Account with id %d not found", id.Id)
			}
			return err
		}
	} else if email, ok := req.GetIdentifier().(*database.GetAccountRequest_Email); ok {
		if err := e.DB.WithContext(ctx).
			Where("email = ?", email.Email).
			First(&account).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return microErr.NotFound("database.GetAccount", "Account with email %s not found", email.Email)
			}
			return err
		}
	} else {
		return microErr.BadRequest("database.GetAccount", "identifier not provided")
	}

	protoAccount := database.Account{
		Id:        account.ID,
		Email:     account.Email,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		JoinDate:  timestamppb.New(account.JoinDate),
	}

	for _, number := range account.PhoneNumbers {
		protoNumber := database.PhoneNumber{
			Number:     number.Number,
			SmsEnabled: number.SmsEnabled,
		}
		protoAccount.PhoneNumbers = append(protoAccount.PhoneNumbers, &protoNumber)
	}

	rsp.Account = &protoAccount

	return nil
}
