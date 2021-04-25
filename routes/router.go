package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onunez-g/jsonresume-api/controllers"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()
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

	//Work
	r.GET("/resume/works", controllers.GetWorks)
	r.HEAD("/resume/works", controllers.GetWorks)
	r.GET("/resume/works/{company}", controllers.GetWork)
	r.HEAD("/resume/works/{company}", controllers.GetWork)
	auth.POST("/resume/works", controllers.PostWork)
	auth.PUT("/resume/works/{company}", controllers.PutWork)
	auth.PATCH("/resume/works/{company}", controllers.PatchWork)
	auth.DELETE("/resume/works/{company}", controllers.DeleteWork)
	return r
}
