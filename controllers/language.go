package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

func PostLanguage(c *gin.Context) {
	defer utils.UpdateResume(m.MyResume.Languages, m.MyResume.Languages)
	body := c.Request.Body
	var language m.Language
	if err := utils.ReadFromBody(body, &language); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if language.IfLanguageExists(m.MyResume.Languages) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s language already exists", language.Language)})
		return
	}
	m.MyResume.Languages = append(m.MyResume.Languages, language)
	c.JSON(http.StatusOK, language)
}

func GetLanguages(c *gin.Context) {
	c.JSON(http.StatusOK, m.MyResume.Languages)
}

func GetLanguage(c *gin.Context) {
	lang := c.Param("lang")
	language, _ := m.FindLanguage(m.MyResume.Languages, lang)
	if language.Language == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s language not found", language.Language)})
		return
	}
	c.JSON(http.StatusOK, &language)
}

func PutLanguage(c *gin.Context) {
	lang := c.Param("lang")
	languageToUpdate, _ := m.FindLanguage(m.MyResume.Languages, lang)
	if languageToUpdate.Language == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s language not found", languageToUpdate.Language)})
		return
	}
	body := c.Request.Body
	var language m.Language
	if err := utils.ReadFromBody(body, &language); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	languageToUpdate = &language
	c.JSON(http.StatusOK, language)
}

func PatchLanguage(c *gin.Context) {
	lang := c.Param("lang")
	languageToUpdate, _ := m.FindLanguage(m.MyResume.Languages, lang)
	if languageToUpdate.Language == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s language not found", languageToUpdate.Language)})
		return
	}
	body := c.Request.Body
	var language map[string]interface{}
	if err := utils.ReadFromBody(body, &language); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	languageToUpdate.Patch(language)
	c.JSON(http.StatusOK, &languageToUpdate)
}

func DeleteLanguage(c *gin.Context) {
	lang := c.Param("lang")
	language, index := m.FindLanguage(m.MyResume.Languages, lang)
	if language.Language == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s language not found", language.Language)})
		return
	}
	m.MyResume.Languages = append(m.MyResume.Languages[:index], m.MyResume.Languages[index+1:]...)
	c.JSON(http.StatusOK, language)
}
