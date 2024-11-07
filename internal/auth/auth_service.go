package auth

import "context"

type AuthService interface {
	CreateUser(ctx context.Context, request AuthRegisterRequest) AuthRegisterResponse
}
