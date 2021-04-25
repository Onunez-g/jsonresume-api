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
	defer utils.UpdateResume(models.MyResume.Work, works)
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

func GetWorks(c *gin.Context) {
	c.JSON(http.StatusOK, works)
}

func GetWork(c *gin.Context) {
	company := c.Param("company")
	work, _ := models.FindWork(works, company)
	if work.Company == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s company not found", work.Company)})
		return
	}
	c.JSON(http.StatusOK, &work)
}

func PutWork(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Work, works)
	company := c.Param("company")
	workToUpdate, _ := models.FindWork(works, company)
	if workToUpdate.Company == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s company not found", workToUpdate.Company)})
		return
	}
	body := c.Request.Body
	var work models.Work
	if err := utils.ReadFromBody(body, &work); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if work.IfCompanyExists(works) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s profile already exists", work.Company)})
		return
	}
	workToUpdate = &work
	c.JSON(http.StatusOK, work)
}

func PatchWork(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Work, works)
	company := c.Param("company")
	workToUpdate, _ := models.FindWork(works, company)
	if workToUpdate.Company == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s company not found", workToUpdate.Company)})
		return
	}
	body := c.Request.Body
	var work map[string]interface{}
	if err := utils.ReadFromBody(body, &work); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	workToUpdate.Patch(work)
	c.JSON(http.StatusOK, &workToUpdate)
}

func DeleteWork(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Work, works)
	company := c.Param("company")
	work, index := models.FindWork(works, company)
	if work.Company == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s company not found", work.Company)})
		return
	}
	works = append(works[:index], works[index+1:]...)
	c.JSON(http.StatusOK, work)
}
