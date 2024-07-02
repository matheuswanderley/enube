package auth

import (
    "testing"
    "time"

    "github.com/dgrijalva/jwt-go"
)

func TestGenerateToken(t *testing.T) {
    username := "testuser"
    token, err := GenerateToken(username)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    claims := &Claims{}
    parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !parsedToken.Valid {
        t.Fatalf("Expected valid token, got %v", err)
    }

    if claims.Username != username {
        t.Fatalf("Expected username %v, got %v", username, claims.Username)
    }
}

func TestValidateToken(t *testing.T) {
    username := "testuser"
    token, err := GenerateToken(username)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    claims, err := ValidateToken(token)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if claims.Username != username {
        t.Fatalf("Expected username %v, got %v", username, claims.Username)
    }
}
