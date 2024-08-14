package http

import (
	"context"

	"go-kit-template/admin"
	"go-kit-template/pkg/common"

	"github.com/go-kit/kit/endpoint"
)

func getAllUsersEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := svc.GetAllUsers(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(users), nil
	}
}

func createUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		user := admin.User{
			Name:     req.Name,
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
			Phone:    req.Phone,
			Role:     req.Role,
			Status:   req.Status,
		}
		if err := svc.CreateUser(ctx, user); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func updateUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		user := admin.User{
			ID:       req.ID,
			Name:     req.Name,
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
			Phone:    req.Phone,
			Role:     req.Role,
			Status:   req.Status,
		}
		if err := svc.UpdateUser(ctx, user); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func getUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		user, err := svc.GetUserById(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(user), nil
	}
}

func deleteUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		if err := svc.DeleteUser(ctx, req.ID); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func activeUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		if err := svc.ActiveUser(ctx, req.ID); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func deactiveUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		if err := svc.DeactiveUser(ctx, req.ID); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}




func getTotalUsersEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := svc.GetTotalUsers(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(users), nil
	}
}

func getTotalGamesEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		games, err := svc.GetTotalGames(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(games), nil
	}
}

func getTotalEnterprisesEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		enterprises, err := svc.GetTotalEnterprises(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(enterprises), nil
	}
}

func getTotalEndUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		endusers, err := svc.GetTotalEndUser(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(endusers), nil
	}
}

func getTotalActiveEndUsersEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		endusers, err := svc.GetTotalActiveEndUsers(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(endusers), nil
	}
}

func getTotalActiveEnterprisesEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		enterprises, err := svc.GetTotalActiveEnterprises(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(enterprises), nil
	}
}

func getTotalNewEnterprisesInTimeEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var req statisticInTimeRequest
		if err := req.validate(); err != nil {
			return nil, err
		}
		enterprises, err := svc.GetTotalNewEnterprisesInTime(ctx, req.Start, req.End)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(enterprises), nil
	}
}

func getTotalNewEndUsersInTimeEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var req statisticInTimeRequest
		if err := req.validate(); err != nil {
			return nil, err
		}
		endusers, err := svc.GetTotalNewEndUsersInTime(ctx, req.Start, req.End)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(endusers), nil
	}
}

func getTotalNewEndUsersInWeekEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		endusers, err := svc.GetTotalNewEndUsersInWeek(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(endusers), nil
	}
}

func getTotalNewEnterprisesInWeekEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		enterprises, err := svc.GetTotalNewEnterprisesInWeek(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(enterprises), nil
	}
}

func loginEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		token, err := svc.Login(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(token), nil
	}
}
