package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/controllers"
	"github.com/onunez-g/jsonresume-api/middlewares"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.ErrorManager())
	auth := r.Group("/", gin.BasicAuth(gin.Accounts{
		"Onunez-g": "Hola1234",
	}))

	//Basics
	r.GET("/resume/basics", controllers.GetBasics)
	r.HEAD("/resume/basics", controllers.GetBasics)
	auth.PUT("/resume/basics", controllers.PutBasics)
	auth.PATCH("/resume/basics", controllers.PatchBasics)
	auth.DELETE("/resume/basics", controllers.DeleteBasics)

	//Basics/profiles
	r.GET("/resume/basics/profiles", controllers.GetProfiles)
	r.HEAD("/resume/basics/profiles", controllers.GetProfiles)
	r.GET("/resume/basics/profiles/{network}", controllers.GetProfile)
	r.HEAD("/resume/basics/profiles/{network}", controllers.GetProfile)
	auth.POST("/resume/basics/profiles", controllers.PostProfile)
	auth.PUT("/resume/basics/profiles/{network}", controllers.PutProfile)
	auth.PATCH("/resume/basics/profiles/{network}", controllers.PatchProfile)
	auth.DELETE("/resume/basics/profiles/{network}", controllers.DeleteProfile)

	//Basics/Location
	r.GET("/resume/basics/location", controllers.GetBasics)
	r.HEAD("/resume/basics/location", controllers.GetBasics)
	auth.PUT("/resume/basics/location", controllers.PutBasics)
	auth.PATCH("/resume/basics/location", controllers.PatchBasics)
	auth.DELETE("/resume/basics/location", controllers.DeleteBasics)

	//Work
	r.GET("/resume/works", controllers.GetWorks)
	r.HEAD("/resume/works", controllers.GetWorks)
	r.GET("/resume/works/{company}", controllers.GetWork)
	r.HEAD("/resume/works/{company}", controllers.GetWork)
	auth.POST("/resume/works", controllers.PostWork)
	auth.PUT("/resume/works/{company}", controllers.PutWork)
	auth.PATCH("/resume/works/{company}", controllers.PatchWork)
	auth.DELETE("/resume/works/{company}", controllers.DeleteWork)

	//Volunteer
	r.GET("/resume/volunteers", controllers.GetVolunteers)
	r.HEAD("/resume/volunteers", controllers.GetVolunteers)
	r.GET("/resume/volunteers/{organization}", controllers.GetVolunteer)
	r.HEAD("/resume/volunteers/{organization}", controllers.GetVolunteer)
	auth.POST("/resume/volunteers", controllers.PostVolunteer)
	auth.PUT("/resume/volunteers/{organization}", controllers.PutVolunteer)
	auth.PATCH("/resume/volunteers/{organization}", controllers.PatchVolunteer)
	auth.DELETE("/resume/volunteers/{organization}", controllers.DeleteVolunteer)

	//Education
	r.GET("/resume/educations", controllers.GetEducations)
	r.HEAD("/resume/educations", controllers.GetEducations)
	r.GET("/resume/educations/{institution}", controllers.GetEducation)
	r.HEAD("/resume/educations/{institution}", controllers.GetEducation)
	auth.POST("/resume/educations", controllers.PostEducation)
	auth.PUT("/resume/educations/{institution}", controllers.PutEducation)
	auth.PATCH("/resume/educations/{institution}", controllers.PatchEducation)
	auth.DELETE("/resume/educations/{institution}", controllers.DeleteEducation)

	//Awards
	r.GET("/resume/awards", controllers.GetAwards)
	r.HEAD("/resume/awards", controllers.GetAwards)
	r.GET("/resume/awards/{title}", controllers.GetAward)
	r.HEAD("/resume/awards/{title}", controllers.GetAward)
	auth.POST("/resume/awards", controllers.PostAward)
	auth.PUT("/resume/awards/{title}", controllers.PutAward)
	auth.PATCH("/resume/awards/{title}", controllers.PutAward)
	auth.DELETE("/resume/awards/{title}", controllers.DeleteAward)

	//Publications
	r.GET("/resume/publications", controllers.GetPublications)
	r.HEAD("/resume/publications", controllers.GetPublications)
	r.GET("/resume/publications/{name}", controllers.GetPublication)
	r.HEAD("/resume/publications/{name}", controllers.GetPublication)
	auth.POST("/resume/publications", controllers.PostPublication)
	auth.PUT("/resume/publications/{name}", controllers.PutPublication)
	auth.PATCH("/resume/publications/{name}", controllers.PatchPublication)
	auth.DELETE("/resume/publications/{name}", controllers.DeletePublication)

	//Skills
	r.GET("/resume/skills", controllers.GetSkills)
	r.HEAD("/resume/skills", controllers.GetSkills)
	r.GET("/resume/skills/{name}", controllers.GetSkill)
	r.HEAD("/resume/skills/{name}", controllers.GetSkill)
	auth.POST("/resume/skills", controllers.PostSkill)
	auth.PUT("/resume/skills/{name}", controllers.PutSkill)
	auth.PATCH("/resume/skills/{name}", controllers.PatchSkill)
	auth.DELETE("/resume/skills/{name}", controllers.DeleteSkill)

	//Interests
	r.GET("/resume/interests", controllers.GetSkills)
	r.HEAD("/resume/interests", controllers.GetSkills)
	r.GET("/resume/interests/{name}", controllers.GetSkill)
	r.HEAD("/resume/interests/{name}", controllers.GetSkill)
	auth.POST("/resume/interests", controllers.PostSkill)
	auth.PUT("/resume/interests/{name}", controllers.PutSkill)
	auth.PATCH("/resume/interests/{name}", controllers.PatchSkill)
	auth.DELETE("/resume/interests/{name}", controllers.DeleteSkill)

	//Skills/Keywords
	r.GET("/resume/skills/{name}/keywords", controllers.GetKeywords)
	r.HEAD("/resume/skills/{name}/keywords", controllers.GetKeywords)
	r.GET("/resume/skills/{name}/keywords/{key}", controllers.GetKeyword)
	r.HEAD("/resume/skills/{name}/keywords/{key}", controllers.GetKeyword)
	auth.POST("/resume/skills/{name}/keywords", controllers.PostKeyword)
	auth.PUT("/resume/skills/{name}/keywords/{key}", controllers.PutKeyword)
	auth.DELETE("/resume/skills/{name}/keywords/{key}", controllers.DeleteKeyword)

	//Interests/Keywords
	r.GET("/resume/interests/{name}/keywords", controllers.GetKeywords)
	r.HEAD("/resume/interests/{name}/keywords", controllers.GetKeywords)
	r.GET("/resume/interests/{name}/keywords/{key}", controllers.GetKeyword)
	r.HEAD("/resume/interests/{name}/keywords/{key}", controllers.GetKeyword)
	auth.POST("/resume/interests/{name}/keywords", controllers.PostKeyword)
	auth.PUT("/resume/interests/{name}/keywords/{key}", controllers.PutKeyword)
	auth.DELETE("/resume/interests/{name}/keywords/{key}", controllers.DeleteKeyword)

	//Languages
	r.GET("/resume/languages", controllers.GetLanguages)
	r.HEAD("/resume/languages", controllers.GetLanguages)
	r.GET("/resume/languages/{lang}", controllers.GetLanguage)
	r.HEAD("/resume/languages/{lang}", controllers.GetLanguage)
	auth.POST("/resume/languages", controllers.PostLanguage)
	auth.PUT("/resume/languages/{lang}", controllers.PutLanguage)
	auth.PATCH("/resume/languages/{lang}", controllers.PatchLanguage)
	auth.DELETE("/resume/languages/{lang}", controllers.DeleteLanguage)

	//References
	r.GET("/resume/references", controllers.GetReferences)
	r.HEAD("/resume/references", controllers.GetReferences)
	r.GET("/resume/references/{name}", controllers.GetReference)
	r.HEAD("/resume/references/{name}", controllers.GetReference)
	auth.POST("/resume/references", controllers.PostReference)
	auth.PUT("/resume/references/{name}", controllers.PutReference)
	auth.PATCH("/resume/references/{name}", controllers.PatchReference)
	auth.DELETE("/resume/references/{name}", controllers.DeleteReference)
	return r
}
