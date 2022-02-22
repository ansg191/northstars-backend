package handler

import (
	"context"
	"errors"
	database "github.com/ansg191/northstars-backend/database/proto"
	"github.com/ansg191/northstars-backend/database/utils"
	"github.com/gotidy/ptr"
	microErr "github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// CreateAccount is a single request handler called via client.Call or the generated client code
func (e *Database) CreateAccount(ctx context.Context, req *database.CreateAccountRequest, _ *database.CreateAccountResponse) error {
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
				return nil
			}
			return err
		}
	} else if email, ok := req.GetIdentifier().(*database.GetAccountRequest_Email); ok {
		if err := e.DB.WithContext(ctx).
			Where("email = ?", email.Email).
			First(&account).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
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

func (e *Database) AddPhoneNumber(ctx context.Context, req *database.AddPhoneNumberRequest, _ *database.AddPhoneNumberResponse) error {
	if req.PhoneNumber == nil {
		return microErr.BadRequest("database.AddPhoneNumber", "Phone Number not provided")
	}

	account := utils.Account{ID: req.Id}

	formattedNumber, err := utils.FormatPhoneNumbers(req.PhoneNumber.Number)
	if err != nil {
		return err
	}

	phoneNumber := utils.PhoneNumber{
		Number:     formattedNumber,
		SmsEnabled: req.PhoneNumber.SmsEnabled,
		AccountID:  req.Id,
	}

	var exists bool
	err = e.DB.WithContext(ctx).
		Model(&phoneNumber).
		Select("count(*) > 0").
		Where("number = ? AND account_id = ?", formattedNumber, req.Id).
		Find(&exists).Error
	if err != nil {
		return err
	}

	if exists {
		return microErr.Conflict("database.AddPhoneNumber", "Phone number %s for account %d already exists", formattedNumber, req.Id)
	}

	err = e.DB.WithContext(ctx).
		Model(&account).
		Association("PhoneNumbers").
		Append(&phoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (e *Database) RemovePhoneNumber(ctx context.Context, req *database.RemovePhoneNumberRequest, _ *database.RemovePhoneNumberResponse) error {
	formattedNumber, err := utils.FormatPhoneNumbers(req.PhoneNumber)
	if err != nil {
		return err
	}

	var number utils.PhoneNumber
	err = e.DB.WithContext(ctx).
		Where("number = ? AND account_id = ?", formattedNumber, req.Id).
		First(&number).
		Error
	if err != nil {
		return err
	}

	err = e.DB.WithContext(ctx).
		Delete(&number).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (e *Database) CreateSwimmer(ctx context.Context, req *database.CreateSwimmerRequest, _ *database.CreateSwimmerResponse) error {
	swimmer := req.Swimmer
	if swimmer == nil {
		return microErr.BadRequest("database.CreateSwimmer", "Swimmer not provided")
	}

	middleInitial := (*string)(nil)
	if swimmer.MiddleInitial != "" {
		middleInitial = &swimmer.MiddleInitial
	}
	preferredName := (*string)(nil)
	if swimmer.PreferredName != "" {
		preferredName = &swimmer.PreferredName
	}

	protoSwimmer := utils.Swimmer{
		ID:              swimmer.Id,
		AccountID:       swimmer.AccountId,
		DOB:             swimmer.Dob.AsTime(),
		DateJoined:      swimmer.DateJoined.AsTime(),
		FirstName:       swimmer.FirstName,
		MiddleInitial:   middleInitial,
		LastName:        swimmer.LastName,
		PreferredName:   preferredName,
		Sex:             swimmer.Sex.String(),
		SwimmerIdentity: swimmer.SwimmerIdentity,
		RosterID:        swimmer.RosterId,
		Watchers:        nil,
	}

	err := e.DB.WithContext(ctx).Create(&protoSwimmer).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *Database) WatchSwimmer(ctx context.Context, req *database.WatchSwimmerRequest, _ *database.WatchSwimmerResponse) error {
	account := utils.Account{ID: req.AccountId}
	swimmer := utils.Swimmer{ID: req.SwimmerId}

	err := e.DB.WithContext(ctx).
		Model(&account).
		Association("WatchedSwimmers").
		Append(&swimmer)
	if err != nil {
		return err
	}

	return nil
}

func (e *Database) GetSwimmer(ctx context.Context, req *database.GetSwimmerRequest, rsp *database.GetSwimmerResponse) error {
	var swimmer utils.Swimmer
	if id, ok := req.Identifier.(*database.GetSwimmerRequest_Id); ok {
		err := e.DB.WithContext(ctx).
			Debug().
			Preload("Watchers").
			First(&swimmer, id.Id).
			Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
	} else if identity, ok := req.Identifier.(*database.GetSwimmerRequest_SwimmerIdentity); ok {
		err := e.DB.WithContext(ctx).
			Where("swimmer_identity = ?", identity.SwimmerIdentity).
			First(&swimmer).
			Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
	} else if identifier, ok := req.Identifier.(*database.GetSwimmerRequest_Name); ok {
		err := e.DB.WithContext(ctx).
			Where(
				"first_name = ? AND last_name = ? AND dob = ?",
				identifier.Name.FirstName,
				identifier.Name.LastName,
				identifier.Name.Dob.AsTime(),
			).
			First(&swimmer).
			Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
	} else {
		return microErr.BadRequest("database.GetSwimmer", "Identifier not provided")
	}

	protoSwimmer := database.Swimmer{
		Id:              swimmer.ID,
		AccountId:       swimmer.AccountID,
		Dob:             timestamppb.New(swimmer.DOB),
		DateJoined:      timestamppb.New(swimmer.DateJoined),
		FirstName:       swimmer.FirstName,
		MiddleInitial:   ptr.ToString(swimmer.MiddleInitial),
		LastName:        swimmer.LastName,
		PreferredName:   ptr.ToString(swimmer.PreferredName),
		SwimmerIdentity: swimmer.SwimmerIdentity,
		RosterId:        swimmer.RosterID,
		Watchers:        nil,
	}

	switch swimmer.Sex {
	case "MALE":
		protoSwimmer.Sex = database.Swimmer_MALE
	case "FEMALE":
		protoSwimmer.Sex = database.Swimmer_FEMALE
	case "OTHER":
		protoSwimmer.Sex = database.Swimmer_OTHER
	}

	for _, watcher := range swimmer.Watchers {
		protoAccount := database.Account{
			Id:           watcher.ID,
			Email:        watcher.Email,
			FirstName:    watcher.FirstName,
			LastName:     watcher.LastName,
			JoinDate:     timestamppb.New(watcher.JoinDate),
			PhoneNumbers: nil,
		}
		protoSwimmer.Watchers = append(protoSwimmer.Watchers, &protoAccount)
	}

	rsp.Swimmer = &protoSwimmer

	return nil
}

func (e *Database) UnWatchSwimmer(ctx context.Context, req *database.UnWatchSwimmerRequest, _ *database.UnWatchSwimmerResponse) error {
	account := utils.Account{ID: req.AccountId}
	swimmer := utils.Swimmer{ID: req.SwimmerId}

	err := e.DB.WithContext(ctx).
		Model(&account).
		Association("WatchedSwimmers").
		Delete(&swimmer)
	if err != nil {
		return err
	}

	return nil
}
