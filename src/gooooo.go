package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	type User struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Birthdate string `json:"date_of_birth"`
	}

	e := echo.New()

	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")

		u := &User{
			Id:        id,
			Name:      "Karel Novák",
			Email:     "karel.novak@whalebone.io",
			Birthdate: time.Now().Format("2006-01-02T15:04:05-07:00"),
		}

		return c.JSON(http.StatusOK, u)
	})

	e.POST("/save", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":80"))
}
