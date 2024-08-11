package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

func main() {
	type UserDto struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Birthdate string `json:"date_of_birth"`
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")

		u := &UserDto{
			Id:        id,
			Name:      "Karel Novák",
			Email:     "karel.novak@whalebone.io",
			Birthdate: time.Now().Format("2006-01-02T15:04:05-07:00"),
		}

		return c.JSON(http.StatusOK, u)
	})

	e.POST("/save", func(c echo.Context) error {
		u := new(UserDto)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":80"))
}
