package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": r,
				})
			}
		}()
		c.Next()
	}
}
