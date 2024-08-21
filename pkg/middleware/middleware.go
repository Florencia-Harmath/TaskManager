package middleware

import (
	"net/http"
	"strings"
	"taskmanager/pkg/auth"
)

func JwtAuthentication(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Authorization header required", http.StatusUnauthorized)
            return
        }

        // Eliminar el prefijo "Bearer " si está presente
        token = strings.TrimPrefix(token, "Bearer ")

        _, err := auth.ValidateToken(token)
        if err != nil {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
