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
		return c.HTML(http.StatusOK, getCatto())
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func getCatto() string {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	
	cats := []string{
		"Maine Coon",
		"Ragdoll",
		"Sphynx",
		"Russian Blue",
		"Scottish Fold",
		"British Shorthair",
	}

	return cats[rand.Intn(len(cats))]
}
