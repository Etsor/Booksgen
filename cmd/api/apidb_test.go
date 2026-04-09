package api

import (
	"path/filepath"
	"testing"
)

func TestInitDB_CreatesTable(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")

	initDB(dbPath)
	t.Cleanup(func() { db.Close() })

	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='requests'")
	if err != nil {
		t.Fatalf("Query error: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Error("Expected 'requests' table to exist after initDB")
	}
}

func TestInitDB_Idempotent(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")

	// CREATE TABLE IF NOT EXISTS — calling twice must not error.
	initDB(dbPath)
	t.Cleanup(func() { db.Close() })
	initDB(dbPath)
}

func TestLogRequest_InsertsRow(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")

	initDB(dbPath)
	t.Cleanup(func() { db.Close() })

	logRequest("127.0.0.1", "/books/json", 5)

	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM requests").Scan(&count); err != nil {
		t.Fatalf("QueryRow error: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 row, got %d", count)
	}
}

func TestLogRequest_Fields(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")

	initDB(dbPath)
	t.Cleanup(func() { db.Close() })

	logRequest("10.0.0.1", "/books/xml", 42)

	var ip, endpoint string
	var amount int
	err := db.QueryRow("SELECT ip, endpoint, amount FROM requests").Scan(&ip, &endpoint, &amount)
	if err != nil {
		t.Fatalf("QueryRow error: %v", err)
	}

	if ip != "10.0.0.1" {
		t.Errorf("ip: got %q, want %q", ip, "10.0.0.1")
	}
	if endpoint != "/books/xml" {
		t.Errorf("endpoint: got %q, want %q", endpoint, "/books/xml")
	}
	if amount != 42 {
		t.Errorf("amount: got %d, want 42", amount)
	}
}

func TestLogRequest_MultipleRows(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")

	initDB(dbPath)
	t.Cleanup(func() { db.Close() })

	logRequest("1.1.1.1", "/books/json", 1)
	logRequest("2.2.2.2", "/books/csv", 10)
	logRequest("3.3.3.3", "/books/xml", 100)

	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM requests").Scan(&count); err != nil {
		t.Fatalf("QueryRow error: %v", err)
	}
	if count != 3 {
		t.Errorf("Expected 3 rows, got %d", count)
	}
}
