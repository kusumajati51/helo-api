package controllers

import (
	jwttoken "example/hello-api/auth/jwt_token"
	authdto "example/hello-api/dto/auth_dto"
	usersdto "example/hello-api/dto/users_dto"
	"example/hello-api/helper"
	"example/hello-api/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	db.Find(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	var user_input usersdto.CreateUserInputDTO

	if err := c.ShouldBindJSON(&user_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: user_input.Name, Email: &user_input.Email, No_HP: user_input.No_HP, Password: helper.GeneratePassword(user_input.Password)}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UserLogin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var login_input authdto.UserLoginDTO
	if err := c.ShouldBindJSON(&login_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	db.Where("email = ?", login_input.Email).First(&user)
	credentialError := helper.ComparePassword(user.Password, login_input.Password)
	if credentialError {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		c.Abort()
		return
	}
	token, err := jwttoken.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user, "token": token})

}

func GetDataUserLogin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	user_id, err := jwttoken.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := db.First(&user, user_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
