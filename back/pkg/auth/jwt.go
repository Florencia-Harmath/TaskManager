package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var mySecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateToken genera un nuevo token JWT
func GenerateToken(userID uuid.UUID) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(12 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(mySecretKey)
}

// ValidateToken valida un token JWT
 func ValidateToken(tokenString string) (*jwt.Token, error) {
     return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
         if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
             return nil, ErrInvalidSigningMethod
         }
         return mySecretKey, nil
     })
}

var ErrInvalidSigningMethod = errors.New("invalid signing method")
