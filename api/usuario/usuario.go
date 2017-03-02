package usuario

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const (
	prefix = "/v1/usuario"
)

type response struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

// Usuario struct
type Usuario struct {
	GetAllURL string
	GetOneURL string
}

// NewUsuarioAPI instance
func NewUsuarioAPI() *Usuario {
	return &Usuario{
		GetAllURL: prefix + "/list",
		GetOneURL: prefix + "/get",
	}
}

// GetAll -> lista todos os usuários
func (u *Usuario) GetAll(c echo.Context) error {
	r := &response{Nome: "Rodolfo Azevedo", Idade: 33}
	return c.JSONPretty(http.StatusOK, r, "   ")
}

// GetOne -> retorna usuário pelo id
func (u *Usuario) GetOne(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
