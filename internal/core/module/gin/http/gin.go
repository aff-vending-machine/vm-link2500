package http

import (
	"fmt"
	"net/http"

	"vm-link2500/pkg/utils/errs"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// 200 - OK
func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":   http.StatusOK,
		"status": "done",
		"data":   data,
	})
}

// 204 - NoContent
func NoContent(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, map[string]interface{}{
		"code":   http.StatusNoContent,
		"status": "done",
	})
}

// 201 - Created
func Created(ctx *gin.Context, id string) {
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"code":   http.StatusCreated,
		"id":     id,
		"status": "done",
	})
}

func UsecaseError(ctx *gin.Context, err error) {
	code, msg := translateError(err)
	ctx.AbortWithStatusJSON(code, map[string]interface{}{
		"code":    code,
		"status":  "error",
		"message": msg,
	})
}

// 400 - Bad Request
func BadRequest(ctx *gin.Context, cause error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"status":  "error",
		"message": cause.Error(),
	})
}

// 401 - Unauthorized
func Unauthorized(ctx *gin.Context, cause error) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "error",
		"message": cause.Error(),
	})
}

// 403 - Forbidden
func Forbidden(ctx *gin.Context, cause error) {
	ctx.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
		"code":    http.StatusForbidden,
		"status":  "error",
		"message": cause.Error(),
	})
}

func translateError(err error) (int, string) {
	// 400
	if errs.Is(err, "invalid request") {
		return http.StatusBadRequest, errors.Cause(err).Error()
	}

	if errs.Is(err, "exist") {
		return http.StatusBadRequest, errors.Cause(err).Error()
	}

	if errs.Is(err, "device id") {
		return http.StatusBadRequest, "device ID is invalid"
	}

	if errs.Is(err, "decrypt") {
		return http.StatusBadRequest, "data is invalid"
	}

	if errs.Is(err, "invalid data") {
		return http.StatusBadRequest, "data is invalid"
	}

	// 403
	if errs.Is(err, "signature") {
		return http.StatusBadRequest, errors.Cause(err).Error()
	}

	if errs.Is(err, "forbidden") {
		return http.StatusForbidden, "no permission"
	}

	if errs.Is(err, "no permission") {
		return http.StatusForbidden, err.Error()
	}

	// 404
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, "no data"
	}

	// 500
	if errs.Is(err, "rpc error") {
		return http.StatusInternalServerError, errors.Cause(err).Error()
	}

	return http.StatusBadRequest, fmt.Sprintf("unexpected error: (%s)", err.Error())
}
