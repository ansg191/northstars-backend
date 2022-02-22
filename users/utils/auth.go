package utils

import (
	"context"
	database "github.com/ansg191/northstars-backend/database/proto"
)

func CheckAccountExists(ctx context.Context, service database.DatabaseService, email string) (bool, error) {
	account, err := service.GetAccount(ctx, &database.GetAccountRequest{
		Identifier: &database.GetAccountRequest_Email{Email: email},
	})
	if err != nil {
		return false, err
	}

	return account.Account != nil, nil
}
