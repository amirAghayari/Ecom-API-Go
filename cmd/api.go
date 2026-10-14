package main

import (
	"log"
	"net/http"
	"time"

	"github.com/amiraghayari/Ecom-API-Go/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// middleware
	r.Use(middleware.RequestID) // important for ratelimiting
	r.Use(middleware.RealIP)    // important for ratelimiting and analytics , tracing
	r.Use(middleware.Logger)    //
	r.Use(middleware.Recoverer) // recover from crashes

	r.Use(middleware.Timeout((60 * time.Second)))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all Goooooood!"))
	})

	productService := products.NewService()
	productsHandler := products.NewHandler(productService)

	r.Get("/products", productsHandler.ListProducts)

	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	// logger
	// db driver
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
