package http

import (
	"github.com/ggicci/httpin/integration"
	"github.com/go-chi/chi/v5"
	http2 "github.com/go-modulus/modulus/http"
	"github.com/go-modulus/modulus/module"
)

func OverrideHttpRouter(m *module.Module) *module.Module {
	return http2.OverrideRouter[chi.Router](m)
}

func NewModule() *module.Module {
	return module.NewModule("modulus/chihttp").
		AddDependencies(
			http2.NewModule(),
		).
		AddProviders(
			NewRouter,
		)
}

func NewManifesto() module.Manifesto {
	httpModule := module.NewManifesto(
		NewModule(),
		"github.com/go-modulus/chihttp",
		"Chi HTTP router that is working over the base http modulus module.",
		"1.0.0",
	)

	return httpModule
}

func init() {
	integration.UseGochiURLParam("path", chi.URLParam)
}
