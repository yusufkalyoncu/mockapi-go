package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Success responses
var walletSuccessJSON = []byte(`{
  "id": "wal_29a1f3c8",
  "currency": "TRY",
  "balance": 1250.75,
  "created_at": "2025-09-15T10:30:00Z"
}`)

var childrenSuccessJSON = []byte(`[
  {
    "id": "child_001",
    "first_name": "Elif",
    "last_name": "Yılmaz",
    "avatar_url": null,
    "age": 12,
    "grade": "7th Grade",
    "wallet_balance": 340.00,
    "currency": "TRY",
    "school_name": "Bahçeşehir Koleji"
  },
  {
    "id": "child_002",
    "first_name": "Can",
    "last_name": "Yılmaz",
    "avatar_url": null,
    "age": 9,
    "grade": "4th Grade",
    "wallet_balance": 125.50,
    "currency": "TRY",
    "school_name": "Bahçeşehir Koleji"
  },
  {
    "id": "child_003",
    "first_name": "Zeynep",
    "last_name": "Yılmaz",
    "avatar_url": null,
    "age": 15,
    "grade": "10th Grade",
    "wallet_balance": 580.00,
    "currency": "TRY",
    "school_name": "Doğa Koleji"
  }
]`)

var transactionsSuccessJSON = []byte(`[
  {
    "id": "txn_a001",
    "type": "expense",
    "category": "cafeteria",
    "description": "School Cafeteria — Elif",
    "amount": -45.00,
    "currency": "TRY",
    "date": "2026-03-28T12:35:00Z",
    "child_id": "child_001",
    "status": "completed"
  },
  {
    "id": "txn_a002",
    "type": "income",
    "category": "topup",
    "description": "Wallet Top-Up",
    "amount": 500.00,
    "currency": "TRY",
    "date": "2026-03-27T09:15:00Z",
    "child_id": null,
    "status": "completed"
  },
  {
    "id": "txn_a003",
    "type": "expense",
    "category": "trip",
    "description": "Cappadocia School Trip — Can",
    "amount": -1200.00,
    "currency": "TRY",
    "date": "2026-03-25T16:00:00Z",
    "child_id": "child_002",
    "status": "completed"
  },
  {
    "id": "txn_a004",
    "type": "expense",
    "category": "cafeteria",
    "description": "School Cafeteria — Zeynep",
    "amount": -32.50,
    "currency": "TRY",
    "date": "2026-03-25T12:20:00Z",
    "child_id": "child_003",
    "status": "completed"
  },
  {
    "id": "txn_a005",
    "type": "expense",
    "category": "campus_store",
    "description": "Campus Bookstore — Elif",
    "amount": -89.90,
    "currency": "TRY",
    "date": "2026-03-24T14:10:00Z",
    "child_id": "child_001",
    "status": "completed"
  },
  {
    "id": "txn_a006",
    "type": "income",
    "category": "topup",
    "description": "Wallet Top-Up",
    "amount": 1000.00,
    "currency": "TRY",
    "date": "2026-03-22T08:45:00Z",
    "child_id": null,
    "status": "completed"
  },
  {
    "id": "txn_a007",
    "type": "expense",
    "category": "event",
    "description": "Science Fair Registration — Can",
    "amount": -150.00,
    "currency": "TRY",
    "date": "2026-03-20T10:30:00Z",
    "child_id": "child_002",
    "status": "completed"
  },
  {
    "id": "txn_a008",
    "type": "expense",
    "category": "cafeteria",
    "description": "School Cafeteria — Elif",
    "amount": -38.00,
    "currency": "TRY",
    "date": "2026-03-19T12:45:00Z",
    "child_id": "child_001",
    "status": "completed"
  },
  {
    "id": "txn_a009",
    "type": "income",
    "category": "cashback",
    "description": "Cashback Reward — March",
    "amount": 25.00,
    "currency": "TRY",
    "date": "2026-03-18T00:00:00Z",
    "child_id": null,
    "status": "completed"
  },
  {
    "id": "txn_a010",
    "type": "expense",
    "category": "trip",
    "description": "Visa Fee — Italy Trip — Zeynep",
    "amount": -320.00,
    "currency": "TRY",
    "date": "2026-03-15T11:00:00Z",
    "child_id": "child_003",
    "status": "completed"
  }
]`)

// Empty responses
var walletEmptyJSON = []byte(`{
  "id": "wal_new_user_01",
  "currency": "TRY",
  "balance": 0.00,
  "created_at": "2026-03-30T08:00:00Z"
}`)

var emptyArrayJSON = []byte(`[]`)

// Error response (500 Internal Server Error)
var errorResponseJSON = []byte(`{
  "code": "INTERNAL_SERVER_ERROR",
  "message": "The requested information could not be loaded."
}`)

func serveScenario(w http.ResponseWriter, r *http.Request, successJSON, emptyJSON []byte) {
	// Add CORS headers for browser/webview/mobile testing convenience
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	scenario := r.URL.Query().Get("scenario")
	switch scenario {
	case "error":
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorResponseJSON)
	case "empty":
		w.WriteHeader(http.StatusOK)
		w.Write(emptyJSON)
	default:
		// Default to success response if scenario is empty, "success", or any other value
		w.WriteHeader(http.StatusOK)
		w.Write(successJSON)
	}
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		scenario := r.URL.Query().Get("scenario")
		if scenario == "" {
			scenario = "default(success)"
		}
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s ?scenario=%s — %v", r.Method, r.URL.Path, scenario, time.Since(start))
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("/api/v1/wallet", logMiddleware(func(w http.ResponseWriter, r *http.Request) {
		serveScenario(w, r, walletSuccessJSON, walletEmptyJSON)
	}))

	mux.HandleFunc("/api/v1/children", logMiddleware(func(w http.ResponseWriter, r *http.Request) {
		serveScenario(w, r, childrenSuccessJSON, emptyArrayJSON)
	}))

	mux.HandleFunc("/api/v1/transactions", logMiddleware(func(w http.ResponseWriter, r *http.Request) {
		serveScenario(w, r, transactionsSuccessJSON, emptyArrayJSON)
	}))

	// Health check endpoint for container monitoring
	mux.HandleFunc("/health", logMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}))

	log.Printf("🚀 Mock API Server is running on port %s", port)
	log.Printf("Available GET endpoints:")
	log.Printf("  • http://localhost:%s/api/v1/wallet", port)
	log.Printf("  • http://localhost:%s/api/v1/children", port)
	log.Printf("  • http://localhost:%s/api/v1/transactions", port)
	log.Printf("Supported query parameters: ?scenario=success | ?scenario=empty | ?scenario=error")

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
