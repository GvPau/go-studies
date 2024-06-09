package middlewares

import (
	"net/http"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		expectedToken := "API_TOKEN" // Replace with your actual token

		// Perform case-insensitive comparison
		if !strings.EqualFold(token, "Bearer "+expectedToken) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
