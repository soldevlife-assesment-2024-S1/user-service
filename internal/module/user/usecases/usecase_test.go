package usecases_test

import (
	"context"
	"errors"
	"testing"
	"user-service/internal/module/user/mocks"
	"user-service/internal/module/user/models/entity"
	"user-service/internal/module/user/models/request"
	"user-service/internal/module/user/models/response"
	"user-service/internal/module/user/usecases"
	"user-service/internal/pkg/log"

	"github.com/stretchr/testify/assert"
)

var (
	repositories *mocks.Repositories
	logTest      *log.Logger
	uc           usecases.Usecases
	ctx          context.Context
)

func setup() {
	repositories = new(mocks.Repositories)
	logZap := log.SetupLogger()
	log.Init(logZap)
	logTest := log.GetLogger()
	uc = usecases.New(repositories, logTest)
	ctx = context.Background()
}

func teardown() {
	repositories = nil
	logTest = nil
	uc = nil
}

func TestGetProfile(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// mock data
		payload := &request.GetProfile{
			UserID: 1,
		}

		entityMock := entity.Profile{
			ID:             1,
			UserID:         1,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Kota",
			City:           "Jakarta",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asia",
			Phone:          "08123456789",
			PersonalID:     "1234567890",
			TypePersonalID: "KTP",
		}

		expectedResponse := response.GetProfileResponse{
			ID:             1,
			UserID:         1,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Kota",
			City:           "Jakarta",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asia",
			Phone:          "08123456789",
			PersonalID:     "1234567890",
			TypePersonalID: "KTP",
		}
		// mock repository
		repositories.On("FindProfileByUserID", ctx, payload.UserID).Return(entityMock, nil)
		// run usecase
		response, err := uc.GetProfile(ctx, payload)

		// assert result

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)

	})

	t.Run("error", func(t *testing.T) {
		// mock data
		payload := &request.GetProfile{
			UserID: 2,
		}

		expectedResponseError := response.GetProfileResponse{}

		// mock repository
		repositories.On("FindProfileByUserID", ctx, payload.UserID).Return(entity.Profile{}, errors.New("error"))
		// run usecase
		responseError, err := uc.GetProfile(ctx, payload)

		// assert result
		assert.Error(t, err)
		assert.Equal(t, expectedResponseError, responseError)
	})
}

func TestGetUser(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// mock data
		payload := &request.GetUser{
			ID: 1,
		}

		entityMock := entity.User{
			ID:       1,
			Email:    "test@test.com",
			Password: "password",
		}

		expectedResponse := response.GetUserResponse{
			ID:    1,
			Email: "test@test.com",
		}

		// mock repository
		repositories.On("FindUserByID", ctx, payload.ID).Return(entityMock, nil)
		// run usecase
		response, err := uc.GetUser(ctx, payload)

		// assert result
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("error", func(t *testing.T) {
		// mock data
		payload := &request.GetUser{
			ID: 2,
		}

		expectedResponseError := response.GetUserResponse{}

		// mock repository
		repositories.On("FindUserByID", ctx, payload.ID).Return(entity.User{}, errors.New("error"))
		// run usecase
		responseError, err := uc.GetUser(ctx, payload)

		// assert result
		assert.Error(t, err)
		assert.Equal(t, expectedResponseError, responseError)
	})
}

func TestRegister(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// mock data
		// payload := &request.Register{
		// 	Email:    "test@test.com",
		// 	Password: "password",
		// }

		// hashedPassword := helpers.CheckPasswordHash(payload.Password, "password")

		// // mock repository
		// repositories.On("FindUserByEmail", ctx, payload.Email).Return(entity.User{}, nil)
		// repositories.On("UpsertUser", ctx, &entity.User{
		// 	Email:    payload.Email,
		// 	Password: hashedPassword,
		// }).Return(nil)

		// // run usecase
		// err := uc.Register(ctx, payload)

		// // assert result
		// assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// mock data
		payload := &request.Register{
			Email:    "testError@test.com",
			Password: "password",
		}

		// mock repository
		repositories.On("FindUserByEmail", ctx, payload.Email).Return(entity.User{}, errors.New("error"))

		// run usecase
		err := uc.Register(ctx, payload)

		// assert result
		assert.Error(t, err)
	})
}

