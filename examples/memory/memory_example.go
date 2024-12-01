package main

import (
	"github.com/chenyahui/gin-cache/persist/memory"
	"time"

	"github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	memoryStore := memory.NewMemoryStore(1 * time.Minute)

	app.GET("/hello",
		cache.CacheByRequestURI(memoryStore, 2*time.Second),
		func(c *gin.Context) {
			c.String(200, "hello world")
		},
	)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
