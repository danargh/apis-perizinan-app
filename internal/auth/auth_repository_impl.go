package auth

import (
	"context"
	"time"

	"github.com/danargh/apis-perizinan-app/pkg/database"
)

type User struct {
	ID             int       `db:"id"`
	Created        time.Time `db:"created"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
}

type AuthRepositoryImpl struct{}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (r *AuthRepositoryImpl) InsertUser(db *database.DB, input AuthRegisterRequest) (AuthRegisterRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // ini aslinya menggunakan default time out
	defer cancel()

	var resData AuthRegisterRequest

	query := `
		INSERT INTO users (created, email, hashed_password)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := db.GetContext(ctx, &resData, query, time.Now(), input.Email, input.Password)
	if err != nil {
		return AuthRegisterRequest{}, err
	}

	return resData, err
}

// func (r *AuthRepositoryImpl) GetUser(db *database.DB, id int) (*User, bool, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
// 	defer cancel()

// 	var user User

// 	query := `SELECT * FROM users WHERE id = $1`

// 	err := db.GetContext(ctx, &user, query, id)
// 	if errors.Is(err, sql.ErrNoRows) {
// 		return nil, false, nil
// 	}

// 	return &user, true, err
// }

// func (r *AuthRepositoryImpl) GetUserByEmail(db *database.DB, email string) (*User, bool, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
// 	defer cancel()

// 	var user User

// 	query := `SELECT * FROM users WHERE email = $1`

// 	err := db.GetContext(ctx, &user, query, email)
// 	if errors.Is(err, sql.ErrNoRows) {
// 		return nil, false, nil
// 	}

// 	return &user, true, err
// }

// func (r *AuthRepositoryImpl) UpdateUserHashedPassword(db *database.DB, id int, hashedPassword string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
// 	defer cancel()

// 	query := `UPDATE users SET hashed_password = $1 WHERE id = $2`

// 	_, err := db.ExecContext(ctx, query, hashedPassword, id)
// 	return err
// }
