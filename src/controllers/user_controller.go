package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkyar/rest-bookapp/src/config"
	"github.com/hokkyar/rest-bookapp/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := config.GetDB().Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.GetDB().First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": user})
}

func PostUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Enkripsi password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hashed password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := config.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create user success",
		"user":    user,
	})
}

func PutUser(c *gin.Context) {

	id := c.Param("id")
	var user models.User

	if err := config.GetDB().First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updatedData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
		FullName string `json:"full_name"`
	}

	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Enkripsi password
	if updatedData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = string(hashedPassword)
	}

	// Update fields lainnya
	user.Email = updatedData.Email
	user.Username = updatedData.Username
	user.FullName = updatedData.FullName

	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success update user",
		"user":    user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.GetDB().First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := config.GetDB().Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
