package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const PORT = "8080"

func webServerListen(ctx context.Context) error {
	r := chi.NewRouter()
	webEngine := NewWebEngine()

	// Business logic routes
	r.Route("/counters", func(r chi.Router) {
		r.Get("/counters", webEngine.GetCounters)
		r.Route("/{counterId:[0-9]+}", func(r chi.Router) {
			// r.Get("/", webEngine.GetCounter)
			r.Post("/increase", webEngine.IncreaseCounter)
			r.Post("/decrease", webEngine.DecreaseCounter)
		})
	})

	// Serve the main page
	r.Get("/", webEngine.Serve)

	// Serve static assets
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("cmd/web/static/"))))

	srv := &http.Server{
		Addr:    "localhost:" + PORT,
		Handler: r,
	}

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Println("server shutting down (reason: context_cancelled)")
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
			return
		}
		log.Printf("graceful shutdown complete")
	}()

	log.Println("the server is listening to port :" + PORT)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
