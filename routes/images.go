package routes

import "github.com/gin-gonic/gin"

func ImgRout(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})
}
