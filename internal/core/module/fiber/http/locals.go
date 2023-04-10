package http

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func String(ctx *fiber.Ctx, key string) string {
	if ctx.Locals(key) == nil {
		return ""
	}

	casted, ok := ctx.Locals(key).(string)
	if !ok {
		return ""
	}

	return casted
}

func Int(ctx *fiber.Ctx, key string) int {
	if ctx.Locals(key) == nil {
		return 0
	}

	casted, ok := ctx.Locals(key).(int)
	if !ok {
		str := fmt.Sprintf("%v", ctx.Locals(key))
		parsed, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return parsed
	}

	return casted
}
