package utils

import (
	database "github.com/ansg191/northstars-backend/database/proto"
	"github.com/gotidy/ptr"
	"github.com/nyaruka/phonenumbers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPhoneNumber(pnString string) (*phonenumbers.PhoneNumber, error) {
	return phonenumbers.Parse(pnString, "US")
}

func FormatPhoneNumbers(pnString string) (string, error) {
	phoneNumber, err := ToPhoneNumber(pnString)
	if err != nil {
		return "", err
	}
	return phonenumbers.Format(phoneNumber, phonenumbers.E164), nil
}

func ConvertSwimmerToProto(swimmer *Swimmer) *database.Swimmer {
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
		protoSwimmer.Watchers = append(
			protoSwimmer.Watchers,
			ConvertAccountToProto(&watcher),
		)
	}

	return &protoSwimmer
}

func ConvertAccountToProto(account *Account) *database.Account {
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

	for _, swimmer := range account.WatchedSwimmers {
		protoAccount.Watches = append(
			protoAccount.Watches,
			ConvertSwimmerToProto(&swimmer),
		)
	}

	return &protoAccount
}
