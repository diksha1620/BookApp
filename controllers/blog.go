package controllers

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	// "github.com/books/handlers"
	"github.com/books/models"
	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {

	var existingBlog models.Blog
	claims := jwt.ExtractClaims(c)
	user_email, _ := claims["email"]
	var User models.User
	var blog models.Blog

	// Check if the current user had admin role.
	if err := models.DB.Where("email = ? AND user_role_id=3", user_email).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blog can only be added by user"})
		return
	}

	//Check if the blog title already exists.
	err := models.DB.Where("title = ?").First(&existingBlog).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "blog title already exists, please! choose different Title ."})
		return
	}
	if err := c.ShouldBindJSON(&existingBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DB.Create(&blog).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         blog.ID,
		"name":       blog.Title,
		"body":       blog.Body,
		"created_by": blog.CreatedBy,
	})

}
