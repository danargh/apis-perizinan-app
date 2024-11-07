package auth

import (
	"context"
	"net/http"

	"github.com/danargh/apis-perizinan-app/pkg/database"
	"github.com/danargh/apis-perizinan-app/pkg/password"
	"github.com/danargh/apis-perizinan-app/pkg/response"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository AuthRepository
	db             *database.DB
	Validate       *validator.Validate
}

func NewAuthService(authRepository AuthRepository, db *database.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		db:             db,
		Validate:       validate,
	}
}

func (s AuthServiceImpl) CreateUser(ctx context.Context, request AuthRegisterRequest) AuthRegisterResponse {
	var input struct {
		Email     string             `json:"Email"`
		Password  string             `json:"Password"`
		Validator validator.Validate `json:"-"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	_, found, err := app.db.GetUserByEmail(input.Email)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	input.Validator.CheckField(input.Email != "", "Email", "Email is required")
	input.Validator.CheckField(validator.Matches(input.Email, validator.RgxEmail), "Email", "Must be a valid email address")
	input.Validator.CheckField(!found, "Email", "Email is already in use")

	input.Validator.CheckField(input.Password != "", "Password", "Password is required")
	input.Validator.CheckField(len(input.Password) >= 8, "Password", "Password is too short")
	input.Validator.CheckField(len(input.Password) <= 72, "Password", "Password is too long")
	input.Validator.CheckField(validator.NotIn(input.Password, password.CommonPasswords...), "Password", "Password is too common")

	if input.Validator.HasErrors() {
		app.failedValidation(w, r, input.Validator)
		return
	}

	hashedPassword, err := password.Hash(input.Password)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	_, err = app.db.InsertUser(input.Email, hashedPassword)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := map[string]interface{}{
		"Status":  "Success",
		"Message": "User created successfully",
		"Data": map[string]interface{}{
			"email": input.Email,
		},
	}

	// w.WriteHeader(http.StatusNoContent)
	err = response.JSON(w, http.StatusOK, res)
	if err != nil {
		app.serverError(w, r, err)
	}
}
