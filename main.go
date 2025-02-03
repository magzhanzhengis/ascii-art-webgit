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
	switch r.Method {
	case http.MethodGet:
		// Render the main page for GET requests
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, "Failed to load page", http.StatusInternalServerError)
		}
	case http.MethodPost:
		// Handle POST requests as usual
		processAsciiArt(w, r)
	default:
		// ✅ Handle unsupported HTTP methods
		renderError(w, http.StatusMethodNotAllowed, fmt.Sprintf("405 - Method '%s' Not Allowed", r.Method))
	}
}

func processAsciiArt(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		renderError(w, http.StatusBadRequest, "400 - Bad Request: Failed to parse form data")
		return
	}

	text := r.FormValue("inputText")
	bannerType := r.FormValue("banner")

	if internal.ContainsNonASCII(text) {
		renderError(w, http.StatusBadRequest, "400 - Input contains non-ASCII characters")
		return
	}

	bannerPath := filepath.Join("banners", bannerType+".txt")
	lines, err := internal.ReadBannerFile(bannerPath)
	if err != nil {
		renderError(w, http.StatusNotFound, fmt.Sprintf("404 - Banner '%s' Not Found", bannerType))
		return
	}

	asciiTemplates := internal.ParseBanner(lines)
	output := internal.RenderASCIIArt(text, asciiTemplates)

	data := struct {
		AsciiArt string
	}{
		AsciiArt: output,
	}

	err = templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "500 - Internal Server Error")
	}
}

func renderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	data := struct {
		Status  int
		Message string
	}{
		Status:  status,
		Message: message,
	}
	err := templates.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
	}
}

// notFoundHandler - 404 not found page
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderError(w, http.StatusNotFound, "404 - Page Not Found")
}
