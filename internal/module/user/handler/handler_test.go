package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"user-service/internal/module/user/handler"
	"user-service/internal/module/user/mocks"
	"user-service/internal/module/user/models/request"
	"user-service/internal/module/user/models/response"
	"user-service/internal/pkg/log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

var (
	usecases      *mocks.Usecases
	logTest       *log.Logger
	h             handler.UserHandler
	validatorTest *validator.Validate
	app           *fiber.App
)

func setup() {
	usecases = new(mocks.Usecases)
	logZap := log.SetupLogger()
	log.Init(logZap)
	logTest := log.GetLogger()
	validatorTest = validator.New()
	h = handler.UserHandler{
		Log:       logTest,
		Usecase:   usecases,
		Validator: validatorTest,
	}

	app = fiber.New()

}

func teardown() {
	usecases = nil
	logTest = nil
	h = handler.UserHandler{}
}

func TestRegister(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// Prepare test data
		payload := request.Register{
			// Set the required fields of the request struct
			Email:    "test@test.com",
			Password: "password",
		}

		// Prepare the request
		reqBody, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

		// Set the request body
		ctx.Request().SetBody(reqBody)
		ctx.Request().Header.SetMethod("POST")
		ctx.Request().Header.SetContentType("application/json")
		ctx.Request().SetRequestURI("/api/v1/register")

		// Set the expected response
		usecases.On("Register", ctx.Context(), &payload).Return(nil)

		// Call the function
		err := h.Register(ctx)

		// assert the result
		assert.Nil(t, err)

	})
}

func TestLogin(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// Prepare test data
		payload := request.Login{
			// Set the required fields of the request struct
			Email:    "test@test.com",
			Password: "password",
		}

		// Prepare the request
		reqBody, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		// Set the request body
		ctx.Request().SetBody(reqBody)
		ctx.Request().Header.SetMethod("POST")
		ctx.Request().Header.SetContentType("application/json")
		ctx.Request().SetRequestURI("/api/v1/login")

		// Set the expected response
		usecases.On("Login", ctx.Context(), &payload).Return(response.LoginResponse{
			Token:     "token",
			ExpiredAt: 0,
		}, nil)

		// Call the function
		err := h.Login(ctx)

		// assert the result
		assert.Nil(t, err)
	})
}

func TestGetUser(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// Prepare test data
		payload := request.GetUser{
			// Set the required fields of the request struct
			ID: 1,
		}

		// Prepare the request
		reqBody, _ := json.Marshal(payload)
		req := httptest.NewRequest("GET", "/api/v1/user", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		// Set the request body
		ctx.Request().SetBody(reqBody)
		ctx.Request().Header.SetMethod("GET")
		ctx.Request().Header.SetContentType("application/json")
		ctx.Request().SetRequestURI("/api/v1/user")
		ctx.Locals("userID", 1)

		// Set the expected response
		usecases.On("GetUser", ctx.Context(), &payload).Return(response.GetUserResponse{
			ID:    1,
			Email: "test@test.com",
		}, nil)

		// Call the function
		err := h.GetUser(ctx)

		// assert the result
		assert.Nil(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// Prepare test data
		payload := request.UpdateUser{
			// Set the required fields of the request struct
			ID:    1,
			Email: "testUpdate@test.com",
		}

		// Prepare the request
		reqBody, _ := json.Marshal(payload)
		req := httptest.NewRequest("PUT", "/api/v1/user", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		// Set the request body
		ctx.Request().SetBody(reqBody)
		ctx.Request().Header.SetMethod("PUT")
		ctx.Request().Header.SetContentType("application/json")
		ctx.Request().SetRequestURI("/api/v1/user")

		// Set the expected response
		usecases.On("UpdateUser", ctx.Context(), &payload).Return(nil)

		// Call the function
		err := h.UpdateUser(ctx)

		// assert the result
		assert.Nil(t, err)
	})
}

func TestCreateProfile(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// Prepare test data
		payload := request.CreateProfile{
			// Set the required fields of the request struct
			UserID:         1,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Senayan",
			City:           "Jakarta Selatan",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asean",
			Phone:          "08123456789",
			PersonalID:     "7890123456789012",
			TypePersonalID: "KTP",
		}

		// Prepare the request
		reqBody, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/api/v1/profile", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		// Set the request body
		ctx.Request().SetBody(reqBody)
		ctx.Request().Header.SetMethod("POST")
		ctx.Request().Header.SetContentType("application/json")
		ctx.Request().SetRequestURI("/api/v1/profile")

		// Set the expected response
		usecases.On("CreateProfile", ctx.Context(), &payload).Return(nil)

		// Call the function
		err := h.CreateProfile(ctx)

		// assert the result
		assert.Nil(t, err)
	})
}

func TestGetProfile(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// Prepare test data
		payload := request.GetProfile{
			// Set the required fields of the request struct
			UserID: 1,
		}

		// Prepare the request
		reqBody, _ := json.Marshal(payload)
		req := httptest.NewRequest("GET", "/api/private/profile", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		// Set the request body
		ctx.Request().SetBody(reqBody)
		ctx.Request().Header.SetMethod("GET")
		ctx.Request().Header.SetContentType("application/json")
		ctx.Request().SetRequestURI("/api/private/profile")
		ctx.Locals("userID", 1)

		// Set the expected response
		usecases.On("GetProfile", ctx.Context(), &payload).Return(response.GetProfileResponse{
			UserID: 1,
		}, nil)

		// Call the function
		err := h.GetProfile(ctx)

		// assert the result
		assert.Nil(t, err)
	})
}

func TestUpdateProfile(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// Prepare test data
		payload := request.UpdateProfile{
			ID:             1,
			UserID:         1,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Senayan",
			City:           "Jakarta Selatan",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asean",
			Phone:          "08123456789",
			PersonalID:     "7890123456789012",
			TypePersonalID: "KTP",
		}

		// Prepare the request
		reqBody, _ := json.Marshal(payload)
		req := httptest.NewRequest("PUT", "/api/v1/profile", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		// Set the request body
		ctx.Request().SetBody(reqBody)
		ctx.Request().Header.SetMethod("PUT")
		ctx.Request().Header.SetContentType("application/json")
		ctx.Request().SetRequestURI("/api/v1/profile")

		// Set the expected response
		usecases.On("UpdateProfile", ctx.Context(), &payload).Return(nil)

		// Call the function
		err := h.UpdateProfile(ctx)

		// assert the result
		assert.Nil(t, err)
	})
}

func TestValidateToken(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		payload := request.ValidateToken{
			Token: "token",
		}
		// Prepare the request
		req := httptest.NewRequest("GET", "/api/private/user/validate", nil)
		req.Header.Set("Content-Type", "application/json")

		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
		// Set the request body
		ctx.Request().Header.SetMethod("GET")
		ctx.Request().Header.SetContentType("application/json")
		ctx.Request().SetRequestURI("/api/private/user/validate")
		ctx.Request().URI().QueryArgs().Add("token", "token")

		// Set the expected response
		usecases.On("ValidateToken", ctx.Context(), &payload).Return(response.ValidateToken{
			IsValid:   true,
			UserID:    1,
			EmailUser: "test@test.com",
		}, nil)

		// Call the function
		err := h.ValidateToken(ctx)

		// assert the result
		assert.Nil(t, err)
	})
}
