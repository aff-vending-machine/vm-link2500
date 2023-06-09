package link2500

import (
	"context"
	"time"

	"vm-link2500/internal/layer/usecase/link2500/request"
	"vm-link2500/internal/layer/usecase/link2500/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Sale(ctx context.Context, req *request.Sale) (*response.Result, error) {
	if v := validate.Struct(req); !v.Validate() {
		log.Error().Interface("req", req).Msg("unable to validate request")
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	c, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	res, err := uc.link2500.Sale(c, req)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("unable to sale EDC")
		return nil, errors.Wrapf(err, "unable to sale EDC")
	}

	return res, nil
}
