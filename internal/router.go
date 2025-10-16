package internal

import (
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Create a limiter struct.
	limit := tollbooth.NewLimiter(5, nil)

	// r.GET("/", tollbooth_gin.LimitHandler(limit), func(c *gin.Context) {
	// 	c.String(200, "Hello, world!")
	// })
	r.Use(tollbooth_gin.LimitHandler(limit))
	return r

}
