package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var config oauth2.Config
var ctx context.Context
var verifier *oidc.IDTokenVerifier

const state = "default"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx = context.Background()
	provider, err := oidc.NewProvider(ctx, os.Getenv("OIDC_ISSUER_URL"))
	if err != nil {
		log.Panic(err)
	}

	clientID := os.Getenv("OIDC_CLIENT_ID")
	config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: os.Getenv("OIDC_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OIDC_REDIRECT_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	verifier = provider.Verifier(&oidc.Config{
		ClientID:          clientID,
		SkipClientIDCheck: true,
	})

	initRoutes()

	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
