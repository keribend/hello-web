package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"sandbox/hello-web/ui"
)

func cliExecute(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/", ui.Handler())
	mux.HandleFunc("/api/quotes", Server)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
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

	log.Println("the server is listening to port :8080")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

// type ServerResponse struct {
// 	Quote Quote `json:"quote"`
// }

var quotes = []Quote{
	{
		Content: "I'm not here to be perfect, I'm here to be real.",
		Author:  "Lady Gaga",
	},
	{
		Content: "I'm not interested in money. I just want to be wonderful.",
		Author:  "Marilyn Monroe",
	},
	{
		Content: "The only thing that feels better than winning is winning when nobody thought you could.",
		Author:  "Hank Aaron",
	},
	{
		Content: "Success is not final, failure is not fatal: It is the courage to continue that counts.",
		Author:  "Winston Churchill",
	},
	{
		Content: "If you can dream it, you can do it.",
		Author:  "Walt Disney",
	},
	{
		Content: "If you want something done, ask a busy person to do it.",
		Author:  "Laura Ingalls Wilder",
	},
	{
		Content: "If your actions inspire others to dream more, learn more, do more and become more, you are a leader.",
		Author:  "John Quincy Adams",
	},
	{
		Content: "The best way to find out if you can trust somebody is to trust them.",
		Author:  "Ernest Hemingway",
	},
	{
		Content: "The only Limit to our realization of tomorrow will be our doubts of today.",
		Author:  "Franklin D. Roosevelt",
	},
	{
		Content: "We may encounter many defeats but we must not be defeated.",
		Author:  "Maya Angelou",
	},
}

func Server(w http.ResponseWriter, r *http.Request) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	resp := quotes[r1.Intn(len(quotes))]

	output, _ := json.Marshal(resp)

	fmt.Fprint(w, string(output))
}
