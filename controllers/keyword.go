package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostKeyword(c *gin.Context) {
	skill, _ := findKeywordIndex(c, false)
	if skill.Name == "" {
		return
	}
	var keyword string
	body := c.Request.Body
	if err := utils.ReadFromBody(body, &keyword); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	index := skill.IndexKeyword(keyword)
	if index != -1 {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s keyword already exists", keyword)})
		return
	}
	skill.Keywords = append(skill.Keywords, keyword)
	c.JSON(http.StatusOK, keyword)
}

func GetKeywords(c *gin.Context) {
	skill, _ := findKeywordIndex(c, false)
	if skill.Name == "" {
		return
	}
	c.JSON(http.StatusOK, skill.Keywords)
}
func GetKeyword(c *gin.Context) {
	skill, index := findKeywordIndex(c, true)
	if skill.Name == "" || index == -1 {
		return
	}
	c.JSON(http.StatusOK, skill.Keywords[index])
}
func PutKeyword(c *gin.Context) {
	skill, index := findKeywordIndex(c, true)
	if skill.Name == "" || index == -1 {
		return
	}
	var keyword string
	body := c.Request.Body
	if err := utils.ReadFromBody(body, &keyword); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	skill.Keywords[index] = keyword
	c.JSON(http.StatusOK, keyword)
}
func DeleteKeyword(c *gin.Context) {
	skill, index := findKeywordIndex(c, true)
	if skill.Name == "" || index == -1 {
		return
	}
	key := skill.Keywords[index]
	skill.Keywords = append(skill.Keywords[:index], skill.Keywords[index+1:]...)
	c.JSON(http.StatusOK, key)
}

func findKeywordIndex(c *gin.Context, needsKey bool) (*models.Skill, int) {
	name := c.Param("name")
	key := c.Param("key")
	list := SkillorInterest(c.Request.URL.Path)
	skill, _ := models.FindSkill(list, name)
	if skill.Name == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", skill.Name)})
		return skill, -1
	}
	if !needsKey {
		return skill, -1
	}
	index := skill.IndexKeyword(key)
	if index == -1 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s keyword not found", key)})
		return skill, -1
	}
	return skill, index
}
