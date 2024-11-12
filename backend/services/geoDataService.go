package services

import (
    "GeoDataApp/models"
    "GeoDataApp/utils"
)

// SaveGeoData saves new geospatial data to the database
func SaveGeoData(geoData *models.GeoData) error {
    return utils.DB.Create(geoData).Error
}

// GetGeoDataByUser retrieves geospatial data by user ID
func GetGeoDataByUser(userID uint) ([]models.GeoData, error) {
    var geoData []models.GeoData
    err := utils.DB.Where("user_id = ?", userID).Find(&geoData).Error
    return geoData, err
}

// UpdateGeoData updates an existing geospatial data entry in the database
func UpdateGeoData(geoData *models.GeoData) error {
    return utils.DB.Save(geoData).Error
}
