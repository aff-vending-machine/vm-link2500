package usecase

import (
	"context"

	"vm-link2500/internal/layer/usecase/link2500/request"
	"vm-link2500/internal/layer/usecase/link2500/response"
)

type Link2500 interface {
	Test(context.Context) error
	Sale(context.Context, *request.Sale) (*response.Result, error)
	Void(context.Context, *request.Void) (*response.Result, error)
	Refund(context.Context, *request.Refund) (*response.Result, error)
	Settlement(context.Context, *request.Settlement) (*response.Settlement, error)
}
