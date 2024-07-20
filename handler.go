package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	rawAccessToken := r.Header.Get("Authorization")
	if rawAccessToken == "" {
		log.Println("no token found")
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
		return
	}

	parts := strings.Split(rawAccessToken, " ")
	if len(parts) != 2 {
		w.WriteHeader(400)
		return
	}

	_, err := verifier.Verify(ctx, parts[1])
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
		return
	}

	w.Write([]byte("authenticated"))
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("state") != state {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
