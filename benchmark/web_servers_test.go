package benchmark

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

// Handler using net/http
func netHTTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

// Handler using gin
func ginHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

// Handler using echo
func echoHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func BenchmarkNetHTTP(b *testing.B) {
	handler := http.HandlerFunc(netHTTPHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	for i := 0; i < b.N; i++ {
		http.Get(server.URL)
	}
}

func BenchmarkGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", ginHandler)
	server := httptest.NewServer(r)
	defer server.Close()

	for i := 0; i < b.N; i++ {
		http.Get(server.URL)
	}
}

func BenchmarkEcho(b *testing.B) {
	e := echo.New()
	e.GET("/", echoHandler)
	server := httptest.NewServer(e)
	defer server.Close()

	for i := 0; i < b.N; i++ {
		http.Get(server.URL)
	}
}
