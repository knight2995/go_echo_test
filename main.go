package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	// 첫 화면
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World! 테스트중")
	})

	e.GET("/text/:text", getText)
	e.POST(("/post"), postHandler)

	e.Logger.Fatal(e.Start(":80")) // localhost:1323
}

func getText(c echo.Context) error {

	text := c.Param("text")
	return c.String(http.StatusOK, text)
}

func postHandler(c echo.Context) error {

	id := c.FormValue("id")
	return c.String(http.StatusOK, id)
}
