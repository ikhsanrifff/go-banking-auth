package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/ikhsanrifff/go-banking-auth/config"
)

// AuthMiddleware untuk validasi JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// remove "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := config.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// get context from request then set data
		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", claims.ID)
		ctx = context.WithValue(ctx, "username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
