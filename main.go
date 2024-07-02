package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    _ "teste-enube/docs"
    "teste-enube/controllers"
    "teste-enube/routes"
    "teste-enube/auth"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    "teste-enube/models"
)

// @title Excel Data API
// @version 1.0
// @description This API serves data from an Excel sheet with JWT authentication and SQLite3 for user management.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.bearerAuth Bearer
// @in header
// @name Authorization
// @BearerFormat JWT
// @description JWT Authorization header using the Bearer scheme. Example: "Bearer {token}"
func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize database
    db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrate the schema
    db.AutoMigrate(&models.User{})

    router := gin.Default()

    // Middleware to set db in context
    router.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    router.POST("/login", controllers.Login)
    router.POST("/register", controllers.Register)

    // Use authentication middleware
    authenticated := router.Group("/")
    authenticated.Use(auth.AuthMiddleware())
    {
        routes.SetupDataRoutes(authenticated)
    }

    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    router.Run(":" + port)
}
