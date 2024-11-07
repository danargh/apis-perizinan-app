package auth

import (
	"github.com/danargh/apis-perizinan-app/pkg/database"
)

type AuthRepository interface {
	InsertUser(db *database.DB, input AuthRegisterRequest) (AuthRegisterRequest, error)
}
