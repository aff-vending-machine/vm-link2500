package fiber

import "github.com/aff-vending-machine/vm-link2500/internal/core/infra/fiber"

type server struct {
	*fiber.Wrapper
}

func New(client *fiber.Wrapper) *server {
	return &server{
		client,
	}
}
