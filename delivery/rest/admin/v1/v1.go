package v1

import (
	"mona/delivery/router"
	"mona/domain/service"
	"net/http"

	"github.com/rs/zerolog"
)

type AdminV1 struct {
	Svc    *service.ServiceFacade
	Logger zerolog.Logger
}

func New(svc *service.ServiceFacade, logger zerolog.Logger) *AdminV1 {
	return &AdminV1{
		Svc:    svc,
		Logger: logger.With().Str("component", "AdminV1").Logger(),
	}
}

func (a *AdminV1) InitRoutes(mux *http.ServeMux) {
	adminsV1 := router.NewGroup(mux, "admin/v1", nil)
	{
		_ = adminsV1
		// adminsV1.Get("/example", a.exampleHandler)
	}
}
