package requests

import (
	"github.com/gin-gonic/gin"
)

// GetHello            godoc
// @Summary      Get hello
// @Description  first request
// @Tags         hello
// @Produce      json
// @Success      200 "good": "hello"
// @Router       /hello [get]

func hello_req(c *gin.Context) {
	c.JSON(200, gin.H{"good": "hello"})
	return
}
