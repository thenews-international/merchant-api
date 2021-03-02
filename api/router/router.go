package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"merchant/api/handler"
	"merchant/api/router/middleware"
)

func New(srv *handler.Server) *chi.Mux {
	r := chi.NewRouter()

	// Prepare CORS.
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "pragma", "X-Organization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Total-Count"},
		MaxAge:           300,
	})

	// Route to swagger specification

	r.Route("/auth", func(r chi.Router) {
		r.Use(cors.Handler)
		r.Use(middleware.ContentTypeJson)

		r.MethodFunc(http.MethodPost, "/login", srv.HandleLogin)
		r.MethodFunc(http.MethodPost, "/register", srv.HandleRegister)
	})

	// Routes for APIs
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(cors.Handler)
		r.Use(middleware.ContentTypeJson)
		r.Use(middleware.JwtAuthentication)

		// Routes for merchants
		r.MethodFunc(http.MethodGet, "/merchants", srv.HandleListMerchant)
		r.MethodFunc(http.MethodGet, "/merchants/{id}", srv.HandleReadMerchant)
		r.MethodFunc(http.MethodPut, "/merchants/{id}", srv.HandleUpdateMerchant)
		r.MethodFunc(http.MethodDelete, "/merchants/{id}", srv.HandleDeleteMerchant)

		// Routes for team members
		r.MethodFunc(http.MethodGet, "/team-members", srv.HandleListTeamMember)
		r.MethodFunc(http.MethodPost, "/team-members", srv.HandleCreateTeamMember)
		r.MethodFunc(http.MethodGet, "/team-members/{id}", srv.HandleReadTeamMember)
		r.MethodFunc(http.MethodPut, "/team-members/{id}", srv.HandleUpdateTeamMember)
		r.MethodFunc(http.MethodDelete, "/team-members/{id}", srv.HandleDeleteTeamMember)
	})

	return r
}
