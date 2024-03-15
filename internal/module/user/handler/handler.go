package handler

import (
	"user-service/internal/module/user/models/request"
	"user-service/internal/module/user/usecases"
	"user-service/internal/pkg/helpers"
	"user-service/internal/pkg/helpers/errors"
	"user-service/internal/pkg/log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Log       log.Logger
	Validator *validator.Validate
	Usecase   usecases.Usecases
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var req request.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	// call usecase
	if err := h.Usecase.Register(ctx.Context(), &req); err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return nil
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var req request.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	return nil
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	var req request.GetUserRequest
	if err := ctx.QueryParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	return nil
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var req request.UpdateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	// call usecase
	if err := h.Usecase.UpdateUser(ctx.Context(), &req); err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return nil
}

// private
func (h *UserHandler) ValidateToken(ctx *fiber.Ctx) error {
	var req request.ValidateTokenRequest
	if err := ctx.QueryParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	return nil
}
