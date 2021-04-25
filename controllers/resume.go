package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func GetResume(c *gin.Context) {
	c.JSON(http.StatusOK, models.MyResume)
}

func PutResume(c *gin.Context) {
	body := c.Request.Body
	var resume models.Resume
	if err := utils.ReadFromBody(body, &resume); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	models.MyResume = resume
	c.JSON(http.StatusOK, resume)
}

func DeleteResume(c *gin.Context) {
	resume := models.MyResume
	models.MyResume = models.Resume{}
	c.JSON(http.StatusOK, resume)
}

func PatchResume(c *gin.Context) {
	body := c.Request.Body
	var resume map[string]interface{}
	if err := utils.ReadFromBody(body, &resume); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	models.MyResume.Patch(resume)
	c.JSON(http.StatusOK, models.MyResume)
}
