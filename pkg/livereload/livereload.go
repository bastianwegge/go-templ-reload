package livereload

import (
	"fmt"
	"net/http"
	"time"
)

func LiveReloadEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Client connected")
	w.Header().Set("Access-Control-Allow-Origin", "*") // must have since it enable cors
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("Could not init http.Flusher")
	}

	for {
		select {
		case <-r.Context().Done():
			fmt.Println("Connection closed")
			return
		default:
			//fmt.Println("case message... sending message")
			fmt.Fprintf(w, "data: Ping\n\n")
			flusher.Flush()
			time.Sleep(5 * time.Second)
		}
	}
}
