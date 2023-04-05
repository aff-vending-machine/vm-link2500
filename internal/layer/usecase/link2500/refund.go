package link2500

import (
	"context"

	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/request"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Refund(ctx context.Context, req *request.Refund) (*response.Result, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	res, err := uc.link2500.Refund(ctx, req)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to refund EDC")
	}

	return res, nil
}
