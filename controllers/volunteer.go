package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var volunteers = models.MyResume.Volunteer

func PostVolunteer(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Volunteer, volunteers)
	body := c.Request.Body
	var volunteer models.Work
	if err := utils.ReadFromBody(body, &volunteer); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if volunteer.IfOrganizationExists(volunteers) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s organization already exists", volunteer.Organization)})
		return
	}
	volunteers = append(volunteers, volunteer)
	c.JSON(http.StatusOK, volunteer)
}

func GetVolunteers(c *gin.Context) {
	c.JSON(http.StatusOK, volunteers)
}

func GetVolunteer(c *gin.Context) {
	organization := c.Param("organization")
	volunteer, _ := models.FindVolunteer(volunteers, organization)
	if volunteer.Organization == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s organization not found", volunteer.Organization)})
		return
	}
	c.JSON(http.StatusOK, &volunteer)
}

func PutVolunteer(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Volunteer, volunteers)
	organization := c.Param("organization")
	volunteerToUpdate, _ := models.FindVolunteer(volunteers, organization)
	if volunteerToUpdate.Organization == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s organization not found", volunteerToUpdate.Organization)})
		return
	}
	body := c.Request.Body
	var volunteer models.Work
	if err := utils.ReadFromBody(body, &volunteer); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if volunteer.IfOrganizationExists(volunteers) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s organization already exists", volunteer.Organization)})
		return
	}
	volunteerToUpdate = &volunteer
	c.JSON(http.StatusOK, volunteer)
}

func PatchVolunteer(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Volunteer, volunteers)
	organization := c.Param("organization")
	volunteerToUpdate, _ := models.FindVolunteer(volunteers, organization)
	if volunteerToUpdate.Organization == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s organization not found", volunteerToUpdate.Organization)})
		return
	}
	body := c.Request.Body
	var volunteer map[string]interface{}
	if err := utils.ReadFromBody(body, &volunteer); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	volunteerToUpdate.Patch(volunteer)
	c.JSON(http.StatusOK, &volunteerToUpdate)
}

func DeleteVolunteer(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Volunteer, volunteers)
	organization := c.Param("organization")
	volunteer, index := models.FindVolunteer(volunteers, organization)
	if volunteer.Organization == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s organization not found", volunteer.Organization)})
		return
	}
	volunteers = append(volunteers[:index], volunteers[index+1:]...)
	c.JSON(http.StatusOK, volunteer)
}
