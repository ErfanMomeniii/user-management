package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/erfanmomeniii/user-management/internal/config"
	"github.com/erfanmomeniii/user-management/internal/http/handler"
	"github.com/erfanmomeniii/user-management/internal/http/server"
	"github.com/erfanmomeniii/user-management/internal/repository"
)

var validUserRequest = handler.UserRequest{
	FirstName: "A",
	LastName:  "B",
	Nickname:  "C",
	Email:     "abc@efg.com",
	Password:  "1234",
	Country:   "UK",
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

type UserTestSuite struct {
	suite.Suite
	assert *require.Assertions
	engine *echo.Echo
}

func (suite *UserTestSuite) SetupSuite() {
	suite.assert = suite.Require()

	cfg, err := config.Init("./../../../config.defaults.yaml")
	suite.assert.NoError(err)

	server.Init(cfg)

	suite.engine = server.E
	repository.User = repository.NewUserRepositoryMock()
}

func (suite *UserTestSuite) Test_Save_User() {
	url := "/v1/user"
	cases := []struct {
		name    string
		request handler.UserRequest
		status  int
	}{
		{
			name:    "with valid request should be passed",
			request: validUserRequest,
			status:  http.StatusOK,
		},
		{
			name:    "with invalid request should be failed",
			request: handler.UserRequest{},
			status:  http.StatusBadRequest,
		},
	}
	for _, c := range cases {
		suite.Run(c.name, func() {
			data, err := json.Marshal(c.request)
			suite.NoError(err)

			request := httptest.NewRequest(echo.POST, url, bytes.NewReader(data))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			recorder := httptest.NewRecorder()
			ctx := suite.engine.NewContext(request, recorder)

			suite.assert.NoError(handler.SaveUser(ctx))
			suite.assert.Equal(c.status, recorder.Code)
		})
	}
}

func (suite *UserTestSuite) Test_Update_User() {
	url := "/v1/user/"
	cases := []struct {
		name    string
		userId  string
		request handler.UserRequest
		status  int
	}{
		{
			name:    "with user id should be passed",
			userId:  "1",
			request: validUserRequest,
			status:  http.StatusOK,
		},
		{
			name:    "without user id should be failed",
			userId:  "",
			request: validUserRequest,
			status:  http.StatusBadRequest,
		},
		{
			name:    "with invalid request should be failed",
			userId:  "1",
			request: handler.UserRequest{},
			status:  http.StatusBadRequest,
		},
	}
	for _, c := range cases {
		suite.Run(c.name, func() {
			data, err := json.Marshal(c.request)
			suite.NoError(err)

			request := httptest.NewRequest(echo.PUT, url+c.userId, bytes.NewReader(data))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			recorder := httptest.NewRecorder()

			ctx := suite.engine.NewContext(request, recorder)
			ctx.SetParamNames("userId")
			ctx.SetParamValues(c.userId)

			suite.assert.NoError(handler.UpdateUser(ctx))
			suite.assert.Equal(c.status, recorder.Code)
		})
	}
}

func (suite *UserTestSuite) Test_Delete_User() {
	url := "/v1/user/"
	cases := []struct {
		name   string
		userId string
		status int
	}{
		{
			name:   "with user id should be passed",
			userId: "1",
			status: http.StatusOK,
		},
		{
			name:   "without user id should be failed",
			userId: "",
			status: http.StatusBadRequest,
		},
	}
	for _, c := range cases {
		suite.Run(c.name, func() {

			request := httptest.NewRequest(echo.DELETE, url+c.userId, bytes.NewReader([]byte{}))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			recorder := httptest.NewRecorder()

			ctx := suite.engine.NewContext(request, recorder)
			ctx.SetParamNames("userId")
			ctx.SetParamValues(c.userId)

			suite.assert.NoError(handler.DeleteUser(ctx))
			suite.assert.Equal(c.status, recorder.Code)
		})
	}
}

func (suite *UserTestSuite) Test_Get_User() {
	url := "/v1/user/"
	cases := []struct {
		name   string
		userId string
		status int
	}{
		{
			name:   "with user id should be passed",
			userId: "1",
			status: http.StatusOK,
		},
		{
			name:   "without user id should be failed",
			userId: "",
			status: http.StatusBadRequest,
		},
	}
	for _, c := range cases {
		suite.Run(c.name, func() {

			request := httptest.NewRequest(echo.GET, url+c.userId, bytes.NewReader([]byte{}))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			recorder := httptest.NewRecorder()

			ctx := suite.engine.NewContext(request, recorder)
			ctx.SetParamNames("userId")
			ctx.SetParamValues(c.userId)

			suite.assert.NoError(handler.GetUser(ctx))
			suite.assert.Equal(c.status, recorder.Code)
		})
	}
}
