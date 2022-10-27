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

func CreateComment(c *gin.Context) {
	//connect database
	db := database.GetDB()

	//get id user
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	comment := models.Comment{}
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}
	comment.User_ID = userID

	//get id photo
	photo := models.Photo{}

	photoID, _ := strconv.Atoi(c.Request.FormValue("photo_id"))

	err := db.Model(&photo).Where("id = ?", photoID).Find(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	comment.Photo_ID = uint(photoID)

	//create comment
	err = db.Debug().Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.Photo_ID,
		"user_id":    comment.User_ID,
		"created_at": comment.CreatedAt,
	})

}

func GetAllComments(c *gin.Context) {
	//connect database
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	comment := []models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	err := db.Model(&models.Comment{}).Preload("Photo").Preload("User").Find(&comment).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func UpdateComment(c *gin.Context) {
	//connect to database
	db := database.GetDB()

	//get id user
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))

	comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}
	comment.User_ID = userID

	//update by commentid
	commentID, _ := strconv.Atoi(c.Param("commentid"))

	comment.User_ID = userID
	comment.ID = uint(commentID)

	err := db.Model(&comment).Where("id = ?", commentID).Updates(models.Comment{
		Message: comment.Message},
	).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	//get comment by commentid
	err = db.Model(&comment).Where("id = ?", commentID).Find(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"user_id":    comment.User_ID,
		"photo_id":   comment.Photo_ID,
		"created_at": comment.CreatedAt,
		"updated_at": comment.UpdatedAt,
	})

}

func DeleteComment(c *gin.Context) {
	//connect database
	db := database.GetDB()

	//get commentid dan userid
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	comment := models.Comment{}

	commentID, _ := strconv.Atoi(c.Param("commentid"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.User_ID = userID
	comment.ID = uint(commentID)

	//delete comment berdasarkan commentid
	err := db.Model(&comment).Where("id = ?", commentID).Delete(&comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
