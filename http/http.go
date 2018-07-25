package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/luantranminh/example-go/endpoints"
	categoryDecode "github.com/luantranminh/example-go/http/decode/json/category"
	userDecode "github.com/luantranminh/example-go/http/decode/json/user"
)

// NewHTTPHandler ...
func NewHTTPHandler(endpoints endpoints.Endpoints,
	logger log.Logger,
	useCORS bool) http.Handler {
	r := chi.NewRouter()

	// if running on local (using `make dev`), include cors middleware
	if useCORS {
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})
		r.Use(cors.Handler)
	}

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Get("/_warm", httptransport.NewServer(
		endpoint.Nop,
		httptransport.NopRequestDecoder,
		httptransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllUser,
			userDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{user_id}", httptransport.NewServer(
			endpoints.FindUser,
			userDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateUser,
			userDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{user_id}", httptransport.NewServer(
			endpoints.UpdateUser,
			userDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{user_id}", httptransport.NewServer(
			endpoints.DeleteUser,
			userDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	r.Route("/categories", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllCategory,
			categoryDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{category_id}", httptransport.NewServer(
			endpoints.FindCategory,
			categoryDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateCategory,
			categoryDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{category_id}", httptransport.NewServer(
			endpoints.UpdateCategory,
			categoryDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{category_id}", httptransport.NewServer(
			endpoints.DeleteCategory,
			categoryDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	return r
}
