package handler

import (
	"context"
	"fmt"
	database "github.com/ansg191/northstars-backend/database/proto"
	"strconv"

	authProto "github.com/micro/micro/v3/proto/auth"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
	"github.com/micro/micro/v3/service/store/client"
	"google.golang.org/protobuf/types/known/timestamppb"

	cookiestealer "github.com/ansg191/northstars-backend/cookie-stealer/proto"
	users "github.com/ansg191/northstars-backend/users/proto"
	"github.com/ansg191/northstars-backend/users/utils"
)

type Users struct {
	Cookies  cookiestealer.CookieStealerService
	Accounts authProto.AccountsService
	DB       database.DatabaseService
}

// NewUser is a single request handler called via client.Call or the generated client code
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
	})
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

	watchStore := client.NewStore(
		store.WithContext(ctx),
		store.Table("watch-table"),
	)

	oldRecord, err := store.Read(account.ID)
	if err.Error() == "not found" {
	} else if err != nil {
		return err
	}

	var watchedSwimmers []int
	if len(oldRecord) > 0 {
		err = oldRecord[0].Decode(&watchedSwimmers)
		if err != nil {
			return err
		}
	}

	for _, swimmer := range watchedSwimmers {
		if swimmer == int(req.Id) {
			for _, swimmer2 := range watchedSwimmers {
				rsp.SwimmerIds = append(rsp.SwimmerIds, int32(swimmer2))
			}
			return nil
		}
	}

	watchedSwimmers = append(watchedSwimmers, int(req.Id))

	newRecord := store.NewRecord(account.ID, watchedSwimmers)
	err = watchStore.Write(newRecord)
	if err != nil {
		return err
	}

	for _, swimmer := range watchedSwimmers {
		rsp.SwimmerIds = append(rsp.SwimmerIds, int32(swimmer))
	}

	return nil
}
