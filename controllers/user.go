package controllers

import (
	"fmt"
	"net/http"

	models "goproject/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input models.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create user
	user := models.User{FirstName: input.FirstName, LastName: input.LastName, UserID: input.UserID, Password: input.Password}
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

// PATCH /users/:id
// Update a user
func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func LoginService(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	fmt.Println(db)
	var input models.LoginUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	fmt.Println(input.UserID, input.Password)
	//json.Unmarshal(input, &user)
	if err := db.Where("id = ?", input.ID).First(&user).Error; err != nil {
		fmt.Println(user)
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}
	if input.Password != user.Password {
		fmt.Println("This is in the password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})

}
