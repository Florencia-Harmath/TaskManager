// middleware.go

package middleware

import (
	"context"
	"net/http"
	"strings"
	"taskmanager/pkg/auth"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Define un tipo clave personalizado
type ContextKey string

const UserIDKey ContextKey = "userID" // Cambiado a mayúscula para que sea exportado

func JwtAuthentication(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Authorization header required", http.StatusUnauthorized)
            return
        }

        // Eliminar el prefijo "Bearer " si está presente
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Validar el token y extraer los claims
        token, err := auth.ValidateToken(tokenString)
        if err != nil {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Extraer el userID de los claims
        userID, err := uuid.Parse(claims["user_id"].(string))
        if err != nil {
            http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
            return
        }

        // Colocar el userID en el contexto usando la clave personalizada
        ctx := context.WithValue(r.Context(), UserIDKey, userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
