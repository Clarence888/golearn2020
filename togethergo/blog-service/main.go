package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"pong"})
	})
	r.Run(":9999")

}

/*
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

 */