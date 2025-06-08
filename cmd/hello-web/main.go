package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "modernc.org/sqlite"

	"github.com/keribend/hello-web/internal/controller"
	"github.com/keribend/hello-web/internal/http/routing"
	"github.com/keribend/hello-web/internal/repository"
	"github.com/keribend/hello-web/internal/service"
)

var DB *sql.DB

func init() {
	initDB()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log.Println(bootstrapApp(ctx))
}

func bootstrapApp(ctx context.Context) error {
	repo := repository.New(DB)
	service := service.New(repo)
	htmlController := controller.NewHTMLController(service)
	apiController := controller.NewAPIController(service)

	mainRouter := http.NewServeMux()
	uiRouter := routing.NewUIRouter(htmlController)
	apiRouter := routing.NewAPIRouter(apiController)

	mainRouter.Handle("/api/", http.StripPrefix("/api", apiRouter))
	mainRouter.Handle("/ui/", http.StripPrefix("/ui", uiRouter))
	mainRouter.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./public/static"))))

	server := http.Server{
		Addr:    ":8080",
		Handler: mainRouter,
	}

	go gracefulShutdown(ctx, &server)

	log.Println("Server listening on port 8080")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func gracefulShutdown(ctx context.Context, server *http.Server) {
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer DB.Close()

	log.Println("server shutting down (reason: context_canceled)")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
		return
	}
	log.Println("graceful shutdown complete")
}

func initDB() {
	var err error
	DB, err = sql.Open("sqlite", "file:./bin/hello-web.db?_foreign_keys=true")
	if err != nil {
		log.Fatal(err)
	}
}
