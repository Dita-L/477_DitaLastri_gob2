package controller

import (
	"ass-03/model"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateStatus(ctx *gin.Context) {
	var status model.Status

	if err := ctx.ShouldBindJSON(&status); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//random seed
	rand.Seed(time.Now().UnixNano())
	//initialize random number
	status.Water = rand.Intn(10) + 1
	status.Wind = rand.Intn(20) + 1

	ctx.JSON(http.StatusOK, gin.H{
		"water": status.Water,
		"wind":  status.Wind,
	})

}

func GetNewStatus(ctx *gin.Context) {
	var status model.Status

	//random seed
	rand.Seed(time.Now().UnixNano())
	//initialize random number
	status.Water = rand.Intn(10) + 1
	status.Wind = rand.Intn(20) + 1

	ctx.JSON(http.StatusOK, gin.H{
		"water": status.Water,
		"wind":  status.Wind,
	})
}
