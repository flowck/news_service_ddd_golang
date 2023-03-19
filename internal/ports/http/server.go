package http

import (
	"context"
	"fmt"
	"net"
	nethttp "net/http"
	"time"

	oapi "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/flowck/news_service_ddd_golang/internal/app"
	"github.com/flowck/news_service_ddd_golang/internal/common/logs"
	"github.com/flowck/news_service_ddd_golang/internal/ports/http/static"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type Port struct {
	logger *logs.Logger
	server *nethttp.Server
}

var _ static.ServerInterface = (*handlers)(nil)

type handlers struct {
	application *app.App
}

func NewPort(
	ctx context.Context,
	port int16,
	allowedCorsOrigin []string,
	application *app.App,
	logger *logs.Logger,
) *Port {
	router := chi.NewRouter()

	swagger, err := static.GetSwagger()
	if err != nil {
		panic(err)
	}

	registerMiddlewares(router, allowedCorsOrigin, logger)
	registerHandlers(router, application, swagger)

	return &Port{
		server: &nethttp.Server{
			Addr:              fmt.Sprintf(":%d", port),
			Handler:           router,
			ReadTimeout:       time.Second * 5,
			ReadHeaderTimeout: time.Second * 5,
			WriteTimeout:      time.Second * 5,
			IdleTimeout:       time.Second * 5,
			BaseContext: func(listener net.Listener) context.Context {
				return ctx
			},
		},
		logger: logger,
	}
}

func (p *Port) Start() {
	p.logger.Infof("Server will run shortly at http://localhost%s", p.server.Addr)
	if err := p.server.ListenAndServe(); err != nil {
		p.logger.Infof("The http server has been shutdown: %v", err)
	}
}

func (p *Port) Stop(ctx context.Context) {
	if err := p.server.Shutdown(ctx); err != nil {
		p.logger.Errorf("An error occurred during http server shutdown: %v", err)
		return
	}

	p.logger.Infof("The http server has been shutdown gracefully")
}

func registerMiddlewares(r *chi.Mux, allowedOrigin []string, logger *logs.Logger) {
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(cors.Handler(corsConfig(allowedOrigin)))

	logger.Infof("CORS enabled for the following origins: %s", allowedOrigin)
}

func registerHandlers(r *chi.Mux, application *app.App, swagger *openapi3.T) {
	h := handlers{application: application}

	r.Get("/openapi.json", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		render.Respond(w, r, swagger)
	})

	r.Route("/", func(r chi.Router) {
		r.Use(oapi.OapiRequestValidatorWithOptions(swagger, &oapi.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: func(ctx context.Context, inp *openapi3filter.AuthenticationInput) error {
					return nil
				},
			},
			ErrorHandler: globalErrorHandler,
		}))
	})

	r.Route("/news", func(router chi.Router) {
		router.Get("/", func(w nethttp.ResponseWriter, rq *nethttp.Request) {
			params, err := mapRequestQueryToPaginationParams(rq)
			if err != nil {
				reply(w, rq, newErrorWithStatus(err, "invalid-query", nethttp.StatusBadRequest))
				return
			}

			h.GetNews(w, rq, params)
		})

		router.Post("/", h.PublishNews)
		router.Route("/{newsID}", func(router chi.Router) {
			router.Get("/", func(w nethttp.ResponseWriter, r *nethttp.Request) {
				h.GetNewsByID(w, r, chi.URLParam(r, "newsID"))
			})

			router.Patch("/unpublish", func(w nethttp.ResponseWriter, r *nethttp.Request) {
				h.UnPublishNews(w, r, chi.URLParam(r, "newsID"))
			})
		})
	})

	r.Route("/topics", func(router chi.Router) {
		router.Get("/", h.GetTopics)
		router.Post("/", h.CreateTopic)
		router.Route("/{topicID}", func(r chi.Router) {
			r.Get("/", func(w nethttp.ResponseWriter, req *nethttp.Request) {
				h.GetTopicByID(w, req, chi.URLParam(req, "topicID"))
			})

			r.Delete("/", func(w nethttp.ResponseWriter, r *nethttp.Request) {
				h.RemoveTopicByID(w, r, chi.URLParam(r, "topicID"))
			})

			r.Put("/", func(w nethttp.ResponseWriter, r *nethttp.Request) {
				h.EditTopic(w, r, chi.URLParam(r, "topicID"))
			})
		})
	})
}

// func strToFloat(value string) float32 {}
