package login

import (
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

// Admin represents the Admin model
type Admin struct {
	ID         uint   `gorm:"primaryKey"`
	Email      string `gorm:"uniqueIndex"`
	Password   string // Note: In production, store hashed passwords, not plain text.
	IsVerified bool   `gorm:"default:false"`
	Role       string `gorm:"admin"`
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
		logService.Errorln("error" + "Invalid email domain. Only " + h.c.Domain + " emails are allowed.")
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
	}

	if err := h.service.PostgesQL.Create(admin, &adminInfo).Error; err != nil {
		logService.Errorln("error : Failed to create admin")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "Admin registered successfully! Please check your email to verify your account."})

}
