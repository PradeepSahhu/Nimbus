package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func CheckValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email) && len(email) <= 254
}

func Hash(password string, cost int) (string, error) {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		return "", fmt.Errorf("invalid bcrypt cost %d, must be between %d and %d",
			cost, bcrypt.MinCost, bcrypt.MaxCost)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// JWT
func GenerateToken(userID string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "c869821be627c0c49ba13002a47f31ddacf52bdaf7bcea57f415c5788b119bf48273c1725b456530f211498679bed5c0644d21a86511293ccc34c621c7c0cddef51be0c261043d4962d8bb0d3457d4c66918017b0e20bedfd7fbb6f02b6ee96c927695b8845ae8905f0c6aa9b61b1acc4c7da17ea9cffb90d498179e024b7b72"
	}

	expirationTime := time.Now().Add(14 * 24 * time.Hour) // 14 days

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "nimbus",
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "c869821be627c0c49ba13002a47f31ddacf52bdaf7bcea57f415c5788b119bf48273c1725b456530f211498679bed5c0644d21a86511293ccc34c621c7c0cddef51be0c261043d4962d8bb0d3457d4c66918017b0e20bedfd7fbb6f02b6ee96c927695b8845ae8905f0c6aa9b61b1acc4c7da17ea9cffb90d498179e024b7b72" // Default secret
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", jwt.ErrSignatureInvalid
	}

	return claims.UserID, nil
}
