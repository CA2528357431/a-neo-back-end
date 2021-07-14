package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)



func main()  {
	ech := echo.New()

	ech.Use(middleware.Recover())



	ech.POST("/signIn",signIn)

	ech.POST("/signUp",signUp)

	ech.POST("/addFavorite/:id",addFavorite)

	ech.POST("/deleteFavorite/:id",deleteFavorite)

	ech.GET("/showFavorite/:id",showFavorite)

	ech.POST("/codeEmail",codeEmail)





	ech.Start(":8888")

}


