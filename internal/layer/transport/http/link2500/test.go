package link2500

import (
	"github.com/aff-vending-machine/vm-link2500/internal/core/infra/fiber/http"
	"github.com/aff-vending-machine/vm-link2500/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (t *httpImpl) Test(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	// usecase execution
	err := t.usecase.Test(ctx)
	if err != nil {
		trace.RecordError(span, err)
		return http.UsecaseError(c, err)
	}

	return http.NoContent(c)
}
