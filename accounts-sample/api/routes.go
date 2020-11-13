package api

import (
	"github.com/go-chi/chi"
)

// SetupRoutes contains all http route definitions
func SetupRoutes(r *chi.Mux, wa *WebApp) error {
	var err error

	r.Get("/", wa.Home)
	r.Route("/v1", func(ra chi.Router) {
		ra.Use(wa.AuthMiddleware)
		ra.Post("/account", wa.AddAccount)
		ra.Post("/deposit/{accountID}", wa.DepositFunds)
		ra.Get("/account/{accountID}", wa.GetAccount)
		ra.Post("/transfer", wa.TransferCash)
		ra.Get("/accounts/{clientID}", wa.GetClientAccounts)
		ra.Post("/accounts", wa.PostClientAccounts)
	})

	r.Route("/auth", func(rt chi.Router) {
		rt.Get("/", wa.Ping)
	})

	return err
}
