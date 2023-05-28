package gin

import (
	"vm-link2500/internal/boot/modules"

	"github.com/rs/zerolog/log"
)

func (s *server) Serve(driver modules.HTTPTransport) {
	v1 := s.Engine.Group("/api/v1")
	routeLink2500(v1, driver.Link2500)

	go func() {
		err := s.Engine.Run(s.Address)
		if err != nil {
			log.Error().Err(err).Str("address", s.Address).Msg("unable to start http server")
		}
	}()

	log.Info().Str("address", s.Address).Msg("http server listen")
}
