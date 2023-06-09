package link2500

import (
	"vm-link2500/internal/layer/usecase"
)

type httpImpl struct {
	usecase usecase.Link2500
}

func New(uc usecase.Link2500) *httpImpl {
	return &httpImpl{uc}
}
