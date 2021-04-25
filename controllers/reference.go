package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostReference(c *gin.Context) {
	body := c.Request.Body
	var reference m.Reference
	if err := utils.ReadFromBody(body, &reference); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if reference.IfNameExists(m.MyResume.References) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s name already exists", reference.Name)})
		return
	}
	m.MyResume.References = append(m.MyResume.References, reference)
	c.JSON(http.StatusOK, reference)
}

func GetReferences(c *gin.Context) {
	c.JSON(http.StatusOK, m.MyResume.References)
}

func GetReference(c *gin.Context) {
	name := c.Param("name")
	reference, _ := m.FindReference(m.MyResume.References, name)
	if reference.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", reference.Name)})
		return
	}
	c.JSON(http.StatusOK, &reference)
}

func PutReference(c *gin.Context) {
	name := c.Param("name")
	referenceToUpdate, index := m.FindReference(m.MyResume.References, name)
	if referenceToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", referenceToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var reference m.Reference
	if err := utils.ReadFromBody(body, &reference); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	referenceToUpdate = &reference
	m.MyResume.References[index] = *referenceToUpdate
	c.JSON(http.StatusOK, reference)
}

func PatchReference(c *gin.Context) {
	name := c.Param("name")
	referenceToUpdate, index := m.FindReference(m.MyResume.References, name)
	if referenceToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", referenceToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var reference map[string]interface{}
	if err := utils.ReadFromBody(body, &reference); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	referenceToUpdate.Patch(reference)
	m.MyResume.References[index] = *referenceToUpdate

	c.JSON(http.StatusOK, &referenceToUpdate)
}

func DeleteReference(c *gin.Context) {
	name := c.Param("name")
	reference, index := m.FindReference(m.MyResume.References, name)
	if reference.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", reference.Name)})
		return
	}
	m.MyResume.References = append(m.MyResume.References[:index], m.MyResume.References[index+1:]...)
	c.JSON(http.StatusOK, reference)
}
