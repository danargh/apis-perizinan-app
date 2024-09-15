package main

import (
	"context"
	"net/http"

	"github.com/danargh/apis-perizinan-app/internal/database"
)

type contextKey string

const (
	// kunci untuk context
	authenticatedUserContextKey = contextKey("authenticatedUser")
)

// menyimpan data pengguna dengan kunci ke context
func contextSetAuthenticatedUser(r *http.Request, user *database.User) *http.Request {
	ctx := context.WithValue(r.Context(), authenticatedUserContextKey, user)
	return r.WithContext(ctx)
}

// mengambil kembali informasi yang disimpan di context
func contextGetAuthenticatedUser(r *http.Request) *database.User {
	user, ok := r.Context().Value(authenticatedUserContextKey).(*database.User)
	if !ok {
		return nil
	}

	return user
}
