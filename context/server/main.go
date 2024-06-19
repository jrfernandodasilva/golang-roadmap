package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	println("❯ context.server start")

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)

	println("❯ context.server end")
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("🛈 Request start")
	defer log.Println("🛈 Request end")

	select {
	case <-time.After(time.Second * 5):
		log.Println("✓ request context handled with success")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("✓ request context handled with success"))
	case <-ctx.Done():
		log.Println("🗙 request context canceled")
		http.Error(w, "request context canceled", http.StatusRequestTimeout)
	}
}
