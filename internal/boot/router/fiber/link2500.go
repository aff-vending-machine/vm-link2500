package fiber

import (
	"github.com/aff-vending-machine/vm-link2500/internal/layer/transport/http"
	"github.com/gofiber/fiber/v2"
)

func routeLink2500(api fiber.Router, endpoint http.Link2500) {
	api.Post("link2500/test", endpoint.Test)
	api.Post("link2500/sale", endpoint.Sale)
	api.Post("link2500/void", endpoint.Void)
	api.Post("link2500/refund", endpoint.Refund)
	api.Post("link2500/settlement", endpoint.Settlement)
}
