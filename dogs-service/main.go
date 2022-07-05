package main

import (
	"net/http"
	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, getDoggo())
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func getDoggo() string {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	
	dogs := []string{
		"Border Collie",
		"Labrador",
		"Husky",
		"Corgi",
		"Shiba Inu",
		"Whippet",
	}

	return dogs[rand.Intn(len(dogs))]
}
