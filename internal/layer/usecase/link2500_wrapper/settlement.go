package link2500_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/request"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/response"
	"github.com/aff-vending-machine/vm-link2500/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Settlement(ctx context.Context, req *request.Settlement) (*response.Settlement, error) {
	ctx, span := trace.Start(ctx)
	defer span.End()

	res, err := w.usecase.Settlement(ctx, req)
	if err != nil {
		log.Error().Interface("request", req).Err(err).Msg("unable to settlement EDC")
		trace.RecordError(span, err)
	}

	return res, err
}
