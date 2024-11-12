package main

import (
    "github.com/gin-gonic/gin"
    "github.com/rs/cors"
    "log"
    "net/http"
)

func main() {
    // Create a Gin router with default middleware
    r := gin.Default()

    // Create a CORS configuration
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"https://geofrontend.vercel.app/"}, // Frontend URL
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    // Use the CORS middleware with Gin
    r.Use(func(ctx *gin.Context) {
        handler := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx.Next()
        }))
        handler.ServeHTTP(ctx.Writer, ctx.Request)
    })

    // Example route: Registration
    r.POST("/auth/register", func(c *gin.Context) {
        var json struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        // Simulate a registration process
        c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
    })

    // Example route: Login
    r.POST("/auth/login", func(c *gin.Context) {
        var json struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        // Simulate a login process
        // Normally you'd verify credentials here
        c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": "some-jwt-token"})
    })

    // Example route: Get User Profile (requires authentication)
    r.GET("/profile", func(c *gin.Context) {
        // Simulate retrieving user profile (authentication logic should be here)
        c.JSON(http.StatusOK, gin.H{"username": "john_doe", "email": "john@example.com"})
    })

    // Example route: Update User Profile (requires authentication)
    r.PUT("/profile", func(c *gin.Context) {
        var json struct {
            Username string `json:"username"`
            Email    string `json:"email"`
        }
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        // Simulate updating user profile
        c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
    })

    // Example route: Get List of GeoData
    r.GET("/geodata", func(c *gin.Context) {
        geoData := []map[string]interface{}{
            {"id": 1, "location": "Location A", "coordinates": "51.505, -0.09"},
            {"id": 2, "location": "Location B", "coordinates": "51.515, -0.10"},
        }
        c.JSON(http.StatusOK, geoData)
    })

    // Example route: Add GeoData
    r.POST("/geodata", func(c *gin.Context) {
        var json struct {
            Location    string `json:"location"`
            Coordinates string `json:"coordinates"`
        }
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Geo data added successfully"})
    })

    // Example route: Delete GeoData
    r.DELETE("/geodata/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{"message": "Geo data with ID " + id + " deleted successfully"})
    })

    // Example route: Update GeoData
    r.PUT("/geodata/:id", func(c *gin.Context) {
        id := c.Param("id")
        var json struct {
            Location    string `json:"location"`
            Coordinates string `json:"coordinates"`
        }
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Geo data with ID " + id + " updated successfully"})
    })

    // New route: Handle Geo Data Upload
    r.POST("/geo/upload", func(c *gin.Context) {
        // Get the file from the form data
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "No file received"})
            return
        }

        // Save the file to the local disk
        if err := c.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
            return
        }

        // Respond with success
        c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": file.Filename})
    })

    // Start the server on port 8080
    log.Fatal(r.Run(":8080"))
}
