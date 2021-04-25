package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var educations = models.MyResume.Education

func PostEducation(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Education, educations)
	body := c.Request.Body
	var education models.Education
	if err := utils.ReadFromBody(body, &education); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if education.IfInstitutionExists(educations) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s institution already exists", education.Institution)})
		return
	}
	educations = append(educations, education)
	c.JSON(http.StatusOK, education)
}

func GetEducations(c *gin.Context) {
	c.JSON(http.StatusOK, educations)
}

func GetEducation(c *gin.Context) {
	institution := c.Param("institution")
	education, _ := models.FindEducation(educations, institution)
	if education.Institution == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s institution not found", education.Institution)})
		return
	}
	c.JSON(http.StatusOK, &education)
}

func PutEducation(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Education, educations)
	institution := c.Param("institution")
	educationToUpdate, _ := models.FindEducation(educations, institution)
	if educationToUpdate.Institution == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s institution not found", educationToUpdate.Institution)})
		return
	}
	body := c.Request.Body
	var education models.Education
	if err := utils.ReadFromBody(body, &education); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if education.IfInstitutionExists(educations) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s institution already exists", education.Institution)})
		return
	}
	educationToUpdate = &education
	c.JSON(http.StatusOK, education)
}

func PatchEducation(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Education, educations)
	institution := c.Param("institution")
	educationToUpdate, _ := models.FindEducation(educations, institution)
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
	defer utils.UpdateResume(models.MyResume.Education, educations)
	institution := c.Param("institution")
	education, index := models.FindEducation(educations, institution)
	if education.Institution == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s institution not found", education.Institution)})
		return
	}
	educations = append(educations[:index], educations[index+1:]...)
	c.JSON(http.StatusOK, education)
}
