package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var (
	signKey = []byte("secret")
)

// Services struct
type Services struct{}

// GetSignKey -> return sign key
func (s *Services) GetSignKey() []byte {
	return signKey
}

// Login -> authentication endpoint
func (s *Services) Login(c echo.Context) error {
	r := new(Request)
	c.Bind(r)

	if r.Login == "admin" && r.Password == "admin" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

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
