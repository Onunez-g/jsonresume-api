package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostWork(c *gin.Context) {

	body := c.Request.Body
	var work m.Work
	if err := utils.ReadFromBody(body, &work); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if work.IfCompanyExists(m.MyResume.Work) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s company already exists", work.Company)})
		return
	}
	m.MyResume.Work = append(m.MyResume.Work, work)
	c.JSON(http.StatusOK, work)
}

func GetWorks(c *gin.Context) {
	c.JSON(http.StatusOK, m.MyResume.Work)
}

func GetWork(c *gin.Context) {
	company := c.Param("company")
	work, _ := m.FindWork(m.MyResume.Work, company)
	if work.Company == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s company not found", work.Company)})
		return
	}
	c.JSON(http.StatusOK, &work)
}

func PutWork(c *gin.Context) {

	company := c.Param("company")
	workToUpdate, _ := m.FindWork(m.MyResume.Work, company)
	if workToUpdate.Company == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s company not found", workToUpdate.Company)})
		return
	}
	body := c.Request.Body
	var work m.Work
	if err := utils.ReadFromBody(body, &work); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	workToUpdate = &work
	c.JSON(http.StatusOK, work)
}

func PatchWork(c *gin.Context) {

	company := c.Param("company")
	workToUpdate, _ := m.FindWork(m.MyResume.Work, company)
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

	company := c.Param("company")
	work, index := m.FindWork(m.MyResume.Work, company)
	if work.Company == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s company not found", work.Company)})
		return
	}
	m.MyResume.Work = append(m.MyResume.Work[:index], m.MyResume.Work[index+1:]...)
	c.JSON(http.StatusOK, work)
}
