package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var skills = models.MyResume.Skills
var interests = models.MyResume.Interests

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
	list := SkillorInterest(c.Request.URL.Path)
	defer updateList(c.Request.URL.Path, list)
	list = append(list, skill)
	c.JSON(http.StatusOK, skill)
}

func GetSkills(c *gin.Context) {
	list := SkillorInterest(c.Request.URL.Path)
	c.JSON(http.StatusOK, list)
}

func GetSkill(c *gin.Context) {
	name := c.Param("name")
	list := SkillorInterest(c.Request.URL.Path)
	skill, _ := models.FindSkill(list, name)
	if skill.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", skill.Name)})
		return
	}
	c.JSON(http.StatusOK, &skill)
}

func PutSkill(c *gin.Context) {
	name := c.Param("name")
	list := SkillorInterest(c.Request.URL.Path)
	skillToUpdate, _ := models.FindSkill(list, name)
	if skillToUpdate.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", skillToUpdate.Name)})
		return
	}
	body := c.Request.Body
	var skill models.Skill
	if err := utils.ReadFromBody(body, &skill); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	skillToUpdate = &skill
	c.JSON(http.StatusOK, skill)
}

func PatchSkill(c *gin.Context) {
	name := c.Param("name")
	list := SkillorInterest(c.Request.URL.Path)
	skillToUpdate, _ := models.FindSkill(list, name)
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
	list := SkillorInterest(c.Request.URL.Path)
	defer updateList(c.Request.URL.Path, list)
	skill, index := models.FindSkill(list, name)
	if skill.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s name not found", skill.Name)})
		return
	}
	list = append(list[:index], list[index+1:]...)
	c.JSON(http.StatusOK, skill)
}

func SkillorInterest(path string) []models.Skill {
	if models.IsSkill(path) {
		return skills
	}
	return interests
}

func updateList(path string, list []models.Skill) {
	if models.IsSkill(path) {
		utils.UpdateResume(models.MyResume.Skills, list)
	} else {
		utils.UpdateResume(models.MyResume.Interests, list)
	}
}
