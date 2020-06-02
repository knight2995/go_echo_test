package main

import (
	"net/http"
	"io/ioutil"
	"github.com/labstack/echo/v4"
)

func main() {

	d1 := []byte("hello\ngo\n")
	ioutil.WriteFile("dat1", d1, 0644)

	e := echo.New()
	// 첫 화면
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "최종테스트 - 개발 - 푸시 - 빌드 - 배포")
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
