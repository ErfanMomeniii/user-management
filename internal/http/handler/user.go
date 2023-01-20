package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-management/internal/model"
	"user-management/internal/usecase"
)

type UserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Nickname  string `json:"nickname" validate:""`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Country   string `json:"country" validate:"required"`
}

func GetAllUser(ctx echo.Context) error {
	var users []model.User

	page := interface{}(ctx.QueryParam("page")).(int)
	country := ctx.QueryParam("country")

	switch country {
	case "":
		users = usecase.GetUsers(page)
	default:
		users = usecase.FilterUsersByCountry(country, page)
	}

	return ctx.JSON(http.StatusOK, users)
}

func SaveUser(ctx echo.Context) error {
	var request UserRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Cannot parse the body.",
		})
	}

	if err := ctx.Validate(request); err != nil {

		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if err := usecase.SaveUser(model.User{FirstName: request.FirstName, LastName: request.LastName, Nickname: request.Nickname,
		Email: request.Email, Password: request.Password, Country: request.Country}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "User saved successfully",
	})
}

func DeleteUser(ctx echo.Context) error {
	return nil
}

func UpdateUser(ctx echo.Context) error {
	return nil
}
