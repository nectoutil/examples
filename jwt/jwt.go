package main

import (
    "fmt"
    "time"

    "npkg.dev/crypto/jwt.v4"
)

// Secret key for signing tokens
var secretKey = []byte("your-secret-key-here")

// Custom claims struct
type Claims struct {
    UserID uint   `json:"user_id"`
    Email  string `json:"email"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token
func GenerateToken(userID uint, email, role string) (string, error) {
    // Create claims with user data and expiration time
    claims := Claims{
        UserID: userID,
        Email:  email,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    // Create token with claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with the secret key
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateToken verifies and parses the JWT token
func ValidateToken(tokenString string) (*Claims, error) {
    // Parse the token
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        // Validate signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    // Extract claims
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("invalid token")
}

func main() {
    // Example usage
    // Generate a token
    tokenString, err := GenerateToken(123, "user@example.com", "admin")
    if err != nil {
        fmt.Printf("Error generating token: %v\n", err)
        return
    }
    fmt.Printf("Generated Token: %v\n\n", tokenString)

    // Validate the token
    claims, err := ValidateToken(tokenString)
    if err != nil {
        fmt.Printf("Error validating token: %v\n", err)
        return
    }

    fmt.Println("Token is valid!")
    fmt.Printf("User ID: %d\n", claims.UserID)
    fmt.Printf("Email: %s\n", claims.Email)
    fmt.Printf("Role: %s\n", claims.Role)
    fmt.Printf("Expires At: %v\n", claims.ExpiresAt)
}
package main

import (
    "fmt"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

// Secret key for signing tokens
var secretKey = []byte("your-secret-key-here")

// Custom claims struct
type Claims struct {
    UserID uint   `json:"user_id"`
    Email  string `json:"email"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token
func GenerateToken(userID uint, email, role string) (string, error) {
    // Create claims with user data and expiration time
    claims := Claims{
        UserID: userID,
        Email:  email,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    // Create token with claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with the secret key
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateToken verifies and parses the JWT token
func ValidateToken(tokenString string) (*Claims, error) {
    // Parse the token
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        // Validate signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    // Extract claims
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("invalid token")
}

func main() {
    // Example usage
    // Generate a token
    tokenString, err := GenerateToken(123, "user@example.com", "admin")
    if err != nil {
        fmt.Printf("Error generating token: %v\n", err)
        return
    }
    fmt.Printf("Generated Token: %v\n\n", tokenString)

    // Validate the token
    claims, err := ValidateToken(tokenString)
    if err != nil {
        fmt.Printf("Error validating token: %v\n", err)
        return
    }

    fmt.Println("Token is valid!")
    fmt.Printf("User ID: %d\n", claims.UserID)
    fmt.Printf("Email: %s\n", claims.Email)
    fmt.Printf("Role: %s\n", claims.Role)
    fmt.Printf("Expires At: %v\n", claims.ExpiresAt)
}
