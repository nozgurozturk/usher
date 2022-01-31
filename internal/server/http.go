package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nozgurozturk/usher/internal/application"
	admin "github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/admin/handler"
	public "github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/public/handler"
)

type HttpServer struct {
	router chi.Router
}

func NewHttpServer(app *application.Application) *HttpServer {

	rootRouter := chi.NewRouter()

	rootRouter.Use(
		middleware.Logger,
		middleware.Recoverer)

	rootRouter.Mount("/api/v1", public.NewAPIHandler(app))
	rootRouter.Mount("/admin/v1", admin.NewAPIHandler(app))

	return &HttpServer{
		router: rootRouter,
	}
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
