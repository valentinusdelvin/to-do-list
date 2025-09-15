package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JWTItf interface {
	GenerateToken(userId, role string) (string, error)
	ValidateToken(token string) (string, string, error)
}

type JWT struct {
	secretKey   string
	expiredTime time.Time
}

type Claims struct {
	userId string
	role   string
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	env := godotenv.Load()
	if env != nil {
		panic("Error loading .env file")
	}

	secretKey := os.Getenv("JWT_SECRET")
	expiredTime := time.Now().Add(2 * time.Hour)

	return &JWT{
		secretKey:   secretKey,
		expiredTime: expiredTime,
	}
}

func (j *JWT) GenerateToken(userId, role string) (string, error) {
	claims := Claims{
		userId: userId,
		role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(j.expiredTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *JWT) ValidateToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return "", "", err
	}

	if !token.Valid {
		return "", "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", "", err
	}

	return claims.userId, claims.role, nil
}
