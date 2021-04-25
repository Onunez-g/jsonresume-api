package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/models"
	"github.com/onunez-g/jsonresume-api/utils"
)

var languages = models.MyResume.Languages

func PostLanguage(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Languages, languages)
	body := c.Request.Body
	var language models.Language
	if err := utils.ReadFromBody(body, &language); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if language.IfLanguageExists(languages) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s language already exists", language.Language)})
		return
	}
	languages = append(languages, language)
	c.JSON(http.StatusOK, language)
}

func GetLanguages(c *gin.Context) {
	c.JSON(http.StatusOK, languages)
}

func GetLanguage(c *gin.Context) {
	lang := c.Param("lang")
	language, _ := models.FindLanguage(languages, lang)
	if language.Language == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s language not found", language.Language)})
		return
	}
	c.JSON(http.StatusOK, &language)
}

func PutLanguage(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Languages, languages)
	lang := c.Param("lang")
	languageToUpdate, _ := models.FindLanguage(languages, lang)
	if languageToUpdate.Language == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s language not found", languageToUpdate.Language)})
		return
	}
	body := c.Request.Body
	var language models.Language
	if err := utils.ReadFromBody(body, &language); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if language.IfLanguageExists(languages) {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s language already exists", language.Language)})
		return
	}
	languageToUpdate = &language
	c.JSON(http.StatusOK, language)
}

func PatchLanguage(c *gin.Context) {
	defer utils.UpdateResume(models.MyResume.Languages, languages)
	lang := c.Param("lang")
	languageToUpdate, _ := models.FindLanguage(languages, lang)
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
	defer utils.UpdateResume(models.MyResume.Languages, languages)
	lang := c.Param("lang")
	language, index := models.FindLanguage(languages, lang)
	if language.Language == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%s language not found", language.Language)})
		return
	}
	languages = append(languages[:index], languages[index+1:]...)
	c.JSON(http.StatusOK, language)
}
