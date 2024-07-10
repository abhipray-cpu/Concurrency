package main

import (
	"web_crawler/middleware"

	"github.com/labstack/echo/v4"
)

// routes registers the API routes with the provided Echo instance.
func (app *Config) routes(e *echo.Echo) {
	g := e.Group("/account")
	g.Use(middleware.JWTAuthMiddleware)
	e.GET("/ping", app.pingHandler)            // health check
	e.POST("/signup", app.signupHandler)       // user signup
	e.POST("/login", app.loginHandler)         // user login
	g.GET("/get", app.getUserHandler)          // get user by id
	g.PUT("/edit", app.updateUserHandler)      // update user by id
	g.DELETE("/delete", app.deleteUserHandler) // delete user by id

}
