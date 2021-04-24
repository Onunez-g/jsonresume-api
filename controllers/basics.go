package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostBasics(c *gin.Context) {
	body := c.Request.Body
	var basics models.Basics
	if err := utils.ReadFromBody(body, &basics); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	models.MyResume.Basics = basics
	c.JSON(http.StatusCreated, basics)
}

func GetBasics(c *gin.Context) {
	c.JSON(http.StatusOK, models.MyResume.Basics)
}

func PutBasics(c *gin.Context) {
	body := c.Request.Body
	var basics models.Basics
	if err := utils.ReadFromBody(body, &basics); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	models.MyResume.Basics = basics
	c.JSON(http.StatusOK, basics)
}

func DeleteBasics(c *gin.Context) {
	models.MyResume.Basics = models.Basics{}
	c.Status(http.StatusOK)
}

func PatchBasics(c *gin.Context) {
	body := c.Request.Body
	var basics map[string]interface{}
	if err := utils.ReadFromBody(body, &basics); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	models.MyResume.Basics.Update(basics)
	c.JSON(http.StatusOK, models.MyResume.Basics)
}
