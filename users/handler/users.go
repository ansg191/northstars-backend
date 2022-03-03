package handler

import (
	"context"
	"fmt"
	database "github.com/ansg191/northstars-backend/database/proto"
	twilio "github.com/ansg191/northstars-backend/twilio/proto"
	"github.com/micro/micro/v3/service/client"
	"strconv"

	authProto "github.com/micro/micro/v3/proto/auth"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"google.golang.org/protobuf/types/known/timestamppb"

	cookiestealer "github.com/ansg191/northstars-backend/cookie-stealer/proto"
	users "github.com/ansg191/northstars-backend/users/proto"
	"github.com/ansg191/northstars-backend/users/utils"
)

type Users struct {
	Cookies  cookiestealer.CookieStealerService
	Accounts authProto.AccountsService
	DB       database.DatabaseService
	Twilio   twilio.TwilioService
}

func (e *Users) CheckAccount(ctx context.Context, req *users.CheckAccountRequest, rsp *users.CheckAccountResponse) error {
	if req.Email == "" {
		return errors.BadRequest("users.CheckAccount", "Email not provided")
	}

	if accountExists, err := utils.CheckAccountExists(ctx, e.DB, req.Email); err != nil {
		return err
	} else if accountExists {
		return errors.Forbidden("users.NewUser", "Account already exists for %s", req.Email)
	}

	cookies, err := utils.GetCookies(ctx, e.Cookies)
	if err != nil {
		return err
	}

	accounts, err := utils.ListAccounts(ctx, cookies)
	if err != nil {
		return err
	}

	var userAccount *utils.Account
	for _, account := range accounts.Result {
		if account.Email == req.Email {
			userAccount = &account
			break
		}
	}
	if userAccount == nil {
		rsp.Found = false
		return nil
	}

	account, err := utils.GetAccount(ctx, cookies, userAccount.Id)
	if err != nil {
		return err
	}

	rsp.Found = true
	rsp.Id = int32(account.Id)
	rsp.FirstName = account.FirstName
	rsp.LastName = account.LastName

	rsp.PhoneNumbers = utils.CollectPhoneNumbers(account)

	for i, number := range rsp.PhoneNumbers {
		rsp.PhoneNumbers[i] = utils.MaskPhoneNumbers(number)
	}

	return nil
}

func (e *Users) VerifyUser(ctx context.Context, req *users.VerifyUserRequest, rsp *users.VerifyUserResponse) error {
	cookies, err := utils.GetCookies(ctx, e.Cookies)
	if err != nil {
		return err
	}

	account, err := utils.GetAccount(ctx, cookies, int(req.Id))
	if err != nil {
		return err
	}

	phoneNumbers := utils.CollectPhoneNumbers(account)
	var phoneNumber string
	for _, number := range phoneNumbers {
		maskedNumber := utils.MaskPhoneNumbers(number)
		if maskedNumber == req.MaskedNumber {
			phoneNumber = number
		}
	}

	response, err := e.Twilio.Verify(ctx, &twilio.VerifyRequest{
		Destination: &twilio.VerifyRequest_Number{Number: phoneNumber},
	}, client.WithAuthToken())
	if err != nil {
		return err
	}

	rsp.Sid = response.Sid

	return nil
}

func (e *Users) CheckVerify(ctx context.Context, req *users.CheckVerifyRequest, rsp *users.CheckVerifyResponse) error {
	response, err := e.Twilio.CheckVerify(ctx, &twilio.CheckVerifyRequest{
		Sid:  req.Sid,
		Code: req.Code,
	}, client.WithAuthToken())
	if err != nil {
		return err
	}

	rsp.Status = response.Status == twilio.CheckVerifyResponse_APPROVED

	return nil
}

