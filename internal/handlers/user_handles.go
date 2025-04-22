package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"train-http/internal/repositories"
	"train-http/internal/services"
	"train-http/internal/validators"
	"train-http/pkg/database"
)

type Response struct {
	Message interface{} `json:"message"`
}

func DocsInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{
		Message: "Training API",
	})
}

func RegisterUser(c echo.Context) error {
	var user validators.PostUser
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
	}
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "Ошибка валидации: " + err.Error(),
		})
	}

	newRepo := repositories.NewUserRepo(database.DB())
	newUserService := services.NewUserService(newRepo)
	jwt, err := newUserService.RegisterUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
	}
	cookie := new(http.Cookie)
	cookie.Name = "JWT"
	cookie.Value = jwt
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, Response{
		Message: jwt,
	})
}

func GetUser(c echo.Context) error {
	var ID = c.QueryParam("id")
	newRepo := repositories.NewUserRepo(database.DB())
	newUserService := services.NewUserService(newRepo)
	user, err := newUserService.GetUser(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, Response{
		Message: user,
	})
}
