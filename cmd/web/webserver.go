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
	router := chi.NewRouter()
	webEngine := NewWebEngine()

	router.Handle("/static", http.FileServer(http.Dir("cmd/web/static")))
	router.Get("/", webEngine.Serve)

	srv := &http.Server{
		Addr:    "127.0.0.1:" + PORT,
		Handler: router,
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
