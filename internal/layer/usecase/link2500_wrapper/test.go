package link2500_wrapper

import (
	"context"

	"github.com/aff-vending-machine/vm-link2500/pkg/trace"
	"github.com/rs/zerolog/log"
)

func (w *wrapperImpl) Test(ctx context.Context) error {
	ctx, span := trace.Start(ctx)
	defer span.End()

	err := w.usecase.Test(ctx)
	if err != nil {
		log.Error().Err(err).Msg("unable to test EDC")
		trace.RecordError(span, err)
	}

	return err
}
