package link2500

import (
	"github.com/aff-vending-machine/vm-link2500/internal/core/module/fiber/http"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/request"
	"github.com/aff-vending-machine/vm-link2500/pkg/trace"
	"github.com/gofiber/fiber/v2"
)

func (t *httpImpl) Settlement(c *fiber.Ctx) error {
	ctx, span := trace.Start(c.Context())
	defer span.End()

	req, err := makeSettlementRequest(c)
	if err != nil {
		trace.RecordError(span, err)
		return http.BadRequest(c, err)
	}

	// usecase execution
	res, err := t.usecase.Settlement(ctx, req)
	if err != nil {
		trace.RecordError(span, err)
		return http.UsecaseError(c, err)
	}

	return http.OK(c, res)
}

func makeSettlementRequest(c *fiber.Ctx) (*request.Settlement, error) {
	var req request.Settlement
	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
