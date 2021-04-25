package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostPublication(c *gin.Context) {
	body := c.Request.Body
	var publication m.Publication
	if err := utils.ReadFromBody(body, &publication); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if publication.IfNameExists(m.MyResume.Publications) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s name already exists", publication.Name)})
		return
	}
	m.MyResume.Publications = append(m.MyResume.Publications, publication)
	c.JSON(http.StatusOK, publication)
}

func GetPublications(c *gin.Context) {
	c.JSON(http.StatusOK, m.MyResume.Publications)
}

func GetPublication(c *gin.Context) {
	name := c.Param("name")
	publication, _ := m.FindPublication(m.MyResume.Publications, name)
	if publication.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", publication.Name)})
		return
	}
	c.JSON(http.StatusOK, &publication)
}

func PutPublication(c *gin.Context) {
	name := c.Param("name")
	publicationToUpdate, index := m.FindPublication(m.MyResume.Publications, name)
	if publicationToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", publicationToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var publication m.Publication
	if err := utils.ReadFromBody(body, &publication); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	publicationToUpdate = &publication
	m.MyResume.Publications[index] = *publicationToUpdate
	c.JSON(http.StatusOK, publication)
}

func PatchPublication(c *gin.Context) {
	name := c.Param("name")
	publicationToUpdate, index := m.FindPublication(m.MyResume.Publications, name)
	if publicationToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", publicationToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var publication map[string]interface{}
	if err := utils.ReadFromBody(body, &publication); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	publicationToUpdate.Patch(publication)
	m.MyResume.Publications[index] = *publicationToUpdate
	c.JSON(http.StatusOK, &publicationToUpdate)
}

func DeletePublication(c *gin.Context) {
	name := c.Param("name")
	publication, index := m.FindPublication(m.MyResume.Publications, name)
	if publication.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", publication.Name)})
		return
	}
	m.MyResume.Publications = append(m.MyResume.Publications[:index], m.MyResume.Publications[index+1:]...)
	c.JSON(http.StatusOK, publication)
}
