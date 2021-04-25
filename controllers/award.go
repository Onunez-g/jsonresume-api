package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostAward(c *gin.Context) {
	body := c.Request.Body
	var award m.Award
	if err := utils.ReadFromBody(body, &award); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if award.IfTitleExists(m.MyResume.Awards) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s title already exists", award.Title)})
		return
	}
	m.MyResume.Awards = append(m.MyResume.Awards, award)
	c.JSON(http.StatusOK, award)
}

func GetAwards(c *gin.Context) {
	c.JSON(http.StatusOK, m.MyResume.Awards)
}

func GetAward(c *gin.Context) {
	title := c.Param("title")
	award, _ := m.FindAward(m.MyResume.Awards, title)
	if award.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s title not found", award.Title)})
		return
	}
	c.JSON(http.StatusOK, &award)
}

func PutAward(c *gin.Context) {
	title := c.Param("title")
	awardToUpdate, index := m.FindAward(m.MyResume.Awards, title)
	if awardToUpdate.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s title not found", awardToUpdate.Title)})
		return
	}
	body := c.Request.Body
	var award m.Award
	if err := utils.ReadFromBody(body, &award); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	awardToUpdate = &award
	m.MyResume.Awards[index] = *awardToUpdate
	c.JSON(http.StatusOK, award)
}

func PatchAward(c *gin.Context) {
	title := c.Param("title")
	awardToUpdate, index := m.FindAward(m.MyResume.Awards, title)
	if awardToUpdate.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s title not found", awardToUpdate.Title)})
		return
	}
	body := c.Request.Body
	var award map[string]interface{}
	if err := utils.ReadFromBody(body, &award); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	awardToUpdate.Patch(award)
	m.MyResume.Awards[index] = *awardToUpdate
	c.JSON(http.StatusOK, &awardToUpdate)
}

func DeleteAward(c *gin.Context) {
	title := c.Param("title")
	award, index := m.FindAward(m.MyResume.Awards, title)
	if award.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s title not found", award.Title)})
		return
	}
	m.MyResume.Awards = append(m.MyResume.Awards[:index], m.MyResume.Awards[index+1:]...)
	c.JSON(http.StatusOK, award)
}
