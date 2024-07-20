package main

import "net/http"

func initRoutes() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/callback", handleCallback)
}
