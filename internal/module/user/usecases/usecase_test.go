package usecases_test

import (
	"context"
	"testing"
	"user-service/internal/module/user/models/response"
	"user-service/internal/module/user/usecases"
	"user-service/internal/pkg/helpers/errors"

	"go.uber.org/mock/gomock"
)

func TestGetProfile(t *testing.T) {
	repoMock := gomock.NewController(t)
	defer repoMock.Finish()

	uc := usecases.NewMockUsecases(repoMock)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {

		uc.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Return(response.GetProfileResponse{}, nil)
		_, err := uc.GetProfile(ctx, nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		uc.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Return(response.GetProfileResponse{}, errors.InternalServerError("error"))
		_, err := uc.GetProfile(ctx, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestGetUser(t *testing.T) {
	repoMock := gomock.NewController(t)
	defer repoMock.Finish()

	uc := usecases.NewMockUsecases(repoMock)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		uc.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(response.GetUserResponse{}, nil)
		_, err := uc.GetUser(ctx, nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		uc.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(response.GetUserResponse{}, errors.InternalServerError("error"))
		_, err := uc.GetUser(ctx, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestRegister(t *testing.T) {
	repoMock := gomock.NewController(t)
	defer repoMock.Finish()

	uc := usecases.NewMockUsecases(repoMock)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		uc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)
		err := uc.Register(ctx, nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		uc.EXPECT().Register(gomock.Any(), gomock.Any()).Return(errors.InternalServerError("error"))
		err := uc.Register(ctx, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestLogin(t *testing.T) {
	repoMock := gomock.NewController(t)
	defer repoMock.Finish()

	uc := usecases.NewMockUsecases(repoMock)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		uc.EXPECT().Login(gomock.Any(), gomock.Any()).Return(response.LoginResponse{}, nil)
		_, err := uc.Login(ctx, nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		uc.EXPECT().Login(gomock.Any(), gomock.Any()).Return(response.LoginResponse{}, errors.InternalServerError("error"))
		_, err := uc.Login(ctx, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestUpdateProfile(t *testing.T) {
	repoMock := gomock.NewController(t)
	defer repoMock.Finish()

	uc := usecases.NewMockUsecases(repoMock)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		uc.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(nil)
		err := uc.UpdateProfile(ctx, nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		uc.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(errors.InternalServerError("error"))
		err := uc.UpdateProfile(ctx, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestValidateToken(t *testing.T) {
	repoMock := gomock.NewController(t)
	defer repoMock.Finish()

	uc := usecases.NewMockUsecases(repoMock)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {

		uc.EXPECT().ValidateToken(gomock.Any(), gomock.Any()).Return(response.ValidateToken{}, nil)
		_, err := uc.ValidateToken(ctx, nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		uc.EXPECT().ValidateToken(gomock.Any(), gomock.Any()).Return(response.ValidateToken{}, errors.InternalServerError("error"))
		_, err := uc.ValidateToken(ctx, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestUpdateUser(t *testing.T) {
	repoMock := gomock.NewController(t)
	defer repoMock.Finish()

	uc := usecases.NewMockUsecases(repoMock)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		uc.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)
		err := uc.UpdateUser(ctx, nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		uc.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.InternalServerError("error"))
		err := uc.UpdateUser(ctx, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestCreateProfile(t *testing.T) {
	repoMock := gomock.NewController(t)
	defer repoMock.Finish()

	uc := usecases.NewMockUsecases(repoMock)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		uc.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(nil)
		err := uc.CreateProfile(ctx, nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		uc.EXPECT().CreateProfile(gomock.Any(), gomock.Any()).Return(errors.InternalServerError("error"))
		err := uc.CreateProfile(ctx, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
