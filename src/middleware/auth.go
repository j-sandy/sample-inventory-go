package middleware

import (
	"net/http"
	"sample-inventory-go/store"
	"sample-inventory-go/utils"
	"strings"
)

func AuthMiddleware(next http.Handler, s *store.InventoryStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		token := tokenParts[1]
		_, ok := s.GetSession(token)
		if !ok {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired session token")
			return
		}

		// If authentication is successful, pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
