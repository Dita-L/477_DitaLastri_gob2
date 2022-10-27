package controllers

import (
	"final/database"
	"final/helpers"
	"final/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	//"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// type respGetUser struct {
// 	username string //`json:"username"`
// 	email    string //`json:"email"`
// }

// type respGet struct {
// 	id         uint      //`json:"photo_id"`
// 	title      string    //`json:"title"`
// 	caption    string    //`json:"caption"`
// 	photo_url  string    //`json:"photo_url"`
// 	user_id    uint      //`json:"user_id"`
// 	created_at time.Time //`json:"created_at"`
// 	updated_at time.Time //`json:"updated_at"`
// 	//user       respGetUser //`json:"user"`
// }

func CreatePhoto(c *gin.Context) {
	//connect database
	db := database.GetDB()
	//get id user
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.User_ID = userID

	//upload single photo
	file, _ := c.FormFile("photo_url")
	log.Println(file.Filename)

	// Upload the file to assets/photo.
	c.SaveUploadedFile(file, "assets/photo"+file.Filename)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	photo.Photo_URL = file.Filename

	//create photo
	err := db.Debug().Create(&photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.Photo_URL,
		"user_id":    photo.User_ID,
		"created_at": photo.CreatedAt,
	})

}

func GetAllPhoto(c *gin.Context) {
	//connect database
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	photo := []models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	err := db.Model(&models.Photo{}).Preload("User").Find(&photo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	// for i := 0; i < len(photo); i++ {
	// 	dataResp := respGet{
	// 		id:         photo[i].ID,
	// 		title:      photo[i].Title,
	// 		caption:    photo[i].Caption,
	// 		photo_url:  photo[i].Photo_URL,
	// 		user_id:    photo[i].User_ID,
	// 		created_at: photo[i].CreatedAt,
	// 		updated_at: photo[i].UpdatedAt,
	// 	}
	// 	a = append(a, dataResp)
	// }

	c.JSON(http.StatusOK, photo)
}

// func abc(photo *[]models.Photo) []respGet {
// 	var a []respGet
// 	for _, v := range *photo {
// 		dataResp := respGet{
// 			id:         v.ID,
// 			title:      v.Title,
// 			caption:    v.Caption,
// 			photo_url:  v.Photo_URL,
// 			user_id:    v.User_ID,
// 			created_at: v.CreatedAt,
// 			updated_at: v.UpdatedAt,
// 		}
// 		a = append(a, dataResp)
// 	}
// 	return a
// }

func Updatephoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	photo := models.Photo{}

	photoID, _ := strconv.Atoi(c.Param("photoid"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.User_ID = userID
	photo.ID = uint(photoID)

	err := db.Model(&photo).Where("id = ?", photoID).Updates(models.Photo{
		Caption: photo.Caption, Title: photo.Title, Photo_URL: photo.Photo_URL},
	).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.Photo_URL,
		"user_id":    photo.User_ID,
		"created_at": photo.CreatedAt,
		"updated_at": photo.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	//connect database
	db := database.GetDB()

	//get photoid dan userid
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	photo := models.Photo{}

	photoID, _ := strconv.Atoi(c.Param("photoid"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.User_ID = userID
	photo.ID = uint(photoID)

	//delete photo berdasarkan photoid
	err := db.Model(&photo).Where("id = ?", photoID).Delete(&photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
