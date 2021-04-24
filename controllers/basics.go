package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func GetBasics(c *gin.Context) {
	c.JSON(http.StatusOK, models.MyResume.Basics)
}

func PutBasics(c *gin.Context) {
	body := c.Request.Body
	var basics models.Basics
	if err := utils.ReadFromBody(body, &basics); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	models.MyResume.Basics = basics
	c.JSON(http.StatusOK, basics)
}

func DeleteBasics(c *gin.Context) {
	basics := models.MyResume.Basics
	models.MyResume.Basics = models.Basics{}
	c.JSON(http.StatusOK, basics)
}

func PatchBasics(c *gin.Context) {
	body := c.Request.Body
	var basics map[string]interface{}
	if err := utils.ReadFromBody(body, &basics); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	models.MyResume.Basics.Patch(basics)
	c.JSON(http.StatusOK, models.MyResume.Basics)
}
