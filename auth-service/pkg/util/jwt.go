package util

import (
	"crypto/rsa"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Load RSA private key from file
func loadRSAPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// Create a JWT token with the user ID in the `sub` claim
func CreateJWTToken(userID int) (string, error) {
	// Get RSA private key from file
	privateKeyPath := GetEnv("PRIVATE_KEY_PATH", "")
	privateKey, err := loadRSAPrivateKey(privateKeyPath)
	if err != nil {
		return "", err
	}

	// Define standard claims with the user ID as the `sub`
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID),                                   // User ID stored in `sub` claim
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // Token expiration time
		IssuedAt:  jwt.NewNumericDate(time.Now()),                         // Token issuance time
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token using the RSA private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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
func ValidateJWTToken(tokenString string) (string, error) {
	// Get RSA public key from file
	publicKeyPath := GetEnv("PUBLIC_KEY_PATH", "")
	publicKey, err := loadRSAPublicKey(publicKeyPath)
	if err != nil {
		return "", err
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
		return "", err
	}

	// Extract registered claims
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("token is invalid")
	}

	// Return the user ID from the `sub` claim
	return claims.Subject, nil
}
