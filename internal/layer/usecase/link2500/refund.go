package link2500

import (
	"context"

	"vm-link2500/internal/layer/usecase/link2500/request"
	"vm-link2500/internal/layer/usecase/link2500/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Refund(ctx context.Context, req *request.Refund) (*response.Result, error) {
	if v := validate.Struct(req); !v.Validate() {
		log.Error().Interface("req", req).Msg("unable to validate request")
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	res, err := uc.link2500.Refund(ctx, req)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("unable to refund EDC")
		return nil, errors.Wrapf(err, "unable to refund EDC")
	}

	return res, nil
}
