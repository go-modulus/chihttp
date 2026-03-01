package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-modulus/modulus/errors/erruser"
	"github.com/go-modulus/modulus/errors/errwrap"
	http2 "github.com/go-modulus/modulus/http"
	"github.com/go-modulus/modulus/http/errhttp"
	"github.com/go-modulus/modulus/module"
)

var (
	ErrMethodNotAllowed = errwrap.Wrap(
		erruser.New("MethodNotAllowed", "Method not allowed"),
		errhttp.With(http.StatusMethodNotAllowed),
	)
	ErrNotFound = errwrap.Wrap(
		erruser.New("NotFound", "Not found"),
		errhttp.With(http.StatusNotFound),
	)
)

func NewRouter(errorPipeline *errhttp.ErrorPipeline, config ServeConfig) chi.Router {
	r := chi.NewRouter()
	r.MethodNotAllowed(
		errhttp.WrapHandler(
			errorPipeline,
			func(w http.ResponseWriter, req *http.Request) error {
				return ErrMethodNotAllowed
			},
		),
	)
	r.NotFound(
		errhttp.WrapHandler(
			errorPipeline,
			func(w http.ResponseWriter, req *http.Request) error {
				return ErrNotFound
			},
		),
	)
	if config.TTL > 0 {
		r.Use(chiMiddleware.Timeout(config.TTL))
	}
	if config.RequestSizeLimit > 0 {
		r.Use(chiMiddleware.RequestSize(int64(config.RequestSizeLimit.Bytes())))
	}
	return r
}

func NewModule() *module.Module {
	return module.NewModule("modulus/chi-http").
		AddDependencies(http2.NewModule()).
		AddCliCommands(
			NewServeCommand,
		).
		AddProviders(
			NewRouter,
			NewServe,
		).
		InitConfig(ServeConfig{})
}
