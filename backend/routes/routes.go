package routes

import (
    "GeoDataApp/controllers"
    "GeoDataApp/middleware"
    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/auth/register", controllers.Register).Methods("POST")
    router.HandleFunc("/auth/login", controllers.Login).Methods("POST")

    // Protected Routes
    api := router.PathPrefix("/api").Subrouter()
    api.Use(middleware.AuthMiddleware)
    api.HandleFunc("/geo/upload", controllers.UploadGeoFile).Methods("POST")
    api.HandleFunc("/geo/data", controllers.GetGeoData).Methods("GET")
    api.HandleFunc("/geo/data/{id}", controllers.UpdateGeoData).Methods("PUT")
    
    return router
}
