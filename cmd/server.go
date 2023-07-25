package main

import (
	"net/http"
	objectController "vrachapi/cmd/objects/controllers"
	userHandler "vrachapi/cmd/users/controllers"
	userModel "vrachapi/cmd/users/models"
	"vrachapi/configs"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//CORS opens for public
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Restricted group
	r := e.Group(configs.Version() + "/user/:username")
	username := e.Group(configs.Version() + "/user")

	/*
		//CORS opens for private domain
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"domain.tld"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
	*/

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.POST(configs.Version()+"/object/create", objectController.CreateObject)
	e.GET(configs.Version()+"/object/list", objectController.GetObjects)
	e.PUT(configs.Version()+"/object/edit", objectController.UpdateObject)
	e.DELETE(configs.Version()+"/object/delete/:id", objectController.DeleteObject)
	e.GET(configs.Version()+"/object/getdata", objectController.GetData)

	e.POST(configs.Version()+"/user/login", userHandler.UserLogin)
	e.GET(configs.Version()+"/users/list", userHandler.GetUsers)
	e.GET(configs.Version()+"/user/:username", userHandler.GetUserByUsername)
	e.GET(configs.Version()+"/user", userHandler.GetUsernameByJwt)

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(userModel.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	r.Use(echojwt.WithConfig(config))
	r.GET("", userHandler.GetUserByUsername)

	username.Use(echojwt.WithConfig(config))
	username.GET("", userHandler.GetUsernameByJwt)

	configs.InitDB()
	e.Logger.Fatal(e.Start(":1323"))
}
