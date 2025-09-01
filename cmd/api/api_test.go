package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBooksJSONHandler(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/books/json", jsonHandler(false))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/books/json")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}

	if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Errorf("Expected Content-Type application/json, got %s", ct)
	}
}

func TestBooksXMLHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/books/xml", xmlHandler(false))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/books/xml")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}

	if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "application/xml") {
		t.Errorf("Expected Content-Type application/xml, got %s", ct)
	}
}

func TestBooksCSVHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/books/csv", csvHandler(false))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/books/csv")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}

	if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "text/csv") {
		t.Errorf("Expected Content-Type text/csv, got %s", ct)
	}
}

func TestParseAmount(t *testing.T) {
	tests := []struct {
		query    string
		expected uint32
	}{
		{"?amount=10", 10},
		{"?amount=0", 1},
		{"?amount=-10", 1},
		{"", 1},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", "/books/json/"+tt.query, nil)

		got := parseAmount(req)
		if got != tt.expected {
			t.Errorf("parseAmount(%q) = %d\nExpected %d",
				tt.query, got, tt.expected)
		}
	}
}

func TestGetRealIP(t *testing.T) {
	// X-Forwarded-For
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "10.0.0.1")
	if ip := getRealIP(req); ip != "10.0.0.1" {
		t.Errorf("Expected 10.0.0.1, got %s", ip)
	}

	// X-Real-IP
	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Real-IP", "192.168.1.100")
	if ip := getRealIP(req); ip != "192.168.1.100" {
		t.Errorf("Expected 192.168.1.100, got %s", ip)
	}

	// fallback
	req = httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:12345"
	if ip := getRealIP(req); ip != "127.0.0.1" {
		t.Errorf("Expected 127.0.0.1, got %s", ip)
	}
}
