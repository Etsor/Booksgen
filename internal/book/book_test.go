package book

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"
)

func TestGenerateBook(t *testing.T) {
	b := generateBook()

	if len(b.ISBN) != 13 {
		t.Errorf("ISBN length: got %d, want 13", len(b.ISBN))
	}
	if b.Year < 1800 || b.Year > 2025 {
		t.Errorf("Year out of range: %d", b.Year)
	}
	if b.Pages < 1 || b.Pages > 2047 {
		t.Errorf("Pages out of range: %d", b.Pages)
	}
	if b.Title == "" {
		t.Error("Title is empty")
	}
	if b.Author == "" {
		t.Error("Author is empty")
	}
	if b.Genre == "" {
		t.Error("Genre is empty")
	}
	if b.Publisher == "" {
		t.Error("Publisher is empty")
	}
}

func TestGenerateBooks_Count(t *testing.T) {
	books := GenerateBooks(10)
	if len(books) != 10 {
		t.Errorf("Expected 10 books, got %d", len(books))
	}
}

func TestGenerateBooks_Zero(t *testing.T) {
	books := GenerateBooks(0)
	if len(books) != 0 {
		t.Errorf("Expected 0 books, got %d", len(books))
	}
}

func TestGenerateBooks_FieldsValid(t *testing.T) {
	books := GenerateBooks(50)
	for i, b := range books {
		if len(b.ISBN) != 13 {
			t.Errorf("books[%d]: ISBN length %d, want 13", i, len(b.ISBN))
		}
		if b.Year < 1800 || b.Year > 2025 {
			t.Errorf("books[%d]: Year %d out of range [1800, 2025]", i, b.Year)
		}
		if b.Pages < 1 || b.Pages > 2047 {
			t.Errorf("books[%d]: Pages %d out of range [1, 2047]", i, b.Pages)
		}
	}
}

func TestBooksToCSV_Header(t *testing.T) {
	books := []Book{}
	data := BooksToCSV(&books)

	r := csv.NewReader(strings.NewReader(string(data)))
	records, err := r.ReadAll()
	if err != nil {
		t.Fatalf("CSV parse error: %v", err)
	}

	if len(records) < 1 {
		t.Fatal("CSV has no rows")
	}

	want := []string{"isbn", "title", "author", "genre", "publisher", "year", "pages"}
	for i, h := range want {
		if records[0][i] != h {
			t.Errorf("Header[%d]: got %q, want %q", i, records[0][i], h)
		}
	}
}

func TestBooksToCSV_Row(t *testing.T) {
	books := []Book{
		{ISBN: "1234567890123", Title: "Test Title", Author: "Test Author",
			Genre: "Fiction", Publisher: "Test Pub", Year: 2000, Pages: 100},
	}

	data := BooksToCSV(&books)
	r := csv.NewReader(strings.NewReader(string(data)))
	records, err := r.ReadAll()
	if err != nil {
		t.Fatalf("CSV parse error: %v", err)
	}

	if len(records) != 2 {
		t.Fatalf("Expected 2 rows (header + 1), got %d", len(records))
	}

	row := records[1]
	if row[0] != "1234567890123" {
		t.Errorf("ISBN: got %q, want %q", row[0], "1234567890123")
	}
	if row[4] != "Test Pub" {
		t.Errorf("Publisher: got %q, want %q", row[4], "Test Pub")
	}
	if row[5] != "2000" {
		t.Errorf("Year: got %q, want %q", row[5], "2000")
	}
	if row[6] != "100" {
		t.Errorf("Pages: got %q, want %q", row[6], "100")
	}
}

func TestBooksToJSON(t *testing.T) {
	books := []Book{
		{ISBN: "1234567890123", Title: "Test", Author: "Author",
			Genre: "Fiction", Publisher: "Pub", Year: 2000, Pages: 300},
	}

	data := BooksToJSON(&books)

	var result struct {
		Books []Book `json:"books"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("JSON unmarshal error: %v", err)
	}

	if len(result.Books) != 1 {
		t.Fatalf("Expected 1 book, got %d", len(result.Books))
	}
	if result.Books[0].ISBN != "1234567890123" {
		t.Errorf("ISBN: got %q, want %q", result.Books[0].ISBN, "1234567890123")
	}
	if result.Books[0].Year != 2000 {
		t.Errorf("Year: got %d, want 2000", result.Books[0].Year)
	}
}

func TestBooksToJSON_Empty(t *testing.T) {
	books := []Book{}
	data := BooksToJSON(&books)

	var result struct {
		Books []Book `json:"books"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("JSON unmarshal error: %v", err)
	}
	if len(result.Books) != 0 {
		t.Errorf("Expected 0 books, got %d", len(result.Books))
	}
}

func TestBooksToXML(t *testing.T) {
	books := []Book{
		{ISBN: "1234567890123", Title: "Test", Author: "Author",
			Genre: "Fiction", Publisher: "Pub", Year: 2000, Pages: 300},
	}

	data := BooksToXML(&books)
	s := string(data)

	if !strings.HasPrefix(s, xml.Header) {
		t.Error("XML output missing XML declaration header")
	}

	xmlBody := strings.TrimPrefix(s, xml.Header)
	var result Books
	if err := xml.Unmarshal([]byte(xmlBody), &result); err != nil {
		t.Fatalf("XML unmarshal error: %v", err)
	}

	if len(result.Items) != 1 {
		t.Fatalf("Expected 1 book, got %d", len(result.Items))
	}
	if result.Items[0].ISBN != "1234567890123" {
		t.Errorf("ISBN: got %q, want %q", result.Items[0].ISBN, "1234567890123")
	}
	if result.Items[0].Pages != 300 {
		t.Errorf("Pages: got %d, want 300", result.Items[0].Pages)
	}
}

func TestBooksToXML_Empty(t *testing.T) {
	books := []Book{}
	data := BooksToXML(&books)
	s := string(data)

	if !strings.HasPrefix(s, xml.Header) {
		t.Error("XML output missing XML declaration header")
	}

	xmlBody := strings.TrimPrefix(s, xml.Header)
	var result Books
	if err := xml.Unmarshal([]byte(xmlBody), &result); err != nil {
		t.Fatalf("XML unmarshal error: %v", err)
	}
	if len(result.Items) != 0 {
		t.Errorf("Expected 0 books, got %d", len(result.Items))
	}
}
