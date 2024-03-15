package usecases

import (
	"context"
	"fmt"

	"user-service/internal/module/user/models/entity"
	"user-service/internal/module/user/models/request"
	"user-service/internal/module/user/repositories"
	"user-service/internal/pkg/helpers"
	"user-service/internal/pkg/helpers/errors"
	"user-service/internal/pkg/log"
)

type usecases struct {
	repositories repositories.Repositories
	log          log.Logger
}

// GetProfile implements Usecases.
func (u *usecases) GetProfile(ctx context.Context, payload *request.GetProfileRequest) error {
	panic("unimplemented")
}

// GetUser implements Usecases.
func (u *usecases) GetUser(ctx context.Context, payload *request.GetUserRequest) error {
	panic("unimplemented")
}

// Login implements Usecases.
func (u *usecases) Login(ctx context.Context, payload *request.LoginRequest) error {
	panic("unimplemented")
}

// Register implements Usecases.
func (u *usecases) Register(ctx context.Context, payload *request.RegisterRequest) error {
	// check if user already exists
	userExisting, err := u.repositories.FindUserByEmail(ctx, payload.Email)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error finding user by email: %s", err.Error()))
	}

	if userExisting.ID != 0 {
		return errors.BadRequest("user already exists")
	}

	// hash password
	hashedPassword, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error hashing password: %s", err.Error()))
	}

	// create user
	user := entity.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	}

	if err := u.repositories.UpsertUser(ctx, &user); err != nil {
		return errors.InternalServerError(fmt.Sprintf("error upserting user: %s", err.Error()))
	}

	return nil
}

// UpdateProfile implements Usecases.
func (u *usecases) UpdateProfile(ctx context.Context, payload *request.UpdateProfileRequest) error {
	// check if record exists
	profileExisting, err := u.repositories.FindProfileByID(ctx, payload.ID)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error finding profile by id: %s", err.Error()))
	}

	// update profile
	profile := entity.Profile{
		ID:             payload.ID,
		UserID:         profileExisting.UserID,
		Address:        payload.Address,
		District:       payload.District,
		City:           payload.City,
		State:          payload.State,
		Country:        payload.Country,
		Region:         payload.Region,
		Phone:          payload.Phone,
		PersonalID:     payload.PersonalID,
		TypePersonalID: payload.TypePersonalID,
	}

	if err := u.repositories.UpsertProfile(ctx, &profile); err != nil {
		return errors.InternalServerError(fmt.Sprintf("error upserting profile: %s", err.Error()))
	}

	return nil
}

// UpdateUser implements Usecases.
func (u *usecases) UpdateUser(ctx context.Context, payload *request.UpdateUserRequest) error {
	// check if record exists
	userExisting, err := u.repositories.FindUserByID(ctx, payload.ID)
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error finding user by id: %s", err.Error()))
	}

	// update user
	user := entity.User{
		ID:        payload.ID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  userExisting.Password,
	}

	if err := u.repositories.UpsertUser(ctx, &user); err != nil {
		return errors.InternalServerError(fmt.Sprintf("error upserting user: %s", err.Error()))
	}

	return nil
}

// ValidateToken implements Usecases.
func (u *usecases) ValidateToken(ctx context.Context, payload *request.ValidateTokenRequest) error {
	panic("unimplemented")
}

type Usecases interface {
	Register(ctx context.Context, payload *request.RegisterRequest) error
	Login(ctx context.Context, payload *request.LoginRequest) error
	GetUser(ctx context.Context, payload *request.GetUserRequest) error
	UpdateUser(ctx context.Context, payload *request.UpdateUserRequest) error
	ValidateToken(ctx context.Context, payload *request.ValidateTokenRequest) error
	GetProfile(ctx context.Context, payload *request.GetProfileRequest) error
	UpdateProfile(ctx context.Context, payload *request.UpdateProfileRequest) error
}

func New(repositories repositories.Repositories, log log.Logger) Usecases {
	return &usecases{
		repositories: repositories,
		log:          log,
	}
}
