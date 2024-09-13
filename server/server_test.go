// Test file
// Configuration.go
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Test configurationServer function
func TestConfigurationServer(t *testing.T) {

	// Dummy Handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("ZxFae testing"))
	})

	// Test with DH
	server := configurationServer(handler)

	type protectionConfig struct {
		name  string
		value interface{}
	}
	//Protections config
	expectedConfigs := []protectionConfig{
		{"ReadTimeout", 10 * time.Second},
		{"ReadHeaderTimeout", 5 * time.Second},
		{"WriteTimeout", 10 * time.Second},
		{"IdleTimeout", 120 * time.Second},
		{"MaxHeaderBytes", 1 << 20},
	}

	// Check each expectedConfig with loop
	for _, expected := range expectedConfigs {
		var currentValue interface{}
		switch expected.name {
		case "ReadTimeout":
			currentValue = server.ReadTimeout
		case "ReadHeaderTimeout":
			currentValue = server.ReadHeaderTimeout
		case "WriteTimeout":
			currentValue = server.WriteTimeout
		case "IdleTimeout":
			currentValue = server.IdleTimeout
		case "MaxHeaderBytes":
			currentValue = server.MaxHeaderBytes
		}

		if currentValue != expected.value {
			t.Errorf("Expected %s %v, got %v", expected.name, expected.value, currentValue)
		}
	}
}

// TestEnableCors tests middleware for various HTTP methods.
func TestEnableCors(t *testing.T) {
	tests := []struct {
		method string
		body   string
	}{
		//All methods expected
		{"GET", "Testing for GET!"},
		{"POST", "Testing for POST!"},
		{"PUT", "Testing for PUT!"},
		{"DELETE", "Testing for DELETE!"},
	}
	// Iterate over each test case.
	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			// Create a dummy handler that responds with the expected body.
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(test.body))
			})
			// Create a new HTTP request with the current method.
			dummyRequest, err := http.NewRequest(test.method, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Simulate a frontend request by setting the Origin header.
			dummyRequest.Header.Set("Origin", "http://localhost:3000")

			// Capture the response with NewRecorder.
			mockRequest := httptest.NewRecorder()

			// Use the EnableCors middleware with the dummy handler.
			enableCors(handler).ServeHTTP(mockRequest, dummyRequest)

			// Check if the response status is 200 OK.
			if status := mockRequest.Code; status != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
			}

			// Check specific CORS parameters.
			if origin := mockRequest.Header().Get("Access-Control-Allow-Origin"); origin != "http://localhost:3000" {
				t.Errorf("Expected Access-Control-Allow-Origin to be 'http://localhost:3000', got '%s'", origin)
			}
			if methods := mockRequest.Header().Get("Access-Control-Allow-Methods"); methods != "POST, GET, OPTIONS, PUT, DELETE" {
				t.Errorf("Expected Access-Control-Allow-Methods to be 'POST, GET, OPTIONS, PUT, DELETE', got '%s'", methods)
			}
			if headers := mockRequest.Header().Get("Access-Control-Allow-Headers"); headers != "Content-Type, Authorization" {
				t.Errorf("Expected Access-Control-Allow-Headers to be 'Content-Type, Authorization', got '%s'", headers)
			}
			if credentials := mockRequest.Header().Get("Access-Control-Allow-Credentials"); credentials != "true" {
				t.Errorf("Expected Access-Control-Allow-Credentials to be 'true', got '%s'", credentials)
			}
		})
	}
}
