package controllers

import (
    "net/http"
    "teste-enube/auth"
    "teste-enube/models"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type RegisterInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
    Message string `json:"message"`
    Token   string `json:"token,omitempty"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param register body RegisterInput true "Register"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Router /register [post]
func Register(c *gin.Context) {
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
        return
    }

    user := models.User{Username: input.Username, Password: input.Password}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&user)

    c.JSON(http.StatusOK, AuthResponse{Message: "Registration successful"})
}

// Login godoc
// @Summary Log in a user
// @Description Log in a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginInput true "Login"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
        return
    }

    var user models.User
    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid username or password"})
        return
    }

    if user.Password != input.Password {
        c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid username or password"})
        return
    }

    token, err := auth.GenerateToken(user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, AuthResponse{Message: "Login successful", Token: token})
}
