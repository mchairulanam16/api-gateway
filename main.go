package main

import (
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define route to forward requests to backend service
	router.Any("/api/*path", func(c *gin.Context) {
		// Create a reverse proxy for the backend service
		proxy := getReverseProxy("http://localhost:8080")

		// Modify the request path to strip the /api prefix
		c.Request.URL.Path = stripPrefix(c.Request.URL.Path, "/api")

		// Forward the modified request to the backend service
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// Start the server on port 8081
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}

// getReverseProxy creates a new reverse proxy for the given target URL
func getReverseProxy(target string) *httputil.ReverseProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		panic(err) // Handle errors appropriately in production
	}
	return httputil.NewSingleHostReverseProxy(targetURL)
}

// stripPrefix removes the specified prefix from the given path
func stripPrefix(path, prefix string) string {
	if strings.HasPrefix(path, prefix) {
		return path[len(prefix):]
	}
	return path
}
