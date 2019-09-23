package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CorsConfig struct {
	AllowedMethods  []string
	AllowedHeaders  []string
	WithCredentials bool
}

var defaultAllowedMethods = []string{
	http.MethodGet,
	http.MethodDelete,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
}

const defaultOrigin = "*"

func main() {
	cc := &CorsConfig{
		AllowedMethods:  defaultAllowedMethods,
		AllowedHeaders:  defaultAllowedHeaders,
		WithCredentials: false,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleCorsRequest(cc))

	server := &http.Server{
		Addr:           ":5656",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

func handleCorsRequest(c *CorsConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		w.Header().Set("Access-Control-Allow-Origin", defaultOrigin)

		// Check config for options here and set appropriate headers

		body := proxyRequest(http.DefaultClient, r, ctx)
		fmt.Fprintf(w, "%s", body)
	}
}

func proxyRequest(client *http.Client, r *http.Request, ctx context.Context) string {
	// Perform request and read body here
}
