package main

import (
	"clondy/database"
	"clondy/handler"

	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	db, err := database.Init()
	if err != nil {
		log.Fatal("Error DataBase")
	}

	h := handler.Handler{Collector: db}

	// Users
	e.POST("signup", h.Signup)
	e.POST("login", h.Login)
	e.GET("user/:id", h.GetUser)
	e.GET("users", h.GetUsers)

	// Projects
	e.GET("projects/:id", h.GetProjects)
	e.POST("project", h.CreateProject)
	e.PUT("project", h.UpdateProject)
	e.DELETE("project/:id", h.DeleteProject)

	// Layer
	e.POST("layer", h.CreateLayer)
	e.GET("layer/:id", h.GetLayers)
	e.DELETE("layer/:id", h.DeleteLayer)
	e.PUT("layer", h.UpdateLayer)

	// Tset
	e.GET("/", h.Index)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
