package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func makeTime(tik int, data chan string) {

	time.Sleep(time.Duration(tik) * time.Second)

	data <- "Successed"
}

// User
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/sign", func(c echo.Context) error {

		if c.FormValue("id") == "duckbo" && c.FormValue("pw") == "1234" {
			type MyCustomClaims struct {
				Identifier string `json:"identifier"`
				jwt.StandardClaims
			}

			// Create the Claims
			claims := MyCustomClaims{
				"jjang",
				jwt.StandardClaims{
					ExpiresAt: 15000,
					Issuer:    "test",
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			ss, _ := token.SignedString([]byte("duckbo"))

			c.Response().Header().Set("x-access-token", ss)
			return c.NoContent(200)

		} else {
			return c.NoContent(401)
		}

	})

	// 첫 화면
	e.GET("/", func(c echo.Context) error {
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = "jon"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		ch := make(chan string)
		go makeTime(3, ch)

		select {
		case v := <-ch:
			fmt.Println(v)
			break
		}

		return c.String(http.StatusOK, "최종테스트 - 개발 - 푸시 - 빌드 - 배포 - 확인 - 최종 - 끝 - 마지막")
	})

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("duckbo")))
	r.GET("", restricted)

	e.GET("/me", func(c echo.Context) error {

		token := c.Request().Header.Get("x-access-token")

		if token != "" {
			return c.String(http.StatusOK, "굳")
		} else {
			return c.String(http.StatusUnauthorized, "에러")
		}

	})

	e.GET("/get", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.JSON(http.StatusOK, u)
	})

	e.GET(("/read"), func(c echo.Context) error {
		cookie, err := c.Cookie("username")
		if err != nil {
			return err
		}
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
		return c.String(http.StatusOK, "read a cookie"+cookie.Value)

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

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}
