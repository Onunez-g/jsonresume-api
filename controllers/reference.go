package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var references = models.MyResume.References

func PostReference(c *gin.Context) {
	body := c.Request.Body
	var reference models.Reference
	if err := utils.ReadFromBody(body, &reference); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if reference.IfNameExists(references) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s name already exists", reference.Name)})
		return
	}
	references = append(references, reference)
	c.JSON(http.StatusOK, reference)
}

func GetReferences(c *gin.Context) {
	c.JSON(http.StatusOK, references)
}

func GetReference(c *gin.Context) {
	name := c.Param("name")
	reference, _ := models.FindReference(references, name)
	if reference.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", reference.Name)})
		return
	}
	c.JSON(http.StatusOK, &reference)
}

func PutReference(c *gin.Context) {
	name := c.Param("name")
	referenceToUpdate, _ := models.FindReference(references, name)
	if referenceToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", referenceToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var reference models.Reference
	if err := utils.ReadFromBody(body, &reference); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if reference.IfNameExists(references) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s name already exists", reference.Name)})
		return
	}
	referenceToUpdate = &reference
	c.JSON(http.StatusOK, reference)
}

func PatchReference(c *gin.Context) {
	name := c.Param("name")
	publicationToUpdate, _ := models.FindReference(references, name)
	if publicationToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", publicationToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var reference map[string]interface{}
	if err := utils.ReadFromBody(body, &reference); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	publicationToUpdate.Patch(reference)
	c.JSON(http.StatusOK, &publicationToUpdate)
}

func DeleteReference(c *gin.Context) {
	name := c.Param("name")
	reference, index := models.FindReference(references, name)
	if reference.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", reference.Name)})
		return
	}
	references = append(references[:index], references[index+1:]...)
	c.JSON(http.StatusOK, reference)
}
