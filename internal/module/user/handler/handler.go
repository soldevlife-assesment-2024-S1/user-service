package handler

import (
	"strconv"
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
	var req request.Register
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	// call usecase
	if err := h.Usecase.Register(ctx.UserContext(), &req); err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return helpers.RespSuccess(ctx, h.Log, nil, "register success")
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var req request.Login
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	// call usecase
	resp, err := h.Usecase.Login(ctx.UserContext(), &req)
	if err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return helpers.RespSuccess(ctx, h.Log, resp, "login success")
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	var req request.GetUser

	req.ID = ctx.Locals("userID").(int)

	// call usecase
	resp, err := h.Usecase.GetUser(ctx.UserContext(), &req)
	if err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return helpers.RespSuccess(ctx, h.Log, resp, "get user success")
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var req request.UpdateUser
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	// call usecase
	if err := h.Usecase.UpdateUser(ctx.UserContext(), &req); err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return helpers.RespSuccess(ctx, h.Log, nil, "update user success")
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	var req request.CreateProfile

	if err := ctx.BodyParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	// call usecase
	if err := h.Usecase.CreateProfile(ctx.UserContext(), &req); err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return helpers.RespSuccess(ctx, h.Log, nil, "create profile success")
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	var req request.GetProfile

	req.UserID = ctx.Locals("userID").(int)

	// call usecase
	resp, err := h.Usecase.GetProfile(ctx.UserContext(), &req)
	if err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return helpers.RespSuccess(ctx, h.Log, resp, "get profile success")
}

func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	var req request.UpdateProfile
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	// call usecase
	if err := h.Usecase.UpdateProfile(ctx.UserContext(), &req); err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return helpers.RespSuccess(ctx, h.Log, nil, "update profile success")
}

// private
func (h *UserHandler) ValidateToken(ctx *fiber.Ctx) error {
	var req request.ValidateToken
	if err := ctx.QueryParser(&req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}

	// validate request
	if err := h.Validator.Struct(req); err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest(err.Error()))
	}

	// call usecase
	resp, err := h.Usecase.ValidateToken(ctx.UserContext(), &req)
	if err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return helpers.RespSuccess(ctx, h.Log, resp, "validate token success")

}

func (h *UserHandler) GetProfilePrivate(ctx *fiber.Ctx) error {
	userID := ctx.Query("user_id")
	var req request.GetProfile

	uid, err := strconv.Atoi(userID)
	if err != nil {
		return helpers.RespError(ctx, h.Log, errors.BadRequest("bad request"))
	}
	req.UserID = uid
	// call usecase
	resp, err := h.Usecase.GetProfile(ctx.UserContext(), &req)
	if err != nil {
		return helpers.RespError(ctx, h.Log, err)
	}

	return helpers.RespSuccess(ctx, h.Log, resp, "get profile private success")
}
