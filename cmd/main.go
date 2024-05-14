package main

import (
	"fmt"
	"github.com/a-h/templ"
	pages "github.com/bastianwegge/go-templ-reload/pkg/feature-x"
	"github.com/bastianwegge/go-templ-reload/pkg/livereload"
	"net/http"
)

func main() {
	// --------------
	// Preparation
	// --------------
	http.HandleFunc("/reload/events", livereload.LiveReloadEventHandler)

	helloComponent := pages.Page()
	http.Handle("/", templ.Handler(helloComponent))

	// --------------
	// Simple listen
	// --------------
	fmt.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
