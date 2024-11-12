package services

import (
    "GeoDataApp/models"
    "GeoDataApp/utils"
    "errors"
    "time"

    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key") // Replace with a secure, environment-loaded secret

// Claims struct for JWT payload
type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// RegisterUser registers a new user with a hashed password and saves it to the database
func RegisterUser(user *models.Credentials) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    
    // Save the user to the database
    if err := utils.DB.Create(user).Error; err != nil {
        return errors.New("unable to register user")
    }
    return nil
}

// LoginUser authenticates a user by email and password, and generates a JWT token if successful
func LoginUser(creds models.Credentials) (string, error) {
    var user models.Credentials
    
    // Find user by email
    if err := utils.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
        return "", errors.New("user not found")
    }

    // Compare stored hashed password with provided password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    // Create JWT token with expiration time
    expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
    claims := &Claims{
        UserID: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // Create token with claims and sign it
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateToken parses and validates the provided JWT token string
func ValidateToken(tokenStr string) (*Claims, error) {
    claims := &Claims{}
    
    // Parse token with the provided claims
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, errors.New("invalid token")
    }

    // Check if token is valid
    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
