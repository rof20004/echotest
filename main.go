package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rof20004/echotest/api/auth"
	"github.com/rof20004/echotest/api/usuario"
)

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Main group endpoint
	mainGroup := e.Group("/echotest")

	// Auth API
	authGroup := mainGroup.Group("/v1/login")
	authRoutes := auth.Routes{}
	authServices := auth.Services{}
	authGroup.POST(authRoutes.Login(), authServices.Login)

	// Usuario API
	usuarioGroup := mainGroup.Group("/v1/usuario")
	usuarioRoutes := usuario.Routes{}
	usuarioServices := usuario.Services{}
	usuarioGroup.Use(middleware.JWT(authServices.GetSignKey()))
	usuarioGroup.GET(usuarioRoutes.List(), usuarioServices.List)
	usuarioGroup.GET(usuarioRoutes.Get(), usuarioServices.Get)

	e.Logger.Fatal(e.Start(":1323"))
}
