package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var works = models.MyResume.Work

func PostWork(c *gin.Context) {
	body := c.Request.Body
	var work models.Work
	if err := utils.ReadFromBody(body, &work); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if work.IfCompanyExists(works) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s company already exists", work.Company)})
		return
	}
	works = append(works, work)
	c.JSON(http.StatusOK, work)
}

func GetWork(c *gin.Context) {
	c.JSON(http.StatusOK, works)
}

func GetWorks(c *gin.Context) {
	network := c.Param("company")
	work := models.FindWork(works, network)
	if work.Company == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s profile not found", work.Company)})
		return
	}
	c.JSON(http.StatusOK, &work)
}
