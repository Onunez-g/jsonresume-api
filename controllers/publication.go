package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var publications = models.MyResume.Publications

func PostPublication(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Publications, publications)
	body := c.Request.Body
	var publication models.Publication
	if err := utils.ReadFromBody(body, &publication); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if publication.IfNameExists(publications) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s name already exists", publication.Name)})
		return
	}
	publications = append(publications, publication)
	c.JSON(http.StatusOK, publication)
}

func GetPublications(c *gin.Context) {
	c.JSON(http.StatusOK, publications)
}

func GetPublication(c *gin.Context) {
	name := c.Param("name")
	publication, _ := models.FindPublication(publications, name)
	if publication.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", publication.Name)})
		return
	}
	c.JSON(http.StatusOK, &publication)
}

func PutPublication(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Publications, publications)
	name := c.Param("name")
	publicationToUpdate, _ := models.FindPublication(publications, name)
	if publicationToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", publicationToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var publication models.Publication
	if err := utils.ReadFromBody(body, &publication); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if publication.IfNameExists(publications) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s name already exists", publication.Name)})
		return
	}
	publicationToUpdate = &publication
	c.JSON(http.StatusOK, publication)
}

func PatchPublication(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Publications, publications)
	name := c.Param("name")
	publicationToUpdate, _ := models.FindPublication(publications, name)
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
	c.JSON(http.StatusOK, &publicationToUpdate)
}

func DeletePublication(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Publications, publications)
	name := c.Param("name")
	publication, index := models.FindPublication(publications, name)
	if publication.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", publication.Name)})
		return
	}
	publications = append(publications[:index], publications[index+1:]...)
	c.JSON(http.StatusOK, publication)
}
