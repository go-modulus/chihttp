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

func init() {
	integration.UseGochiURLParam("path", chi.URLParam)
}
