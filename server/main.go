package main

import (
	"io"
	"log"
	"net/http"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func convertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	html, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	markdown, err := htmltomarkdown.ConvertString(string(html))
	if err != nil {
		log.Printf("Conversion error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(markdown))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	log.Println("HTML→MD sidecar starting on :7777")
	http.HandleFunc("/convert", convertHandler)
	http.HandleFunc("/health", healthHandler)
	log.Println("Endpoints: /convert (POST), /health (GET)")
	log.Fatal(http.ListenAndServe(":7777", nil))
}
