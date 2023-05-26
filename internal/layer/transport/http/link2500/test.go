package link2500

import (
	"context"

	"vm-link2500/internal/core/module/gin/http"

	"github.com/gin-gonic/gin"
)

func (t *httpImpl) Test(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	// usecase execution
	err := t.usecase.Test(ctx)
	if err != nil {
		http.UsecaseError(c, err)
		return
	}

	http.NoContent(c)
}