func TestUpdateProfile(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// mock data
		payload := &request.UpdateProfile{
			UserID:         1,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Kota",
			City:           "Jakarta",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asia",
			Phone:          "08123456789",
			PersonalID:     "1234567890",
			TypePersonalID: "KTP",
		}

		mockEntity := entity.Profile{
			ID:             1,
			UserID:         1,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Kota",
			City:           "Jakarta",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asia",
			Phone:          "08123456789",
			PersonalID:     "1234567890",
			TypePersonalID: "KTP",
		}

		// mock repository
		repositories.On("FindProfileByUserID", ctx, payload.UserID).Return(mockEntity, nil)

		repositories.On("UpsertProfile", ctx, &entity.Profile{
			ID:             0,
			UserID:         1,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Kota",
			City:           "Jakarta",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asia",
			Phone:          "08123456789",
			PersonalID:     "1234567890",
			TypePersonalID: "KTP",
		}).Return(nil)

		// run usecase
		err := uc.UpdateProfile(ctx, payload)

		// assert result
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// mock data
		payload := &request.UpdateProfile{
			UserID:         2,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Kota",
			City:           "Jakarta",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asia",
			Phone:          "08123456789",
			PersonalID:     "1234567890",
			TypePersonalID: "KTP",
		}

		// mock repository
		repositories.On("FindProfileByUserID", ctx, payload.UserID).Return(entity.Profile{}, errors.New("error"))

		// run usecase
		err := uc.UpdateProfile(ctx, payload)

		// assert result
		assert.Error(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// mock data
		payload := &request.UpdateUser{
			ID:    1,
			Email: "test@test.com",
		}

		mockEntity := entity.User{
			ID:    1,
			Email: "test@test.com",
		}

		// mock repository
		repositories.On("FindUserByID", ctx, payload.ID).Return(mockEntity, nil)
		repositories.On("UpsertUser", ctx, &mockEntity).Return(nil)

		// run usecase
		err := uc.UpdateUser(ctx, payload)

		// assert result
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// mock data
		payload := &request.UpdateUser{
			ID:    2,
			Email: "testError@test.com",
		}

		// mock repository
		repositories.On("FindUserByID", ctx, payload.ID).Return(entity.User{}, errors.New("error"))

		// run usecase
		err := uc.UpdateUser(ctx, payload)

		// assert result
		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// mock data
	})

	t.Run("error", func(t *testing.T) {
		// mock data
		payload := &request.Login{
			Email:    "testError@test.com",
			Password: "password",
		}

		// mock repository
		repositories.On("FindUserByEmail", ctx, payload.Email).Return(entity.User{}, errors.New("error"))

		// run usecase
		resp, err := uc.Login(ctx, payload)

		// assert result
		assert.Error(t, err)
		assert.Equal(t, resp, response.LoginResponse{})
	})
}

func TestCreateProfile(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// mock data
		payload := &request.CreateProfile{
			UserID:         1,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Kota",
			City:           "Jakarta",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asia",
			Phone:          "08123456789",
			PersonalID:     "1234567890",
			TypePersonalID: "KTP",
		}

		mockEntity := entity.Profile{
			UserID:         1,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Kota",
			City:           "Jakarta",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asia",
			Phone:          "08123456789",
			PersonalID:     "1234567890",
			TypePersonalID: "KTP",
		}

		// mock repository
		repositories.On("FindUserByID", ctx, payload.UserID).Return(entity.User{
			ID: 1,
		}, nil)
		repositories.On("UpsertProfile", ctx, &mockEntity).Return(nil)

		// run usecase
		err := uc.CreateProfile(ctx, payload)

		// assert result
		assert.NoError(t, err)

	})

	t.Run("error", func(t *testing.T) {
		// mock data
		payload := &request.CreateProfile{
			UserID:         2,
			FirstName:      "John",
			LastName:       "Doe",
			Address:        "Jl. Jendral Sudirman",
			District:       "Kota",
			City:           "Jakarta",
			State:          "DKI Jakarta",
			Country:        "Indonesia",
			Region:         "Asia",
			Phone:          "08123456789",
			PersonalID:     "1234567890",
			TypePersonalID: "KTP",
		}

		// mock repository
		repositories.On("FindUserByID", ctx, payload.UserID).Return(entity.User{}, errors.New("error"))

		// run usecase
		err := uc.CreateProfile(ctx, payload)

		// assert result
		assert.Error(t, err)
	})
}

func TestValidateToken(t *testing.T) {
	setup()
	defer teardown()

	t.Run("success", func(t *testing.T) {
		// mock data

		// mock repository

		// run usecase

		// assert result

	})

	t.Run("error", func(t *testing.T) {
		// mock data

		// mock repository

		// run usecase

		// assert result
	})
}
