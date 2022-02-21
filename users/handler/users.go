package handler

import (
	"context"
	cookiestealer "github.com/ansg191/northstars-backend/cookie-stealer/proto"

	log "github.com/micro/micro/v3/service/logger"

	users "github.com/ansg191/northstars-backend/users/proto"
)

type Users struct {
	cookies cookiestealer.CookieStealerService
}

// NewUser is a single request handler called via client.Call or the generated client code
func (e *Users) NewUser(ctx context.Context, req *users.NewUserRequest, rsp *users.NewUserResponse) error {
	log.Info("Received Users.NewUser request")
	return nil
}
