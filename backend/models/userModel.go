package models

import (
    "gorm.io/gorm"
)

type Credentials struct {
    gorm.Model
    Username string `json:"username" gorm:"unique"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}