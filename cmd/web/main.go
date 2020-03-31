package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"tommynurwantoro/mygoweb/internal/domain"
	"tommynurwantoro/mygoweb/internal/http/handler"
	"tommynurwantoro/mygoweb/internal/http/router"

	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &domain.CustomContext{c, readConfig()}
			return next(cc)
		}
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Set Renderer
	e.Renderer = echoview.Default()

	// Init handler
	h := handler.NewHandler(e)

	// Init router
	router.NewRouter(e, h)

	e.Static("/assets", "assets")

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

// PRIVATE
func readConfig() *domain.StaticConfig {
	conf := &domain.StaticConfig{}
	yamlFile, err := ioutil.ReadFile("static_config.yaml")
	if err != nil {
		log.Printf("Error read yaml file: #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return conf
}
