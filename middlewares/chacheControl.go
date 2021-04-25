package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hhsnopek/etag"
)

var Etags = make(map[string]string)
var versions = make(map[string]int)

func CacheControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Header("cache-control", "public")
		c.Header("Etag", generateEtag(path, c.Params))

		if match := c.Request.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(Etags[path], match) {
				c.AbortWithStatus(http.StatusNotModified)
				return
			}
		}
		if match := c.Request.Header.Get("If-Match"); match != "" {
			if !strings.Contains(Etags[path], match) {
				c.AbortWithStatus(http.StatusConflict)
				return
			}
		}
		c.Next()
	}
}

func generateEtag(path string, params gin.Params) string {
	var e = path
	for _, v := range params {
		e += fmt.Sprintf("-%s:%s", v.Key, v.Value)
	}
	// /basics/profiles-network:github-1
	ver := versions[path]
	ver++
	versions[path] = ver
	e += fmt.Sprintf("-%d", ver)
	Etags[path] = etag.Generate([]byte(e), false)
	return Etags[path]
}
