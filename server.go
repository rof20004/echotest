package main

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rof20004/echotest/api/usuario"
)

var (
	signKey = []byte("secret")
)

// User struct
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

// Auth jwt
type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login API
	e.POST("/echotest/v1/login", login)

	// Main group echotest
	g := e.Group("/echotest")
	g.Use(middleware.JWT(signKey))

	// Usuario API
	u := usuario.NewUsuarioAPI()
	g.GET(u.GetAllURL, u.GetAll)
	g.GET(u.GetOneURL, u.GetOne)

	e.Logger.Fatal(e.Start(":1323"))
}

func login(c echo.Context) error {
	a := new(Auth)
	c.Bind(a)

	if a.Login == "admin" && a.Password == "admin" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Rodolfo Azevedo"
		claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString(signKey)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
