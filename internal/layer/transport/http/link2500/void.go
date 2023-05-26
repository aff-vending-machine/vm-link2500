package link2500

import (
	"context"

	"vm-link2500/internal/core/module/gin/http"
	"vm-link2500/internal/layer/usecase/link2500/request"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (t *httpImpl) Void(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	req, err := makeVoidRequest(c)
	if err != nil {
		log.Error().Err(err).Msg("unable to parse request")
		http.BadRequest(c, err)
		return
	}

	// usecase execution
	res, err := t.usecase.Void(ctx, req)
	if err != nil {
		http.UsecaseError(c, err)
		return
	}

	http.OK(c, res)
}

func makeVoidRequest(c *gin.Context) (*request.Void, error) {
	var req request.Void
	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
