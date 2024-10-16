package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Tes-sudo/online-learning-platform/user-service/auth"
	"github.com/Tes-sudo/online-learning-platform/user-service/errors"
)

// AuthMiddleware is a middleware that checks for a valid JWT in the Authorization header
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")

		// The Authorization header should be in the format: "Bearer <token>"
		// Check if the header is empty or doesn't start with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			HandleError(w, errors.ErrUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			HandleError(w, errors.ErrUnauthorized)
			return
		}

		// If we get here, the token is valid. 
		// You can now use the claims, for example, add the username to the request context
		r = r.WithContext(context.WithValue(r.Context(), "username", claims.Username))

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}