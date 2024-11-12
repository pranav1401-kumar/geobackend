package controllers

import (
    "GeoDataApp/models"
    "GeoDataApp/utils"
    "encoding/json"
    "net/http"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// Register endpoint
func Register(w http.ResponseWriter, r *http.Request) {
    var user models.Credentials
    json.NewDecoder(r.Body).Decode(&user)

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)

    if err := utils.DB.Create(&user).Error; err != nil {
        http.Error(w, "Unable to register user", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode("Registration successful")
}

// Login endpoint
func Login(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&credentials)

    var user models.Credentials
    utils.DB.Where("email = ?", credentials.Email).First(&user)

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := token.SignedString(jwtKey)

    http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   tokenString,
        Expires: expirationTime,
    })

    json.NewEncoder(w).Encode("Login successful")
}
