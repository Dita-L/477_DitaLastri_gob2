package controllers

import (
	"final/database"
	"final/helpers"
	"final/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	//connect database
	db := database.GetDB()

	//menginput data user dari request body
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	//response ketika berhasil
	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})

}

func UserLogin(c *gin.Context) {
	//connect database
	db := database.GetDB()

	//mengambil inputan data dari request body
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	//cek email
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	//cek password
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	//response token
	token := helpers.GenerateToken(uint(User.ID), User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

// update user
func UpdateUser(c *gin.Context) {
	//connect to database
	db := database.GetDB()

	//get user_id, jwt
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User.ID = userID

	//menginput data update-an
	err := db.Model(&User).Where("id = ?", userID).Updates(models.User{Username: User.Username, Email: User.Email, Age: User.Age}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        userID,
		"email":     User.Email,
		"user_name": User.Username,
		"age":       User.Age,
	})
}

// delete
func DeleteUser(c *gin.Context) {
	//connect database
	db := database.GetDB()

	//get id
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User.ID = userID

	//delete user
	err := db.Where("id = ?", userID).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Your account has been sucessfully deleted",
	})
}
