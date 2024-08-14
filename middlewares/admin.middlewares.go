package middlewares

import (
	"context"
	"net/http"

	"go-kit-template/admin"
)

func VerifyRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		userID, err := admin.GetUserIDByAccessToken(accessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		role, err := admin.GetUserRoleByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		type contextKey string
		const userIDKey contextKey = "userID"
		const roleKey contextKey = "role"
		
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		ctx = context.WithValue(ctx, roleKey, role)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}


func VerifyAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		userID, err := admin.GetUserIDByAccessToken(accessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		role, err := admin.GetUserRoleByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if role != "admin" {
			http.Error(w, "You are not authorized to access this resource", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)

	})
}
