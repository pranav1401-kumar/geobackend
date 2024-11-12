package controllers

import (
	"GeoDataApp/models"
	"GeoDataApp/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UploadGeoFile(w http.ResponseWriter, r *http.Request) {
    file, _, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "File upload failed", http.StatusBadRequest)
        return
    }
    defer file.Close()

    fileData, _ := ioutil.ReadAll(file)
    var geoData map[string]interface{}
    if err := json.Unmarshal(fileData, &geoData); err != nil {
        http.Error(w, "Invalid file format", http.StatusBadRequest)
        return
    }

    // Save file data to DB or further processing
    json.NewEncoder(w).Encode("File uploaded successfully")
}

func GetGeoData(w http.ResponseWriter, r *http.Request) {
    // Retrieve and return GeoData based on user authentication
    json.NewEncoder(w).Encode("GeoData retrieval logic here")
}

func UpdateGeoData(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    geoDataID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid geospatial data ID", http.StatusBadRequest)
        return
    }

    var updatedGeoData models.GeoData
    if err := json.NewDecoder(r.Body).Decode(&updatedGeoData); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    updatedGeoData.ID = uint(geoDataID)
    if err := services.UpdateGeoData(&updatedGeoData); err != nil {
        http.Error(w, "Failed to update geospatial data", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode("GeoData updated successfully")
}