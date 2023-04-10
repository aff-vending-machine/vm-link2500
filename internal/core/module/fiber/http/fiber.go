package http

import (
	"fmt"

	"github.com/aff-vending-machine/vm-link2500/pkg/utils/errs"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// 200 - OK
func OK(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "done",
		"data":   data,
	})
}

// 204 - NoContent
func NoContent(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusNoContent,
		"status": "done",
	})
}

// 201 - Created
func Created(ctx *fiber.Ctx, id string) error {
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"id":     id,
		"status": "done",
	})
}

func UsecaseError(ctx *fiber.Ctx, err error) error {
	code, msg := translateError(err)
	return ctx.Status(code).JSON(fiber.Map{
		"code":    code,
		"status":  "error",
		"message": msg,
	})
}

// 400 - Bad Request
func BadRequest(ctx *fiber.Ctx, cause error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"code":    fiber.StatusBadRequest,
		"status":  "error",
		"message": cause.Error(),
	})
}

// 401 - Unauthorized
func Unauthorized(ctx *fiber.Ctx, cause error) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    fiber.StatusUnauthorized,
		"status":  "error",
		"message": cause.Error(),
	})
}

// 403 - Forbidden
func Forbidden(ctx *fiber.Ctx, cause error) error {
	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"code":    fiber.StatusForbidden,
		"status":  "error",
		"message": cause.Error(),
	})
}

func translateError(err error) (int, string) {
	// 400
	if errs.Is(err, "invalid request") {
		return fiber.StatusBadRequest, errors.Cause(err).Error()
	}

	if errs.Is(err, "exist") {
		return fiber.StatusBadRequest, errors.Cause(err).Error()
	}

	if errs.Is(err, "device id") {
		return fiber.StatusBadRequest, "device ID is invalid"
	}

	if errs.Is(err, "decrypt") {
		return fiber.StatusBadRequest, "data is invalid"
	}

	if errs.Is(err, "invalid data") {
		return fiber.StatusBadRequest, "data is invalid"
	}

	// 403
	if errs.Is(err, "signature") {
		return fiber.StatusBadRequest, errors.Cause(err).Error()
	}

	if errors.Is(err, fiber.ErrForbidden) {
		return fiber.StatusForbidden, "no permission"
	}

	if errs.Is(err, "no permission") {
		return fiber.StatusForbidden, err.Error()
	}

	// 404
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.StatusNotFound, "no data"
	}

	// 500
	if errs.Is(err, "rpc error") {
		return fiber.StatusInternalServerError, errors.Cause(err).Error()
	}

	return fiber.StatusBadRequest, fmt.Sprintf("unexpected error: (%s)", err.Error())
}
