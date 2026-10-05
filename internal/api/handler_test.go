package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	router := NewRouter()

	tests := []struct {
		name   string
		method string
		path   string
		status int
	}{
		{"health", http.MethodGet, "/health", http.StatusOK},
		{"list items", http.MethodGet, "/api/v1/items", http.StatusOK},
		{"get item", http.MethodGet, "/api/v1/items/1", http.StatusOK},
		{"item not found", http.MethodGet, "/api/v1/items/999", http.StatusNotFound},
		{"invalid id", http.MethodGet, "/api/v1/items/abc", http.StatusBadRequest},
		{"unknown route", http.MethodGet, "/nope", http.StatusNotFound},
		{"wrong method", http.MethodPost, "/api/v1/items", http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path, nil))
			if rec.Code != tt.status {
				t.Fatalf("status = %d, want %d (body: %s)", rec.Code, tt.status, rec.Body)
			}
		})
	}
}

func TestListItemsBody(t *testing.T) {
	rec := httptest.NewRecorder()
	NewRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/v1/items", nil))

	if ct := rec.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Fatalf("Content-Type = %q", ct)
	}
	var body struct {
		Host  HostInfo `json:"host"`
		Data  []Item   `json:"data"`
		Total int      `json:"total"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if body.Total != len(items) || len(body.Data) != len(items) {
		t.Fatalf("got total=%d len=%d, want %d", body.Total, len(body.Data), len(items))
	}
	if body.Host.Hostname == "" || body.Host.IPAddresses == nil {
		t.Fatalf("host info missing: %+v", body.Host)
	}
}
