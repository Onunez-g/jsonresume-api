package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var skills = models.MyResume.Skills

func PostSkill(c *gin.Context) {
	body := c.Request.Body
	var skill models.Skill
	if err := utils.ReadFromBody(body, &skill); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if skill.IfNameExists(skills) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s name already exists", skill.Name)})
		return
	}
	skills = append(skills, skill)
	c.JSON(http.StatusOK, skill)
}

func GetSkills(c *gin.Context) {
	c.JSON(http.StatusOK, skills)
}

func GetSkill(c *gin.Context) {
	name := c.Param("name")
	skill, _ := models.FindSkill(skills, name)
	if skill.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", skill.Name)})
		return
	}
	c.JSON(http.StatusOK, &skill)
}

func PutSkill(c *gin.Context) {
	name := c.Param("name")
	awardToUpdate, _ := models.FindSkill(skills, name)
	if awardToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", awardToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var skill models.Skill
	if err := utils.ReadFromBody(body, &skill); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if skill.IfNameExists(skills) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s name already exists", skill.Name)})
		return
	}
	awardToUpdate = &skill
	c.JSON(http.StatusOK, skill)
}

func PatchSkill(c *gin.Context) {
	name := c.Param("name")
	skillToUpdate, _ := models.FindSkill(skills, name)
	if skillToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", skillToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var skill map[string]interface{}
	if err := utils.ReadFromBody(body, &skill); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	skillToUpdate.Patch(skill)
	c.JSON(http.StatusOK, &skillToUpdate)
}

func DeleteSkill(c *gin.Context) {
	name := c.Param("name")
	skill, index := models.FindSkill(skills, name)
	if skill.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", skill.Name)})
		return
	}
	skills = append(skills[:index], skills[index+1:]...)
	c.JSON(http.StatusOK, skill)
}
