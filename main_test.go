package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEndpoints(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody string
	}{
		// Wallet endpoint tests
		{
			name:         "Wallet default (no query param)",
			path:         "/api/v1/wallet",
			expectedCode: http.StatusOK,
			expectedBody: `"id": "wal_29a1f3c8"`,
		},
		{
			name:         "Wallet success scenario",
			path:         "/api/v1/wallet?scenario=success",
			expectedCode: http.StatusOK,
			expectedBody: `"balance": 1250.75`,
		},
		{
			name:         "Wallet empty scenario",
			path:         "/api/v1/wallet?scenario=empty",
			expectedCode: http.StatusOK,
			expectedBody: `"balance": 0.00`,
		},
		{
			name:         "Wallet error scenario",
			path:         "/api/v1/wallet?scenario=error",
			expectedCode: http.StatusInternalServerError,
			expectedBody: `"code": "INTERNAL_SERVER_ERROR"`,
		},

		// Children endpoint tests
		{
			name:         "Children default scenario",
			path:         "/api/v1/children",
			expectedCode: http.StatusOK,
			expectedBody: `"id": "child_001"`,
		},
		{
			name:         "Children success scenario",
			path:         "/api/v1/children?scenario=success",
			expectedCode: http.StatusOK,
			expectedBody: `"first_name": "Elif"`,
		},
		{
			name:         "Children empty scenario",
			path:         "/api/v1/children?scenario=empty",
			expectedCode: http.StatusOK,
			expectedBody: `[]`,
		},
		{
			name:         "Children error scenario",
			path:         "/api/v1/children?scenario=error",
			expectedCode: http.StatusInternalServerError,
			expectedBody: `"message": "The requested information could not be loaded."`,
		},

		// Transactions endpoint tests
		{
			name:         "Transactions default scenario",
			path:         "/api/v1/transactions",
			expectedCode: http.StatusOK,
			expectedBody: `"id": "txn_a001"`,
		},
		{
			name:         "Transactions success scenario",
			path:         "/api/v1/transactions?scenario=success",
			expectedCode: http.StatusOK,
			expectedBody: `"category": "cafeteria"`,
		},
		{
			name:         "Transactions empty scenario",
			path:         "/api/v1/transactions?scenario=empty",
			expectedCode: http.StatusOK,
			expectedBody: `[]`,
		},
		{
			name:         "Transactions error scenario",
			path:         "/api/v1/transactions?scenario=error",
			expectedCode: http.StatusInternalServerError,
			expectedBody: `"code": "INTERNAL_SERVER_ERROR"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			rec := httptest.NewRecorder()

			var handler http.HandlerFunc
			switch {
			case strings.HasPrefix(tc.path, "/api/v1/wallet"):
				handler = func(w http.ResponseWriter, r *http.Request) {
					serveScenario(w, r, walletSuccessJSON, walletEmptyJSON)
				}
			case strings.HasPrefix(tc.path, "/api/v1/children"):
				handler = func(w http.ResponseWriter, r *http.Request) {
					serveScenario(w, r, childrenSuccessJSON, emptyArrayJSON)
				}
			case strings.HasPrefix(tc.path, "/api/v1/transactions"):
				handler = func(w http.ResponseWriter, r *http.Request) {
					serveScenario(w, r, transactionsSuccessJSON, emptyArrayJSON)
				}
			}

			handler.ServeHTTP(rec, req)

			if rec.Code != tc.expectedCode {
				t.Errorf("Expected status code %d, got %d", tc.expectedCode, rec.Code)
			}

			contentType := rec.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("Expected Content-Type application/json, got %s", contentType)
			}

			body := rec.Body.String()
			if !strings.Contains(body, tc.expectedBody) {
				t.Errorf("Expected body to contain %q, got %s", tc.expectedBody, body)
			}
		})
	}
}
