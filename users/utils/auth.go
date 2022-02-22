package utils

import (
	"context"
	"github.com/micro/micro/v3/proto/auth"
)

func CheckAccountExists(ctx context.Context, service auth.AccountsService, id string) (bool, error) {
	accounts, err := service.List(ctx, &auth.ListAccountsRequest{})
	if err != nil {
		return false, err
	}

	for _, account := range accounts.Accounts {
		if account.Id == id {
			return true, nil
		}
	}

	return false, nil
}
