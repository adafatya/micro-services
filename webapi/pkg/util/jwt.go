package util

import (
	"crypto/rsa"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

// Load RSA public key from file
func loadRSAPublicKey(path string) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

// Validate JWT token and parse the user ID from `sub` claim
func GetUserID(tokenString string) (int, error) {
	// Get RSA public key from file
	publicKeyPath := GetEnv("PUBLIC_KEY_PATH", "")
	publicKey, err := loadRSAPublicKey(publicKeyPath)
	if err != nil {
		return 0, err
	}

	// Parse the token and extract registered claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is RSA
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return 0, err
	}

	// Extract registered claims
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("token is invalid")
	}

	// Return the user ID from the `sub` claim
	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, fmt.Errorf("subject is invalid")
	}
	return userID, nil
}
