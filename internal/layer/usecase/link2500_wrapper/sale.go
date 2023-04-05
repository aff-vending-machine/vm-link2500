package link2500_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/request"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/response"
	"github.com/aff-vending-machine/vm-link2500/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Sale(ctx context.Context, req *request.Sale) (*response.Result, error) {
	ctx, span := trace.Start(ctx)
	defer span.End()

	res, err := w.usecase.Sale(ctx, req)
	if err != nil {
		log.Error().Interface("request", req).Err(err).Msg("unable to sale EDC")
		trace.RecordError(span, err)
	}

	return res, err
}