func (e *Users) NewUser(ctx context.Context, req *users.NewUserRequest, rsp *users.NewUserResponse) error {
	log.Info("Received Users.NewUser request")

	if req.Email == "" || req.Password == "" {
		return errors.BadRequest("users.NewUser", "Email and Password not provided")
	}

	if accountExists, err := utils.CheckAccountExists(ctx, e.DB, req.Email); err != nil {
		return err
	} else if accountExists {
		return errors.Forbidden("users.NewUser", "Account already exists for %s", req.Email)
	}

	cookies, err := utils.GetCookies(ctx, e.Cookies)
	if err != nil {
		return err
	}

	accounts, err := utils.ListAccounts(ctx, cookies)
	if err != nil {
		return err
	}

	var userAccount *utils.Account
	for _, account := range accounts.Result {
		if account.Email == req.Email {
			userAccount = &account
			break
		}
	}
	if userAccount == nil {
		return errors.NotFound("users.NewUser", "TeamUnify account with the login email %s not found", req.Email)
	}

	account, err := utils.GetAccount(ctx, cookies, userAccount.Id)
	if err != nil {
		return err
	}

	rsp.Id = int32(account.Id)
	rsp.Email = account.Email

	rsp.Sms, err = utils.ConvertPhoneNumbers(account.Sms1)
	if err != nil {
		return err
	}
	rsp.HomePhone, err = utils.ConvertPhoneNumbers(account.HomePhone)
	if err != nil {
		return err
	}
	rsp.WorkPhone, err = utils.ConvertPhoneNumbers(account.WorkPhone)
	if err != nil {
		return err
	}

	rsp.FirstName = account.FirstName
	rsp.MiddleInitial = account.Mi
	rsp.LastName = account.LastName
	rsp.Address = account.Address
	rsp.Address2 = account.Address2
	rsp.City = account.City
	rsp.State = account.State
	rsp.Zip = account.Zip
	rsp.PictureFile = "https://www.teamunify.com" + account.PictureFile

	joinTime, err := utils.ConvertTeamUnifyDate(account.JoinedDate1)
	if err != nil {
		return err
	}

	rsp.JoinedDate = timestamppb.New(joinTime)

	_, err = auth.Generate(
		req.Email,
		auth.WithSecret(req.Password),
		auth.WithName(fmt.Sprintf("%s %s", account.FirstName, account.LastName)),
		auth.WithScopes("user"),
		auth.WithMetadata(map[string]string{
			"tuId": strconv.Itoa(int(rsp.Id)),
		}),
	)
	if err != nil {
		return err
	}

	_, err = e.DB.CreateAccount(ctx, &database.CreateAccountRequest{
		Id:        rsp.Id,
		Email:     rsp.Email,
		FirstName: rsp.FirstName,
		LastName:  rsp.LastName,
		JoinDate:  rsp.JoinedDate,
	}, client.WithAuthToken())
	if err != nil {
		return err
	}

	return nil
}

func (e *Users) GetSwimmers(ctx context.Context, _ *users.GetSwimmersRequest, rsp *users.GetSwimmersResponse) error {
	account, exists := auth.AccountFromContext(ctx)
	if !exists {
		return errors.Unauthorized("users.GetSwimmers", "Unauthorized")
	}

	accountId, err := strconv.Atoi(account.Metadata["tuId"])
	if err != nil {
		return err
	}

	cookies, err := utils.GetCookies(ctx, e.Cookies)
	if err != nil {
		return err
	}

	members, err := utils.GetMembers(ctx, cookies, accountId)
	if err != nil {
		return err
	}

	for _, member := range members {
		newMember := users.Member{
			Id:              int32(member.Id),
			AccountId:       int32(member.AccountId),
			FirstName:       member.FirstName,
			MiddleInitial:   member.Mi,
			LastName:        member.LastName,
			PreferredName:   member.Prefer,
			SwimmerIdentity: member.SwimmerIdentity,
			RosterId:        int32(member.RosterGroupId),
		}

		dob, err := utils.ConvertTeamUnifyDate(member.Dob)
		if err != nil {
			return err
		}
		newMember.Dob = timestamppb.New(dob)

		dateJoined, err := utils.ConvertTeamUnifyDate(member.JoinedDate)
		if err != nil {
			return err
		}
		newMember.DateJoined = timestamppb.New(dateJoined)

		switch member.Sex {
		case 0:
			newMember.Sex = users.Member_FEMALE
		case 1:
			newMember.Sex = users.Member_MALE
		default:
			newMember.Sex = users.Member_OTHER
		}

		rsp.Members = append(rsp.Members, &newMember)
	}

	return nil
}

func (e *Users) WatchSwimmer(ctx context.Context, req *users.WatchSwimmerRequest, rsp *users.WatchSwimmerResponse) error {
	account, exists := auth.AccountFromContext(ctx)
	if !exists {
		return errors.Unauthorized("users.WatchSwimmer", "Unauthorized")
	}

	if req.Id == 0 {
		return errors.BadRequest("users.WatchSwimmer", "Swimmer id not provided")
	}

	accountId, err := strconv.Atoi(account.Metadata["tuId"])
	if err != nil {
		return err
	}

	_, err = e.DB.WatchSwimmer(ctx, &database.WatchSwimmerRequest{
		AccountId: int32(accountId),
		SwimmerId: req.Id,
	})
	if err != nil {
		return err
	}

	dbAccount, err := e.DB.GetAccount(ctx, &database.GetAccountRequest{
		Identifier: &database.GetAccountRequest_Id{Id: int32(accountId)},
	})
	if err != nil {
		return err
	}

	for _, watch := range dbAccount.Account.Watches {
		rsp.SwimmerIds = append(rsp.SwimmerIds, watch.Id)
	}

	return nil
}
