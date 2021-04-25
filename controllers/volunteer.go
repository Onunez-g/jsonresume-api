package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostVolunteer(c *gin.Context) {

	body := c.Request.Body
	var volunteer m.Work
	if err := utils.ReadFromBody(body, &volunteer); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if volunteer.IfOrganizationExists(m.MyResume.Volunteer) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s organization already exists", volunteer.Organization)})
		return
	}
	m.MyResume.Volunteer = append(m.MyResume.Volunteer, volunteer)
	c.JSON(http.StatusOK, volunteer)
}

func GetVolunteers(c *gin.Context) {
	c.JSON(http.StatusOK, m.MyResume.Volunteer)
}

func GetVolunteer(c *gin.Context) {
	organization := c.Param("organization")
	volunteer, _ := m.FindVolunteer(m.MyResume.Volunteer, organization)
	if volunteer.Organization == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s organization not found", volunteer.Organization)})
		return
	}
	c.JSON(http.StatusOK, &volunteer)
}

func PutVolunteer(c *gin.Context) {

	organization := c.Param("organization")
	volunteerToUpdate, _ := m.FindVolunteer(m.MyResume.Volunteer, organization)
	if volunteerToUpdate.Organization == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s organization not found", volunteerToUpdate.Organization)})
		return
	}
	body := c.Request.Body
	var volunteer m.Work
	if err := utils.ReadFromBody(body, &volunteer); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	volunteerToUpdate = &volunteer
	c.JSON(http.StatusOK, volunteer)
}

func PatchVolunteer(c *gin.Context) {

	organization := c.Param("organization")
	volunteerToUpdate, _ := m.FindVolunteer(m.MyResume.Volunteer, organization)
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

	organization := c.Param("organization")
	volunteer, index := m.FindVolunteer(m.MyResume.Volunteer, organization)
	if volunteer.Organization == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s organization not found", volunteer.Organization)})
		return
	}
	m.MyResume.Volunteer = append(m.MyResume.Volunteer[:index], m.MyResume.Volunteer[index+1:]...)
	c.JSON(http.StatusOK, volunteer)
}
