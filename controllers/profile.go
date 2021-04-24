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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if profile.IfNetworkExists(profiles) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s profile already exists", profile.Network)})
		return
	}
	profiles = append(profiles, profile)
	c.JSON(http.StatusOK, profile)
}

func GetProfiles(c *gin.Context) {
	c.JSON(http.StatusOK, profiles)
}

func GetProfile(c *gin.Context) {
	network := c.Param("network")
	profile := models.FindProfile(profiles, network)
	if profile.Network == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s profile not found", profile.Network)})
		return
	}
	c.JSON(http.StatusOK, &profile)
}

func PutProfile(c *gin.Context) {
	network := c.Param("network")
	profileToUpdate := models.FindProfile(profiles, network)
	if profileToUpdate.Network == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s profile not found", profileToUpdate.Network)})
		return
	}
	body := c.Request.Body
	var profile models.Profile
	if err := utils.ReadFromBody(body, &profile); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if profile.IfNetworkExists(profiles) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s profile already exists", profile.Network)})
		return
	}
	profileToUpdate = &profile
	c.JSON(http.StatusOK, profile)
}

func PatchProfile(c *gin.Context) {
	network := c.Param("network")
	profileToUpdate := models.FindProfile(profiles, network)
	if profileToUpdate.Network == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s profile not found", profileToUpdate.Network)})
		return
	}
	body := c.Request.Body
	var profile map[string]interface{}
	if err := utils.ReadFromBody(body, &profile); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	profileToUpdate.Patch(profile)
	c.JSON(http.StatusOK, &profileToUpdate)
}
