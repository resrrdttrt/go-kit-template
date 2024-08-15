package http

import (
	"context"
	"encoding/json"
	"net/http"

	"go-kit-template/admin"
	"go-kit-template/middlewares"
	"go-kit-template/pkg/errors"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
)

func MakeAdminHandler(svc admin.Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := bone.New()

	r.Get("/user", kithttp.NewServer(
		getAllUsersEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/user/:id", kithttp.NewServer(
		getUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))
	r.Post("/user", kithttp.NewServer(
		createUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeResponse,
		opts...,
	))
	r.Put("/user/:id", kithttp.NewServer(
		updateUserEndpoint(svc),
		decodeUpdateUserRequest,
		encodeResponse,
		opts...,
	))
	r.Delete("/user/:id", kithttp.NewServer(
		deleteUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/user/active/:id", kithttp.NewServer(
		activeUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/user/deactive/:id", kithttp.NewServer(
		deactiveUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_users", kithttp.NewServer(
		getTotalUsersEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_games", kithttp.NewServer(
		getTotalGamesEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_enterprises", kithttp.NewServer(
		getTotalEnterprisesEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_end_users", kithttp.NewServer(
		getTotalEndUserEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_active_end_users", kithttp.NewServer(
		getTotalActiveEndUsersEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_active_enterprises", kithttp.NewServer(
		getTotalActiveEnterprisesEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_new_enterprises_in_time", kithttp.NewServer(
		getTotalNewEnterprisesInTimeEndpoint(svc),
		decodeStatisticInTimeRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_new_end_users_in_time", kithttp.NewServer(
		getTotalNewEndUsersInTimeEndpoint(svc),
		decodeStatisticInTimeRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_new_end_users_in_week", kithttp.NewServer(
		getTotalNewEndUsersInWeekEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/statistic/total_new_enterprises_in_week", kithttp.NewServer(
		getTotalNewEnterprisesInWeekEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	handler := middlewares.VerifyAdminMiddleware(r)
	return handler
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch errorVal := err.(type) {
	case errors.Error:
		switch {
		case errors.Contains(errorVal, errors.ErrNotFound):
			w.WriteHeader(http.StatusNotFound)
		case errors.Contains(errorVal, errors.ErrUnsupportedMediaType):
			w.WriteHeader(http.StatusUnsupportedMediaType)
		case errors.Contains(errorVal, errors.ErrMalformedEntity):
			w.WriteHeader(http.StatusBadRequest)
		case errors.Contains(errorVal, errors.ErrBadRequest):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		if errorVal.Msg() != "" {
			if err := json.NewEncoder(w).Encode(errorResponse{Message: errorVal.Msg(), Code: errorVal.Code(), Error: errorVal.Error()}); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
	}
	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req updateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
	}
	id := bone.GetValue(r, "id")
	req.ID = id
	return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getUserRequest
	id := bone.GetValue(r, "id")
	req.ID = id
	return req, nil
}

func decodeNothingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeStatisticInTimeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req statisticInTimeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
	}
	return req, nil
}


func MakeAuthHandler(svc admin.Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := bone.New()

	r.Post("/login", kithttp.NewServer(
		loginEndpoint(svc),
		decodeLoginRequest,
		encodeResponse,
		opts...,
	))
	return r
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
	}
	return req, nil
}



func MakeHandler (svc admin.Service) http.Handler {
	r := bone.New()
	adminHandler := MakeAdminHandler(svc)
	authHandler := MakeAuthHandler(svc)
	r.SubRoute("/admin", adminHandler)
	r.SubRoute("/auth", authHandler)
	return enableCORS(r)
}

func enableCORS(next http.Handler) http.Handler { // enable CORS for all route in handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func enableCORS(w http
