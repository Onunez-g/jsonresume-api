package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostProfile(c *gin.Context) {
	body := c.Request.Body
	var profile m.Profile
	if err := utils.ReadFromBody(body, &profile); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if profile.IfNetworkExists(m.MyResume.Basics.Profiles) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s profile already exists", profile.Network)})
		return
	}
	m.MyResume.Basics.Profiles = append(m.MyResume.Basics.Profiles, profile)
	c.JSON(http.StatusOK, profile)
}

func GetProfiles(c *gin.Context) {
	c.JSON(http.StatusOK, m.MyResume.Basics.Profiles)
}

func GetProfile(c *gin.Context) {
	network := c.Param("network")
	profile, _ := m.FindProfile(m.MyResume.Basics.Profiles, network)
	if profile.Network == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s profile not found", profile.Network)})
		return
	}
	c.JSON(http.StatusOK, &profile)
}

func PutProfile(c *gin.Context) {
	network := c.Param("network")
	profileToUpdate, index := m.FindProfile(m.MyResume.Basics.Profiles, network)
	if profileToUpdate.Network == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s profile not found", profileToUpdate.Network)})
		return
	}
	body := c.Request.Body
	var profile m.Profile
	if err := utils.ReadFromBody(body, &profile); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	profileToUpdate = &profile
	m.MyResume.Basics.Profiles[index] = *profileToUpdate
	c.JSON(http.StatusOK, profile)
}

func PatchProfile(c *gin.Context) {
	network := c.Param("network")
	profileToUpdate, index := m.FindProfile(m.MyResume.Basics.Profiles, network)
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
	m.MyResume.Basics.Profiles[index] = *profileToUpdate
	c.JSON(http.StatusOK, &profileToUpdate)
}

func DeleteProfile(c *gin.Context) {
	network := c.Param("network")
	profile, index := m.FindProfile(m.MyResume.Basics.Profiles, network)
	if profile.Network == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s profile not found", profile.Network)})
		return
	}
	m.MyResume.Basics.Profiles = append(m.MyResume.Basics.Profiles[:index], m.MyResume.Basics.Profiles[index+1:]...)
	c.JSON(http.StatusOK, profile)
}
