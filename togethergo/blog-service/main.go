package main

import (
	"blog-service/internal/routers"
	"net/http"
	"time"
)

func main()  {
	/*
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"pong"})
	})
	r.Run(":9999")
	 */
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":9999",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

/*
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

 */