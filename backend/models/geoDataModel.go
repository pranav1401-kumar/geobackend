package models

import "gorm.io/gorm"

// GeoData represents geospatial data, including GeoJSON or KML attributes
type GeoData struct {
    gorm.Model
    UserID uint   `json:"user_id"`
    Data   string `json:"data"` // Store GeoJSON or KML data as a string
}
