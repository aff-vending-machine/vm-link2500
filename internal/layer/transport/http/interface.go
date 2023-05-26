package http

import (
	"github.com/gin-gonic/gin"
)

type Link2500 interface {
	Refund(*gin.Context)     // POST {link2500/refund}
	Sale(*gin.Context)       // POST {link2500/sale}
	Settlement(*gin.Context) // POST {link2500/settlement}
	Test(*gin.Context)       // POST {link2500/test}
	Void(*gin.Context)       // POST {link2500/void}
}
