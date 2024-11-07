package auth

import "net/http"

type AuthController interface {
	Status(w http.Response, r http.Request)
	CreateUser(w http.Response, r http.Request)
	CreateAuthenticationToken(w http.Response, r http.Request)
	Protected(w http.Response, r http.Request)
}
