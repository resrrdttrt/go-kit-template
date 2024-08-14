package postgres

import (
	"context"

	"go-kit-template/admin"
	"go-kit-template/pkg/db"
	"go-kit-template/pkg/errors"
	log "go-kit-template/pkg/logger"
)

var _ admin.UserRepository = (*usersRepository)(nil)

type usersRepository struct {
	db db.Database
	l  log.Logger
}

func NewUserRepository(db db.Database, l log.Logger) admin.UserRepository {
	return &usersRepository{
		db: db,
		l:  l,
	}
}

func (r *usersRepository) GetAllUsers(ctx context.Context) ([]admin.User, error) {
	query := `SELECT * FROM users`
	params := map[string]interface{}{}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return nil, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var users []admin.User
	for rows.Next() {
		var user admin.User
		if err := rows.StructScan(&user); err != nil {
			return nil, errors.Wrap(ErrSelectDb, err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *usersRepository) GetUserById(ctx context.Context, id string) (admin.User, error) {
	query := `SELECT * FROM users WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return admin.User{}, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var user admin.User
	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return admin.User{}, errors.Wrap(ErrSelectDb, err)
		}
		return user, nil
	} else {
		return admin.User{}, errors.Wrap(ErrNoData, err)
	}
}

func (r *usersRepository) CreateUser(ctx context.Context, user admin.User) error {
	query := `INSERT INTO users (name, username, password, email, phone, role, status) VALUES (:name, :username, :password, :email, :phone, :role, :status) RETURNING id`
	params := map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
		"phone":    user.Phone,
		"role":     user.Role,
		"status":   user.Status,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrInsertDb, err)
	}
	return nil
}

func (r *usersRepository) UpdateUser(ctx context.Context, user admin.User) error {
	query := `UPDATE users SET `
	params := map[string]interface{}{
		"id": user.ID,
	}

	if user.Name != "" {
		query += `name = :name, `
		params["name"] = user.Name
	}

	if user.Username != "" {
		query += `username = :username, `
		params["username"] = user.Username
	}

	if user.Password != "" {
		query += `password = :password, `
		params["password"] = user.Password
	}

	if user.Email != "" {
		query += `email = :email, `
		params["email"] = user.Email
	}

	if user.Phone != "" {
		query += `phone = :phone, `
		params["phone"] = user.Phone
	}

	if user.Role != "" {
		query += `role = :role, `
		params["role"] = user.Role
	}

	if user.Status != "" {
		query += `status = :status, `
		params["status"] = user.Status
	}

	query = query[:len(query)-2] + ` WHERE id = :id RETURNING *`
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrUpdateDb, err)
	}
	return nil
}

func (r *usersRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrDeleteDb, err)
	}
	return nil
}
