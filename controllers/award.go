package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var awards = models.MyResume.Awards

func PostAward(c *gin.Context) {
	body := c.Request.Body
	var award models.Award
	if err := utils.ReadFromBody(body, &award); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if award.IfTitleExists(awards) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s title already exists", award.Title)})
		return
	}
	awards = append(awards, award)
	c.JSON(http.StatusOK, award)
}

func GetAwards(c *gin.Context) {
	c.JSON(http.StatusOK, awards)
}

func GetAward(c *gin.Context) {
	title := c.Param("title")
	award, _ := models.FindAward(awards, title)
	if award.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s title not found", award.Title)})
		return
	}
	c.JSON(http.StatusOK, &award)
}

func PutAward(c *gin.Context) {
	title := c.Param("title")
	awardToUpdate, _ := models.FindAward(awards, title)
	if awardToUpdate.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s title not found", awardToUpdate.Title)})
		return
	}
	body := c.Request.Body
	var award models.Award
	if err := utils.ReadFromBody(body, &award); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if award.IfTitleExists(awards) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s title already exists", award.Title)})
		return
	}
	awardToUpdate = &award
	c.JSON(http.StatusOK, award)
}

func PatchAward(c *gin.Context) {
	title := c.Param("title")
	awardToUpdate, _ := models.FindAward(awards, title)
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
	c.JSON(http.StatusOK, &awardToUpdate)
}

func DeleteAward(c *gin.Context) {
	title := c.Param("title")
	award, index := models.FindAward(awards, title)
	if award.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s title not found", award.Title)})
		return
	}
	awards = append(awards[:index], awards[index+1:]...)
	c.JSON(http.StatusOK, award)
}
