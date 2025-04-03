package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/redrum15/prueba/src/db"
	"github.com/redrum15/prueba/src/handlers"
	"github.com/redrum15/prueba/src/middlewares"
	"github.com/redrum15/prueba/src/services/auth"
	"github.com/redrum15/prueba/src/services/vms"
)

const ADMIN_ROLE = "admin"

func main() {
	db.Init()
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	router.Get("/ws", handlers.HandleWebSocket)

	go func() {
		for vm := range handlers.GetBroadcast() {
			handlers.NotifyClients(vm)
		}
	}()

	router.Route("/api", func(router chi.Router) {

		router.Post("/login", auth.Login)

		router.Group(func(router chi.Router) {
			router.Use(middlewares.JWTVerifier())
			router.Use(middlewares.RequireRole(ADMIN_ROLE))

			router.Post("/vms", vms.CreateVM)
			router.Delete("/vms/{id}", vms.DeleteVM)
			router.Put("/vms/{id}", vms.UpdateVM)
		})

		router.Group(func(router chi.Router) {
			router.Use(middlewares.JWTVerifier())

			router.Get("/vms", vms.ListVMS)
			router.Get("/vms/{id}", vms.DetailVM)
		})
	})

	http.ListenAndServe(":3000", router)
}
