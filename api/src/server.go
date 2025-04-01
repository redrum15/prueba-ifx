package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redrum15/prueba/src/db"
	"github.com/redrum15/prueba/src/middlewares"
	"github.com/redrum15/prueba/src/services/auth"
	"github.com/redrum15/prueba/src/services/vms"
)

const ADMIN_ROLE = "admin"

func main() {
	db.Init()
	router := chi.NewRouter()

	router.Use(middleware.Logger)

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

	http.ListenAndServe(":3000", router)
}
