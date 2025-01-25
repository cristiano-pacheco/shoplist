package tests

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	app := NewTestApp(t)

	// Start the app
	app.Start()

	// Wait for the server to start
	time.Sleep(2 * time.Second)

	// Get the port from the server logs
	port := "3000"

	// Make the request
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/livez", port))
	if err != nil {
		t.Fatalf("Failed to make HTTP request: %v", err)
	}

	// Assert the response
	assert.Equal(t, 200, resp.StatusCode)

	// Cleanup
	app.Stop()
}
