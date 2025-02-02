# ASCII Art Web

## 📄 Description

**ASCII Art Web** is a Go-based web application that converts user input text into stylized ASCII art using various banner styles. It provides a user-friendly interface where users can input text, select a banner style (Standard, Shadow, Thinkertoy), and view the generated ASCII art directly on the webpage.

---

## 👥 Authors

- **azhensio and mzhengis**

---

## 🚀 Usage: How to Run

### 1. **Clone the Repository**
```bash
 git clone https://01.tomorrow-school.ai/git/azhensio/ascii-art-web.git
 cd ascii-art-web
```

### 2. **Project Structure**
```
ascii-art-web/
├── banners/          # Banner styles (standard.txt, shadow.txt, thinkertoy.txt)
├── internal/         # Go internal logic (parsers, renderers)
│   └── parser.go
│   └── renderer.go
│   └── utils.go
├── static/           # Static files (CSS)
│   └── styles.css
├── templates/        # HTML templates
│   └── index.html
├── go.mod
├── main.go           # Entry point of the application
└── README.md
```

### 3. **Run the Application**
```bash
 go run main.go
```

### 4. **Access the Web App**
Open your browser and navigate to:
```
http://localhost:8080
```

---

## ⚙️ Implementation Details: Algorithm

### 1. **Text Input & Request Handling**
- **GET `/`**: Serves the main page using Go templates.
- **POST `/ascii-art`**: Receives user input (text + banner style) and processes it.

### 2. **ASCII Art Generation Workflow**
- **Input Validation**: Checks for non-ASCII characters and invalid inputs.
- **Banner Loading**: Reads the selected banner file (e.g., `standard.txt`).
- **Character Mapping**: Maps each character to its corresponding ASCII template.
- **Rendering**: Combines the ASCII templates to generate the final artwork.

### 3. **Error Handling**
- Returns proper HTTP status codes:
  - `200 OK` for successful operations
  - `400 Bad Request` for invalid input
  - `404 Not Found` for missing templates
  - `500 Internal Server Error` for server-side errors

---

## 🗒️ Notes
- Only standard Go packages are used.
- The app is optimized for ease of use and clean error handling.
- Supports dynamic error messages displayed on the frontend.

🚀

