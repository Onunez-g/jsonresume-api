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
	r.GET("/basics", controllers.GetBasics)
	r.HEAD("/basics", controllers.GetBasics)
	auth.PUT("/basics", controllers.PutBasics)
	auth.PATCH("/basics", controllers.PatchBasics)
	auth.DELETE("/basics", controllers.DeleteBasics)

	//Basics/profiles
	r.GET("/basics/profiles", controllers.GetProfiles)
	r.HEAD("/basics/profiles", controllers.GetProfiles)
	r.GET("/basics/profiles/{network}", controllers.GetProfile)
	r.HEAD("/basics/profiles/{network}", controllers.GetProfile)
	auth.POST("/basics/profiles", controllers.PostProfile)
	auth.PUT("/basics/profiles/{network}", controllers.PutProfile)
	auth.PATCH("/basics/profiles/{network}", controllers.PatchProfile)
	auth.DELETE("/basics/profiles/{network}", controllers.DeleteProfile)

	//Work
	r.GET("/works", controllers.GetWorks)
	r.HEAD("/works", controllers.GetWorks)
	r.GET("/works/{company}", controllers.GetWork)
	r.HEAD("/works/{company}", controllers.GetWork)
	auth.POST("/works", controllers.PostWork)
	auth.PUT("/works/{company}", controllers.PutWork)
	auth.PATCH("/works/{company}", controllers.PatchWork)
	auth.DELETE("/works/{company}", controllers.DeleteWork)
	return r
}
