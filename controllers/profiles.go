package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var profiles = models.MyResume.Basics.Profiles

func PostProfile(c *gin.Context) {
	body := c.Request.Body
	var profile models.Profile
	if err := utils.ReadFromBody(body, &profile); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	if profile.CheckPrimaryKeyExists(profiles) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Profile already exists"})
		return
	}
	profiles = append(profiles, profile)
	c.JSON(http.StatusCreated, profile)
}

func GetProfiles(c *gin.Context) {
	c.JSON(http.StatusOK, profiles)
}

func GetProfile(c *gin.Context) {
	network := c.Param("network")
	profile := models.FindProfile(profiles, network)
	if profile.Network == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s not found", network)})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func PutProfile(c *gin.Context) {
	network := c.Param("network")
	index := models.IndexProfile(profiles, network)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s not found", network)})
		return
	}
	body := c.Request.Body
	var profile models.Profile
	if err := utils.ReadFromBody(body, &profile); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	if profile.CheckPrimaryKeyExists(profiles) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Profile already exists"})
		return
	}
	profiles[index] = profile
	c.JSON(http.StatusOK, profile)
}
func DeleteProfile(c *gin.Context) {
	network := c.Param("network")
	index := models.IndexProfile(profiles, network)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s not found", network)})
		return
	}
	profiles = append(profiles[:index], profiles[index+1:]...)
	c.Status(http.StatusOK)
}

func PatchProfile(c *gin.Context) {
	network := c.Param("network")
	index := models.IndexProfile(profiles, network)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s not found", network)})
		return
	}
	body := c.Request.Body
	var profile map[string]interface{}
	if err := utils.ReadFromBody(body, &profile); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	profiles[index].Update(profile)
	c.JSON(http.StatusOK, profiles[index])
}
