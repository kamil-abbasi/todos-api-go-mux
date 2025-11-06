package main

import (
	"net/http"
	"todos-api/internal"
)

func main() {
	http.ListenAndServe(":3000", internal.NewApiRouter())
}
