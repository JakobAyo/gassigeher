package services

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestAuthService_HashPassword(t *testing.T) {
	service := NewAuthService("test-secret", 24)

	password := "TestPassword123"
	hash, err := service.HashPassword(password)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if hash == "" {
		t.Error("Expected hash to be generated")
	}

	if hash == password {
		t.Error("Expected hash to be different from password")
	}
}

func TestAuthService_CheckPassword(t *testing.T) {
	service := NewAuthService("test-secret", 24)

	password := "TestPassword123"
	hash, _ := service.HashPassword(password)

	// Test correct password
	if !service.CheckPassword(password, hash) {
		t.Error("Expected password to match hash")
	}

	// Test incorrect password
	if service.CheckPassword("WrongPassword", hash) {
		t.Error("Expected incorrect password to not match")
	}
}

func TestAuthService_GenerateToken(t *testing.T) {
	service := NewAuthService("test-secret", 24)

	token1, err := service.GenerateToken()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(token1) != 64 { // 32 bytes = 64 hex characters
		t.Errorf("Expected token length 64, got %d", len(token1))
	}

	// Test uniqueness
	token2, _ := service.GenerateToken()
	if token1 == token2 {
		t.Error("Expected tokens to be unique")
	}
}

func TestAuthService_GenerateJWT(t *testing.T) {
	service := NewAuthService("test-secret", 24)

	tokenString, err := service.GenerateJWT(1, "test@example.com", false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if tokenString == "" {
		t.Error("Expected JWT token to be generated")
	}

	// Parse and verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("test-secret"), nil
	})

	if err != nil {
		t.Errorf("Expected token to be valid, got %v", err)
	}

	if !token.Valid {
		t.Error("Expected token to be valid")
	}

	// Check claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatal("Expected claims to be MapClaims")
	}

	if claims["user_id"].(float64) != 1 {
		t.Error("Expected user_id to be 1")
	}

	if claims["email"].(string) != "test@example.com" {
		t.Error("Expected email to be test@example.com")
	}

	if claims["is_admin"].(bool) != false {
		t.Error("Expected is_admin to be false")
	}
}

func TestAuthService_ValidateJWT(t *testing.T) {
	service := NewAuthService("test-secret", 24)

	// Generate valid token
	tokenString, _ := service.GenerateJWT(1, "test@example.com", true)

	// Validate token
	claims, err := service.ValidateJWT(tokenString)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if (*claims)["user_id"].(float64) != 1 {
		t.Error("Expected user_id to be 1")
	}

	if (*claims)["is_admin"].(bool) != true {
		t.Error("Expected is_admin to be true")
	}

	// Test invalid token
	_, err = service.ValidateJWT("invalid-token")
	if err == nil {
		t.Error("Expected error for invalid token")
	}
}

func TestAuthService_ValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"Valid password", "Test123Pass", false},
		{"Too short", "Test1", true},
		{"No uppercase", "test123pass", true},
		{"No lowercase", "TEST123PASS", true},
		{"No number", "TestPassword", true},
		{"Valid complex", "MyP@ssw0rd", false},
	}

	service := NewAuthService("test-secret", 24)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.ValidatePassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthService_JWTExpiration(t *testing.T) {
	// Create service with 0 hour expiration for testing
	service := &AuthService{
		jwtSecret:           "test-secret",
		jwtExpirationHours: 0,
	}

	// Generate token that expires immediately
	tokenString, _ := service.GenerateJWT(1, "test@example.com", false)

	// Wait a moment
	time.Sleep(1 * time.Second)

	// Try to validate - should fail due to expiration
	_, err := service.ValidateJWT(tokenString)
	if err == nil {
		t.Error("Expected error for expired token")
	}
}
