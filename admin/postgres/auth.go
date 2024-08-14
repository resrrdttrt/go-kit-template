package postgres

import (
	"context"
	"math/rand"
	"time"

	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/db"
	"github.com/resrrdttrt/VOU/pkg/errors"
	log "github.com/resrrdttrt/VOU/pkg/logger"
)

var _ admin.AuthRepository = (*authRepository)(nil)

type authRepository struct {
	db db.Database
	l  log.Logger
}

func NewAuthRepository(db db.Database, l log.Logger) admin.AuthRepository {
	return &authRepository{
		db: db,
		l:  l,
	}
}

func (r *authRepository) Login(ctx context.Context, username string, password string) (admin.Token, error) {
	query := `SELECT id, email, password FROM users WHERE username = :username`
	params := map[string]interface{}{
		"username": username,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return admin.Token{}, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var user admin.User
	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return admin.Token{}, errors.Wrap(ErrSelectDb, err)
		}
		if user.Password != password {
			return admin.Token{}, errors.New("password is incorrect")
		}
		// Generate random access token
		accessToken := generateAccessToken()

		// Add access token to access_token table by SQL
		insertQuery := `INSERT INTO access_tokens (user_id, token) VALUES (:user_id, :token)`
		insertParams := map[string]interface{}{
			"user_id": user.ID,
			"token":   accessToken,
		}
		_, err = r.db.NamedExecContext(ctx, insertQuery, insertParams)
		if err != nil {
			return admin.Token{}, errors.Wrap(ErrInsertDb, err)
		}

		// Return the generated access token
		return admin.Token{
			AccessToken: accessToken,
		}, nil
	} else {
		return admin.Token{}, errors.New("user not found")
	}
}

func generateAccessToken() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 32

	rand.Seed(time.Now().UnixNano())

	token := make([]byte, length)
	for i := 0; i < length; i++ {
		token[i] = charset[rand.Intn(len(charset))]
	}

	return string(token)
}


func (r *authRepository) GetUserIDByAccessToken(accessToken string) (string, error) {
	var userID string
	query := `SELECT user_id FROM access_tokens WHERE token = :token`
	params := map[string]interface{}{
		"token": accessToken,
	}
	ctx := context.Background()
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return "", errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&userID); err != nil {
			return "", errors.Wrap(ErrSelectDb, err)
		}
		return userID, nil
	} else {
		return "", errors.New("user not found")
	}
}

func (r *authRepository) GetUserRoleByID(userID string) (string, error) {
	var role string
	query := `SELECT role FROM users WHERE id = :id`
	params := map[string]interface{}{
		"id": userID,
	}
	ctx := context.Background()
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return "", errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&role); err != nil {
			return "", errors.Wrap(ErrSelectDb, err)
		}
		return role, nil
	} else {
		return "", errors.New("user not found")
	}
}

