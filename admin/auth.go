package admin

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/resrrdttrt/VOU/pkg/common"
)

var DB *sql.DB

func ConnectToPostgres() {
	// Đọc các biến môi trường hoặc sử dụng giá trị mặc định
	host := common.Env("DB_HOST", "localhost")
	port := common.Env("DB_PORT", "5432")
	user := common.Env("DB_USER", "postgres")
	password := common.Env("DB_PASS", "1")
	dbname := common.Env("DB_NAME", "admin")

	// Tạo chuỗi kết nối DSN từ biến môi trường hoặc giá trị mặc định
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	// Kết nối tới cơ sở dữ liệu PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	DB = db
}

func GetUserIDByAccessToken(accessToken string) (string, error) {
	var userID string
	query := `SELECT user_id FROM access_tokens WHERE token = $1`

	err := DB.QueryRow(query, accessToken).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// handle no rows returned case
			return "", fmt.Errorf("no user found with the given access token")
		}
		return "", err
	}

	return userID, nil
}

func GetUserRoleByID(userID string) (string, error) {
	var role string
	query := `SELECT role FROM users WHERE id = $1`

	err := DB.QueryRow(query, userID).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no role found for the given user ID")
		}
		return "", err
	}

	return role, nil
}

func GetUserIDByEventID(eventID string) (string, error) {
	var userID string
	query := `SELECT user_id FROM events WHERE id = $1`

	err := DB.QueryRow(query, eventID).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found for the given event ID")
		}
		return "", err
	}

	return userID, nil
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type AuthRepository interface {
	Login(ctx context.Context, username, password string) (Token, error)
	GetUserIDByAccessToken(accessToken string) (string, error)
	GetUserRoleByID(userID string) (string, error)
}
