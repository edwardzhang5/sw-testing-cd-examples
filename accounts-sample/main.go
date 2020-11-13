package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/drbyronw/accounts/api"
	"github.com/drbyronw/accounts/db"
	"github.com/drbyronw/accounts/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

const (
	accounts = "accounts"
)

func initMain() {
	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Printf("Could not load environment: %v", err)
		}

	}
}

func main() {
	fmt.Println("SW Testing Accounts")
	initMain()
	if os.Getenv("ENVIRONMENT") == "staging" {
		log.Println("---- Staging Environment ---- ")
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use()

	// 30 second timeout on request context
	r.Use(middleware.Timeout(10 * time.Second))

	// API Version
	r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
		js, _ := json.Marshal(map[string]interface{}{
			"api_version": "1.0.1",
		})
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(js)
		if err != nil {
			http.Error(w, "unable to response to api version request", http.StatusInternalServerError)
			return
		}
	})

	wa, err := newWebApp()
	if err != nil {
		log.Fatalf("Unable to setup Web App Service: %v\n", err)
	}

	defer func() {
		err := wa.DB.Client.Close()
		if err != nil {
			log.Fatalf("Unable to close client:  %v", err)
		}
	}()

	err = api.SetupRoutes(r, wa)
	if err != nil {
		log.Fatalf("Unable to setup Routes: %v\n", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4444"
	}

	log.Printf("[BJW main]: running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Unable to Start Server on port %s - %v", port, err)
	}
}

func newWebApp() (*api.WebApp, error) {
	var wa api.WebApp
	var err error
	wa.DB, err = db.NewFSRepoClient()
	if err != nil {
		return nil, err
	}
	wa.Accounts = service.NewAccountsService(accounts, wa.DB)

	return &wa, err
}
