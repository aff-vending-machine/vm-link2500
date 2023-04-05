package fiber

import (
	"github.com/aff-vending-machine/vm-link2500/internal/app/registry"
	"github.com/rs/zerolog/log"
)

func (s *server) Serve(driver registry.AppHTTPDriver) {
	v1 := s.App.Group("/api/v1")
	routeLink2500(v1, driver.Link2500)

	go s.App.Listen(s.Address)

	log.Info().Str("address", s.Address).Msg("http server listen")
}
