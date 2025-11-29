package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	app, err := Initialize()

	// If database is not available, skip the test
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	// Test initialization
	assert.NotNil(t, app, "Initialize should return app")
	assert.NotNil(t, app.Fiber, "Fiber app should be initialized")
	assert.NotNil(t, app.DB, "Database should be initialized")
	assert.NotNil(t, app.Logger, "Logger should be initialized")
	assert.NotNil(t, app.Config, "Config should be initialized")
}

func TestHealthEndpoint(t *testing.T) {
	app, err := Initialize()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	// Create test request
	req := httptest.NewRequest("GET", "/health", nil)

	// Test the request
	resp, err := app.Fiber.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), "ok")
	assert.Contains(t, string(body), "WMS Backend is running")
}

func TestNotFoundRoute(t *testing.T) {
	app, err := Initialize()
	if err != nil {
		t.Skipf("Skipping test - database not available: %v", err)
		return
	}

	// Test non-existent route
	req := httptest.NewRequest("GET", "/non-existent-route", nil)
	resp, err := app.Fiber.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), "Route not found")
}
