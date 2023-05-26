package gin

import (
	"vm-link2500/internal/layer/transport/http"

	"github.com/gin-gonic/gin"
)

func routeLink2500(api *gin.RouterGroup, endpoint http.Link2500) {
	api.POST("link2500/test", endpoint.Test)
	api.POST("link2500/sale", endpoint.Sale)
	api.POST("link2500/void", endpoint.Void)
	api.POST("link2500/refund", endpoint.Refund)
	api.POST("link2500/settlement", endpoint.Settlement)
}
