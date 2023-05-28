package link2500

import (
	"context"

	"vm-link2500/internal/core/infra/network/gin/http"
	"vm-link2500/internal/layer/usecase/link2500/request"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (t *httpImpl) Refund(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	req, err := makeRefundRequest(c)
	if err != nil {
		log.Error().Err(err).Msg("unable to parse request")
		http.BadRequest(c, err)
		return
	}

	// usecase execution
	res, err := t.usecase.Refund(ctx, req)
	if err != nil {
		http.UsecaseError(c, err)
		return
	}

	http.OK(c, res)
}

func makeRefundRequest(c *gin.Context) (*request.Refund, error) {
	var req request.Refund
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
