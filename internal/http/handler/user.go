package handler

import (
	"database/sql"
	"github.com/erfanmomeniii/user-management/internal/model"
	"github.com/erfanmomeniii/user-management/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Nickname  string `json:"nickname" validate:""`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
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

	if _, err := repository.User.Save(model.User{FirstName: request.FirstName, LastName: request.LastName,
		Nickname: request.Nickname, Email: request.Email, Password: request.Password, Country: request.Country}); err != nil {
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

	if _, err := repository.User.Delete(id); err != nil {
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

	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if _, err := repository.User.Update(id, model.User{FirstName: request.FirstName, LastName: request.LastName,
		Nickname: request.Nickname, Email: request.Email, Password: request.Password, Country: request.Country}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "User updated successfully",
	})
}

func GetUsers(ctx echo.Context) error {
	country := ctx.QueryParam("country")

	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	if page == 0 {
		page = 1
	}

	switch country {
	case "":
		users, err := repository.User.GetAll(page)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"users": users,
		})

	default:
		users, err := repository.User.FilterByCountry(country, page)

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

	user, err := repository.User.Get(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
