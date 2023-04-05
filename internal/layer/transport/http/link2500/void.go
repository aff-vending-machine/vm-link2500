package link2500

import (
	"github.com/aff-vending-machine/vm-link2500/internal/core/infra/fiber/http"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/request"
	"github.com/aff-vending-machine/vm-link2500/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (t *httpImpl) Void(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	req, err := makeVoidRequest(c)
	if err != nil {
		trace.RecordError(span, err)
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := t.usecase.Void(ctx, req)
	if err != nil {
		trace.RecordError(span, err)
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}

func makeVoidRequest(c *fiber.Ctx) (*request.Void, error) {
	var req request.Void
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
