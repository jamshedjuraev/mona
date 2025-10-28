package v1

import (
	"mona/delivery/router"
	"mona/domain/service"
	"net/http"

	"github.com/rs/zerolog"
)

type PaymentV1 struct {
	Svc    *service.ServiceFacade
	Logger zerolog.Logger
}

func New(svc *service.ServiceFacade, logger zerolog.Logger) *PaymentV1 {
	return &PaymentV1{
		Svc:    svc,
		Logger: logger.With().Str("component", "PaymentV1").Logger(),
	}
}

func (p *PaymentV1) InitRoutes(mux *http.ServeMux) {
	paymentsV1 := router.NewGroup(mux, "payments/v1", nil)
	{
		_ = paymentsV1
		// paymentsV1.Get("/example", p.exampleHandler)
	}
}