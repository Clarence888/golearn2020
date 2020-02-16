package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("access is now")
		return c.String(http.StatusOK, "Hello, World!王者不可阻挡")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
