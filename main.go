package main

import (
	"embed"
	"log"
	"net/http"
	"prauth/entities"
	"prauth/routes"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

//go:embed templates/**/*.tmpl
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

func main() {
	var err error
	
	// 7 days expiry
	const SessionMaxAge = 7 * 24 * 60 * 60 // seconds
	
	appCtx := entities.AppCtx{
		CookieStore: *sessions.NewCookieStore([]byte("testcookiekey")),
	}

	appCtx.CookieStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   SessionMaxAge,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false, // set true in production with HTTPS
	}

	// load .env vars
	if err = godotenv.Load(); err != nil {
		log.Println("!!!!!!Error loading .env file")
	}

	routes.RunServer(&templateFS,staticFS, &appCtx)
}
