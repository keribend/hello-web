package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	_ "modernc.org/sqlite"

	"github.com/keribend/hello-web/internal/controller"
	"github.com/keribend/hello-web/internal/http/router"
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
	controller := controller.New(service)

	r := router.New()
	r.Get("/", controller.EventList)
	r.Post("/events/{eventId}/checkin", controller.AddCheckinToEvent)
	r.Get("/checkins", controller.Checkins)

	// Create a route along /assets that will serve contents from the ./ui/assets/ folder
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public/static"))
	FileServer(r, "/static", filesDir)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
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

// FileServer conveniently sets up a http.FileServer handler to serve static files from a http.FileSystem
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		log.Panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func initDB() {
	var err error
	
	DB, err = sql.Open("sqlite", "file:./bin/hello-web.db?_foreign_keys=true")
	if err != nil {
		log.Fatal(err)
	}
}
