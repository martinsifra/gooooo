package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

const birthdateFormat = "2006-01-02T15:04:05-07:00"

var (
	database *gorm.DB
)

func main() {
	connect()

	e := echo.New()

	e.GET("/:id", get)
	e.POST("/save", save)

	e.Logger.Fatal(e.Start(":80"))
}

func connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	var err error
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func get(c echo.Context) error {
	id := c.Param("id")

	var user User
	result := database.First(&user, "id = ?", id).Error

	if errors.Is(result, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusNotFound)
	}

	dto := &UserDto{
		Id:        user.Id.String(),
		Name:      *user.Name,
		Email:     user.Email,
		Birthdate: user.Birthdate.Format(birthdateFormat),
	}

	return c.JSON(http.StatusOK, dto)
}

func save(c echo.Context) error {
	dto := new(UserDto)

	errBind := c.Bind(dto)
	if errBind != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	var existing User
	result := database.First(&existing, "id = ?", dto.Id).Error

	id, errUuid := uuid.Parse(dto.Id)
	if errUuid != nil {
		return c.JSON(http.StatusBadRequest, "invalid UUID")
	}

	birthdate, _ := time.Parse(birthdateFormat, dto.Birthdate)

	user := &User{
		Id:        id,
		Name:      &dto.Name,
		Email:     dto.Email,
		Birthdate: &birthdate,
	}

	if errors.Is(result, gorm.ErrRecordNotFound) {
		// Create
		database.Create(&user)
	} else {
		// Update
		database.Save(&user)
	}

	return c.NoContent(http.StatusOK)
}

type UserDto struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Birthdate string `json:"date_of_birth"`

	// todo Validation should be done with https://echo.labstack.com/docs/request#validate-data
}

type User struct {
	Id        uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name      *string
	Email     string `gorm:"unique;not null;"`
	Birthdate *time.Time
}
