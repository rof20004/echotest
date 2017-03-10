package usuario

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Services instance
type Services struct{}

// List -> lista todos os usuários
func (s *Services) List(c echo.Context) error {
	r := &Response{Nome: "Rodolfo Azevedo", Idade: 33}
	return c.JSONPretty(http.StatusOK, r, "   ")
}

// Get -> retorna um usuário pelo id
func (s *Services) Get(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
