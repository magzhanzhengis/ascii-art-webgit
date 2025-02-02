package main

import (
	"ascii-art-web/internal"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/ascii-art", asciiArtHandler)
	mux.HandleFunc("/notfound", notFoundHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Starting server at http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", mux))
}

// indexHandler serves the main page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Failed to load page", http.StatusInternalServerError)
	}
}

// asciiArtHandler processes the form and reloads the same page with ASCII output
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// ✅ For GET requests, just render the main page like the homepage
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, "Failed to load page", http.StatusInternalServerError)
		}
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	text := r.FormValue("inputText")
	bannerType := r.FormValue("banner")

	// Validate input
	if internal.ContainsNonASCII(text) {
		http.Error(w, "Input contains non-ASCII characters", http.StatusBadRequest)
		return
	}

	// Load the banner file
	bannerPath := filepath.Join("banners", bannerType+".txt")
	lines := internal.ReadBannerFile(bannerPath)
	asciiTemplates := internal.ParseBanner(lines)

	// Generate ASCII art
	output := internal.RenderASCIIArt(text, asciiTemplates)

	// Render the same index.html but with ASCII output
	data := struct {
		AsciiArt string
	}{
		AsciiArt: output,
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Failed to render result", http.StatusInternalServerError)
	}
}

// notFoundHandler - 404 not found page
func notFoundHandler(w http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		http.Error(w, "Failed to render result", http.StatusInternalServerError)
	}
	http.Error(w, "", http.StatusNotFound)
	// w.WriteHeader(http.StatusNotFound)
}
