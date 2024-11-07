package auth

import (
	"net/http"

	"github.com/danargh/apis-perizinan-app/internal"
	"github.com/danargh/apis-perizinan-app/pkg/password"
	"github.com/danargh/apis-perizinan-app/pkg/request"
	"github.com/danargh/apis-perizinan-app/pkg/response"
	"github.com/go-playground/validator/v10"
)

type AuthControllerImpl struct {
	AuthService AuthService
}

func NewAuthController(authService AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (c AuthControllerImpl) Status(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"Status": "HELLO WORLD! I LOVE GOLANG",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		// app.serverError(w, r, err)
		internal.PanicIfError(err)
	}
}

func (c AuthControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email     string              `json:"Email"`
		Password  string              `json:"Password"`
		Validator validator.Validator `json:"-"`
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

// func (c AuthControllerImpl) CreateAuthenticationToken(w http.ResponseWriter, r *http.Request) {
// 	var input struct {
// 		Email     string              `json:"Email"`
// 		Password  string              `json:"Password"`
// 		Validator validator.Validator `json:"-"`
// 	}

// 	err := request.DecodeJSON(w, r, &input)
// 	if err != nil {
// 		app.badRequest(w, r, err)
// 		return
// 	}

// 	user, found, err := app.db.GetUserByEmail(input.Email)
// 	if err != nil {
// 		app.serverError(w, r, err)
// 		return
// 	}

// 	input.Validator.CheckField(input.Email != "", "Email", "Email is required")
// 	input.Validator.CheckField(found, "Email", "Email address could not be found")

// 	if found {
// 		passwordMatches, err := password.Matches(input.Password, user.HashedPassword)
// 		if err != nil {
// 			app.serverError(w, r, err)
// 			return
// 		}

// 		input.Validator.CheckField(input.Password != "", "Password", "Password is required")
// 		input.Validator.CheckField(passwordMatches, "Password", "Password is incorrect")
// 	}

// 	if input.Validator.HasErrors() {
// 		app.failedValidation(w, r, input.Validator)
// 		return
// 	}

// 	// create jwt token
// 	var claims jwt.Claims
// 	claims.Subject = strconv.Itoa(user.ID)
// 	claims.Issued = jwt.NewNumericTime(time.Now())    // waktu token diterbitkan
// 	claims.NotBefore = jwt.NewNumericTime(time.Now()) // token ini tidak valid sebelum waktu tertentu
// 	expiry := time.Now().Add(24 * time.Hour)          // waktu expired 1 hari
// 	claims.Expires = jwt.NewNumericTime(expiry)
// 	claims.Issuer = app.config.baseURL
// 	claims.Audiences = []string{app.config.baseURL}

// 	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secretKey))
// 	if err != nil {
// 		app.serverError(w, r, err)
// 		return
// 	}

// 	data := map[string]string{
// 		"AuthenticationToken":       string(jwtBytes),
// 		"AuthenticationTokenExpiry": expiry.Format(time.RFC3339),
// 	}

// 	err = response.JSON(w, http.StatusOK, data)
// 	if err != nil {
// 		app.serverError(w, r, err)
// 	}
// }

// func (c AuthControllerImpl) Protected(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("This is a protected handler"))
// }
