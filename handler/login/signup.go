package login

import (
	"api-service/logger"
	"api-service/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminSignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// Admin represents the Admin model
type Admin struct {
	ID         uint   `gorm:"primaryKey"`
	Email      string `gorm:"uniqueIndex"`
	Password   string // Note: In production, store hashed passwords, not plain text.
	IsVerified bool   `gorm:"default:false"`
	CreatedAt  time.Time
}

func (h *Authenticationhandler) Signup(c *gin.Context) {
	var req AdminSignupRequest
	logService := logger.GetLogger()
	if err := c.ShouldBindJSON(&req); err != nil {
		logService.Errorln("error in Signup API request body parsing reason:", err)
		utils.APIResponse(c, reqBodyParseErr, http.StatusBadRequest, nil)
		return
	}

	// Validate email domain
	if !utils.IsAdminEmail(req.Email, h.c.Domain) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email domain. Only example.com emails are allowed."})
		return
	}

	// Create the admin record in the database
	_ = Admin{
		Email:      req.Email,
		Password:   req.Password,
		IsVerified: false,
		CreatedAt:  time.Now(),
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "Admin registered successfully! Please check your email to verify your account."})

}
