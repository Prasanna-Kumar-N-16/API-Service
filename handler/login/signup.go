package login

import (
	"api-service/encryption"
	"api-service/logger"
	"api-service/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	admin = "admin"
)

type AdminSignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

/*
CREATE TABLE admins (
    id SERIAL PRIMARY KEY,          -- ID field, auto-incremented
    email VARCHAR(100) UNIQUE NOT NULL, -- Email field, unique and not null
    password VARCHAR(255) NOT NULL,     -- Password field, not null
    is_verified BOOLEAN DEFAULT FALSE,  -- IsVerified field, default value of false
    role VARCHAR(50) NOT NULL,         -- Role field, not null
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- CreatedAt field with default current timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP -- UpdatedAt field with default current timestamp
);

*/

// Define the struct with GORM tags
type Admin struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`          // Primary key and auto-increment
	Email      string    `gorm:"type:varchar(100);unique;not null"` // Unique and not null
	Password   string    `gorm:"type:varchar(255);not null"`        // Not null
	IsVerified bool      `gorm:"default:false"`                     // Default value of false
	Role       string    `gorm:"type:varchar(50);not null"`         // Not null
	CreatedAt  time.Time // Automatically managed by GORM
	UpdatedAt  time.Time // Automatically managed by GORM
}

func (h *Authenticationhandler) Signup(c *gin.Context) {
	var req AdminSignupRequest
	logService := logger.GetLogger()
	if err := c.ShouldBindJSON(&req); err != nil {
		logService.Errorln("error in Signup API request body parsing reason:", err)
		utils.APIResponse(c, reqBodyParseErr, http.StatusBadRequest, nil)
		return
	}

	if req.Email == "" || req.Password == "" {
		logService.Errorln("error" + "Invalid email or password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Ensure the database connection is not nil
	if h.service.PostgesQL == nil {
		logService.Errorln("error: Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Auto migrate the Admin struct to create/update the database table
	err := h.service.PostgesQL.DB.AutoMigrate(&Admin{})
	if err != nil {
		logService.Errorln("failed to migrate database", err)
		return
	}

	// Validate email domain
	if !utils.IsAdminEmail(req.Email, h.c.Domain) {
		logService.Errorln("error " + "Invalid email domain. Only " + h.c.Domain + " emails are allowed.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email domain. Only " + h.c.Domain + " emails are allowed."})
		return
	}

	// Create the admin record in the database
	adminInfo := Admin{
		Email:      req.Email,
		Password:   req.Password,
		IsVerified: false,
		Role:       admin,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	encryptedPassword, err := encryption.Encrypt(adminInfo.Password, h.c.EncryptKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}
	adminInfo.Password = encryptedPassword

	if err := h.service.PostgesQL.Create(&adminInfo).Error; err != nil {
		logService.Errorln("error : Failed to create admin reason", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "Admin registered successfully! Please check your email to verify your account."})

}
