package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostEducation(c *gin.Context) {
	body := c.Request.Body
	var education m.Education
	if err := utils.ReadFromBody(body, &education); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if education.IfInstitutionExists(m.MyResume.Education) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s institution already exists", education.Institution)})
		return
	}
	m.MyResume.Education = append(m.MyResume.Education, education)
	c.JSON(http.StatusOK, education)
}

func GetEducations(c *gin.Context) {
	c.JSON(http.StatusOK, m.MyResume.Education)
}

func GetEducation(c *gin.Context) {
	institution := c.Param("institution")
	education, _ := m.FindEducation(m.MyResume.Education, institution)
	if education.Institution == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s institution not found", education.Institution)})
		return
	}
	c.JSON(http.StatusOK, &education)
}

func PutEducation(c *gin.Context) {
	institution := c.Param("institution")
	educationToUpdate, _ := m.FindEducation(m.MyResume.Education, institution)
	if educationToUpdate.Institution == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s institution not found", educationToUpdate.Institution)})
		return
	}
	body := c.Request.Body
	var education m.Education
	if err := utils.ReadFromBody(body, &education); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	educationToUpdate = &education
	c.JSON(http.StatusOK, education)
}

func PatchEducation(c *gin.Context) {
	defer utils.UpdateResume(m.MyResume.Education, m.MyResume.Education)
	institution := c.Param("institution")
	educationToUpdate, _ := m.FindEducation(m.MyResume.Education, institution)
	if educationToUpdate.Institution == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s institution not found", educationToUpdate.Institution)})
		return
	}
	body := c.Request.Body
	var education map[string]interface{}
	if err := utils.ReadFromBody(body, &education); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	educationToUpdate.Patch(education)
	c.JSON(http.StatusOK, &educationToUpdate)
}

func DeleteEducation(c *gin.Context) {
	defer utils.UpdateResume(m.MyResume.Education, m.MyResume.Education)
	institution := c.Param("institution")
	education, index := m.FindEducation(m.MyResume.Education, institution)
	if education.Institution == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s institution not found", education.Institution)})
		return
	}
	m.MyResume.Education = append(m.MyResume.Education[:index], m.MyResume.Education[index+1:]...)
	c.JSON(http.StatusOK, education)
}
