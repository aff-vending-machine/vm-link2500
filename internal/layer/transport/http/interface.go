package http

import "github.com/gofiber/fiber/v2"

type Link2500 interface {
	Test(*fiber.Ctx) error       // POST {link2500/test}
	Sale(*fiber.Ctx) error       // POST {link2500/sale}
	Void(*fiber.Ctx) error       // POST {link2500/void}
	Refund(*fiber.Ctx) error     // POST {link2500/refund}
	Settlement(*fiber.Ctx) error // POST {link2500/settlement}
}
