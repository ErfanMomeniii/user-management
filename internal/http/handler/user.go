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
	id := ctx.Param("userId")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Id is not valid",
		})
	}

	if err := usecase.DeleteUser(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}

func UpdateUser(ctx echo.Context) error {
	var request UserRequest

	id := ctx.Param("userId")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Id is not valid",
		})
	}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Cannot parse the body.",
		})
	}

	if err := usecase.UpdateUser(id, model.User{FirstName: request.FirstName, LastName: request.LastName, Nickname: request.Nickname,
		Email: request.Email, Password: request.Password, Country: request.Country}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "User updated successfully",
	})
}

func GetUsers(ctx echo.Context) error {
	country := ctx.QueryParam("country")
	page := interface{}(ctx.QueryParam("page")).(int)
	if page == 0 {
		page = 1
	}
	switch country {
	case "":
		users, err := usecase.GetUsers(page)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"users": users,
		})

	default:
		users, err := usecase.FilterUsersByCountry(country, page)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"users": users,
		})
	}
}

func GetUser(ctx echo.Context) error {
	id := ctx.Param("userId")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Id is not valid",
		})
	}

	user, err := usecase.GetUserById(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
