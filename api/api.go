package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ResponseData struct for consistent API responses
type ResponseData struct {
	Status  string      `json:"status"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // Flexible data type for Data
}

// HealthyData struct for health check response
type HealthyData struct {
	Up bool `json:"up"`
}

// Reusable error handling function
func handleError(w http.ResponseWriter, err error, statusCode int) {
	if err != nil {
		w.WriteHeader(statusCode)
		responseData := ResponseData{
			Status:  "error",
			Error:   true,
			Message: err.Error(),
		}
		jsonData, err := json.Marshal(responseData)
		if err != nil {
			fmt.Println("Error marshalling error response:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}
}

// Function to send JSON responses
func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	jsonData, err := json.Marshal(data)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// Dashboard handler with logging
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := ResponseData{
		Status:  "success",
		Message: "Welcome to the dashboard!",
		Error:   false,
		Data:    "{}",
	}

	sendJSONResponse(w, http.StatusOK, data)
	fmt.Printf("Dashboard request successful: %s %s\n", r.Method, r.URL.Path) // Log successful request
}

// Health check handler
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthyData{Up: true}

	sendJSONResponse(w, http.StatusOK, response)
}

// Middleware to set Content-Type header
func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Run() {
	fmt.Println("=====================================")
	fmt.Println("=================api=================")
	fmt.Println("=====================================")

	mux := http.NewServeMux()
	mux.HandleFunc("/", dashboardHandler)
	mux.HandleFunc("/health", HealthCheckHandler)

	fmt.Println("❯ Listening on http://localhost:8081")
	fmt.Println("❯ Access http://localhost:8081/health to check the health of the server")
	http.ListenAndServe(":8081", JsonMiddleware(mux))
}
