package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func GetLocation(c *gin.Context) {
	c.JSON(http.StatusOK, models.MyResume.Basics.Location)
}

func PutLocation(c *gin.Context) {
	var location models.Location
	if err := utils.ReadFromBody(c.Request.Body, &location); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	models.MyResume.Basics.Location = location
}

func DeleteLocation(c *gin.Context) {
	location := models.MyResume.Basics.Location
	models.MyResume.Basics.Location = models.Location{}
	c.JSON(http.StatusOK, location)
}

func PatchLocation(c *gin.Context) {
	body := c.Request.Body
	var location map[string]interface{}
	if err := utils.ReadFromBody(body, &location); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	models.MyResume.Basics.Location.Patch(location)
	c.JSON(http.StatusOK, models.MyResume.Basics.Location)
}
