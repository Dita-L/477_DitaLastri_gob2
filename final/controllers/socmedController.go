package controllers

import (
	"final/database"
	"final/helpers"
	"final/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocMed(c *gin.Context) {
	//connect database
	db := database.GetDB()
	//get id user
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socmed := models.Socmed{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&socmed)
	} else {
		c.ShouldBind(&socmed)
	}
	socmed.User_ID = userID

	//create
	err := db.Debug().Create(&socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socmed.ID,
		"name":             socmed.Name,
		"social_media_url": socmed.Socmed_URL,
		"user_id":          socmed.User_ID,
		"created_at":       socmed.CreatedAt,
	})
}

func Updatesocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	socmed := models.Socmed{}

	socmedID, _ := strconv.Atoi(c.Param("socialmediaid"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&socmed)
	} else {
		c.ShouldBind(&socmed)
	}

	socmed.User_ID = userID
	socmed.ID = uint(socmedID)

	err := db.Model(&socmed).Where("id = ?", socmedID).Updates(models.Socmed{
		Name: socmed.Name, Socmed_URL: socmed.Socmed_URL},
	).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socmed.ID,
		"name":             socmed.Name,
		"social_media_url": socmed.Socmed_URL,
		"user_id":          socmed.User_ID,
		"updated_at":       socmed.UpdatedAt,
	})
}

func DeleteSocmed(c *gin.Context) {
	//connect database
	db := database.GetDB()

	//get socmedid dan userid
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	socmed := models.Socmed{}

	socmedID, _ := strconv.Atoi(c.Param("socialmediaid"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&socmed)
	} else {
		c.ShouldBind(&socmed)
	}

	socmed.User_ID = userID
	socmed.ID = uint(socmedID)

	//delete photo berdasarkan photoid
	err := db.Model(&socmed).Where("id = ?", socmedID).Delete(&socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
