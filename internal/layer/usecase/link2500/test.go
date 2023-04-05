package link2500

import (
	"context"

	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/request"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Test(ctx context.Context) error {
	_, err := uc.link2500.Sale(ctx, &request.Sale{
		MerchantID: "000001",
		Price:      1,
	})
	if err != nil {
		return errors.Wrapf(err, "unable to test EDC")
	}

	return nil
}
