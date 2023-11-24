package main

import (
	"echo/config"
	"echo/controller"
	auth "echo/middleware"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Apps Car Documentation API
// @version 1.0
// @description This is a sample service for managing cars
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the value of the PORT variable from .env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("PORT is not set in .env file")
	}

	// initialize database connection
	config.Connect()

	e := echo.New()

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + config.CSRFTokenHeader,
		ContextKey:  config.CSRFKey,
	}))

	e.GET("/index", controller.Index)
	e.POST("/sayhello", controller.SayHello)
	e.POST("/helloworld", controller.HelloWorld)

	//group routes for employee
	emm := e.Group("/employee")
	emm.Use(auth.Authentication())
	emm.PUT("/", controller.UpdateEmployee)
	emm.DELETE("/:id", controller.DeleteEmployee)

	//group routes for item
	itm := e.Group("/item")
	itm.Use(auth.Authentication())
	itm.POST("/", controller.CreateItem)

	//routes for login & register employee
	e.POST("/login", controller.UserLogin)
	e.POST("/register", controller.CreateEmployee)

	// route for swagger

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
