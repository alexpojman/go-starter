package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SaveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	log.Info().Str("name", u.Name).Str("email", u.Email).Msg("User saved")
	return c.JSON(http.StatusCreated, u)
}