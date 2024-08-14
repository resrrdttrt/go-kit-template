package admin

import (
	"context"
	"time"

	log "go-kit-template/pkg/logger"
)

type adminService struct {
	log        log.Logger
	users      UserRepository
	statistic  StatisticRepository
	auth       AuthRepository
}


type Service interface {
	userService
	statisticService
	authService
}

type userService interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserById(ctx context.Context, id string) (User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id string) error
	ActiveUser(ctx context.Context, id string) error
	DeactiveUser(ctx context.Context, id string) error
}

type statisticService interface {
	// General
	GetTotalUsers(ctx context.Context) (int, error)
	GetTotalGames(ctx context.Context) (int, error)

	GetTotalEnterprises(ctx context.Context) (int, error)
	GetTotalEndUser(ctx context.Context) (int, error)

	GetTotalActiveEndUsers(ctx context.Context) (int, error)
	GetTotalActiveEnterprises(ctx context.Context) (int, error)

	GetTotalNewEnterprisesInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error)
	GetTotalNewEndUsersInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error)

	GetTotalNewEndUsersInWeek(ctx context.Context) ([]Statistic, error)
	GetTotalNewEnterprisesInWeek(ctx context.Context) ([]Statistic, error)
}

type authService interface {
	// Auth
	Login(ctx context.Context, username, password string) (Token, error)
	GetUserIDByAccessToken(accessToken string) (string, error)
	GetUserRoleByID(userID string) (string, error)
}


func NewAdminService(log log.Logger, users UserRepository, statistic StatisticRepository, auth AuthRepository) Service {
	return &adminService{
		log:       log,
		users:     users,
		statistic: statistic,
		auth:      auth,
	}
}

func (s *adminService) GetAllUsers(ctx context.Context) ([]User, error) {
	return s.users.GetAllUsers(ctx)
}

func (s *adminService) GetUserById(ctx context.Context, id string) (User, error) {
	return s.users.GetUserById(ctx, id)
}

func (s *adminService) CreateUser(ctx context.Context, user User) error {
	return s.users.CreateUser(ctx, user)
}

func (s *adminService) UpdateUser(ctx context.Context, user User) error {
	return s.users.UpdateUser(ctx, user)
}

func (s *adminService) DeleteUser(ctx context.Context, id string) error {
	return s.users.DeleteUser(ctx, id)
}

func (s *adminService) ActiveUser(ctx context.Context, id string) error {
	user := User{
		ID:     id,
		Status: "active",
	}
	return s.users.UpdateUser(ctx, user)
}

func (s *adminService) DeactiveUser(ctx context.Context, id string) error {
	user := User{
		ID:     id,
		Status: "inactive",
	}
	return s.users.UpdateUser(ctx, user)
}


func (s *adminService) GetTotalUsers(ctx context.Context) (int, error) {
	return s.statistic.GetTotalUsers(ctx)
}

func (s *adminService) GetTotalGames(ctx context.Context) (int, error) {
	return s.statistic.GetTotalGames(ctx)
}

func (s *adminService) GetTotalEnterprises(ctx context.Context) (int, error) {
	return s.statistic.GetTotalEnterprises(ctx)
}

func (s *adminService) GetTotalEndUser(ctx context.Context) (int, error) {
	return s.statistic.GetTotalEndUser(ctx)
}

func (s *adminService) GetTotalActiveEndUsers(ctx context.Context) (int, error) {
	return s.statistic.GetTotalActiveEndUsers(ctx)
}

func (s *adminService) GetTotalActiveEnterprises(ctx context.Context) (int, error) {
	return s.statistic.GetTotalActiveEnterprises(ctx)
}

func (s *adminService) GetTotalNewEnterprisesInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error) {
	return s.statistic.GetTotalNewEnterprisesInTime(ctx, start, end)
}

func (s *adminService) GetTotalNewEndUsersInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error) {
	return s.statistic.GetTotalNewEndUsersInTime(ctx, start, end)
}

func (s *adminService) GetTotalNewEndUsersInWeek(ctx context.Context) ([]Statistic, error) {
	now := time.Now()
	start := now.AddDate(0, 0, -7)
	return s.statistic.GetTotalNewEndUsersInTime(ctx, start, now)
}

func (s *adminService) GetTotalNewEnterprisesInWeek(ctx context.Context) ([]Statistic, error) {
	now := time.Now()
	start := now.AddDate(0, 0, -7)
	return s.statistic.GetTotalNewEnterprisesInTime(ctx, start, now)
}

func (s *adminService) Login(ctx context.Context, username, password string) (Token, error) {
	return s.auth.Login(ctx, username, password)
}

func (s *adminService) GetUserIDByAccessToken(accessToken string) (string, error) {
	return s.auth.GetUserIDByAccessToken(accessToken)
}

func (s *adminService) GetUserRoleByID(userID string) (string, error) {
	return s.auth.GetUserRoleByID(userID)
}


type GRPCService interface {
	// GRPC
	Add(ctx context.Context, numA, numB float32)(sum float32,err error)
}

type grpcService struct {
	log log.Logger
	math MathRepository
}

func NewGRPCService(log log.Logger, math MathRepository) GRPCService {
	return &grpcService{
		log: log,
		math: math,
	}
}

func (s *grpcService) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return s.math.Add(ctx, numA, numB)
}
