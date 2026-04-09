package filewriter

import (
	b "Booksgen/internal/book"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDir(t *testing.T) {
	base := t.TempDir()
	nested := filepath.Join(base, "sub", "nested")

	CreateDir(nested)

	if _, err := os.Stat(nested); os.IsNotExist(err) {
		t.Errorf("Directory was not created: %s", nested)
	}
}

func TestCreateDir_Idempotent(t *testing.T) {
	dir := t.TempDir()

	CreateDir(dir)
	CreateDir(dir)
}

func TestCreateFile_First(t *testing.T) {
	dir := t.TempDir()

	f, err := CreateFile(dir, "books", ".csv")
	if err != nil {
		t.Fatalf("CreateFile error: %v", err)
	}
	f.Close()

	if _, err := os.Stat(filepath.Join(dir, "books.csv")); os.IsNotExist(err) {
		t.Error("Expected books.csv to exist")
	}
}

func TestCreateFile_Collision(t *testing.T) {
	dir := t.TempDir()

	f1, _ := CreateFile(dir, "books", ".csv")
	f1.Close()

	f2, err := CreateFile(dir, "books", ".csv")
	if err != nil {
		t.Fatalf("CreateFile second call error: %v", err)
	}
	f2.Close()

	if _, err := os.Stat(filepath.Join(dir, "books_1.csv")); os.IsNotExist(err) {
		t.Error("Expected books_1.csv to exist after collision")
	}
}

func TestCreateFile_MultipleCollisions(t *testing.T) {
	dir := t.TempDir()

	for range 3 {
		f, err := CreateFile(dir, "report", ".json")
		if err != nil {
			t.Fatalf("CreateFile error: %v", err)
		}
		f.Close()
	}

	for _, name := range []string{"report.json", "report_1.json", "report_2.json"} {
		if _, err := os.Stat(filepath.Join(dir, name)); os.IsNotExist(err) {
			t.Errorf("Expected %s to exist", name)
		}
	}
}

func sampleBooks() []b.Book {
	return []b.Book{
		{ISBN: "1234567890123", Title: "Test Book", Author: "Test Author",
			Genre: "Fiction", Publisher: "Test Pub", Year: 2000, Pages: 200},
	}
}

func TestWriteBooksToCSV(t *testing.T) {
	dir := t.TempDir()
	books := sampleBooks()

	WriteBooksToCSV(&books, dir)

	path := filepath.Join(dir, "books.csv")
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		t.Fatal("Expected books.csv to be created")
	}
	if info.Size() == 0 {
		t.Error("books.csv is empty")
	}
}

func TestWriteBooksToJSON(t *testing.T) {
	dir := t.TempDir()
	books := sampleBooks()

	WriteBooksToJSON(&books, dir)

	path := filepath.Join(dir, "books.json")
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		t.Fatal("Expected books.json to be created")
	}
	if info.Size() == 0 {
		t.Error("books.json is empty")
	}
}

func TestWriteBooksToXML(t *testing.T) {
	dir := t.TempDir()
	books := sampleBooks()

	WriteBooksToXML(&books, dir)

	path := filepath.Join(dir, "books.xml")
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		t.Fatal("Expected books.xml to be created")
	}
	if info.Size() == 0 {
		t.Error("books.xml is empty")
	}
}

func TestWriteBooks_CreatesDir(t *testing.T) {
	base := t.TempDir()
	dir := filepath.Join(base, "output")
	books := sampleBooks()

	WriteBooksToJSON(&books, dir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Error("Expected output directory to be created")
	}
}

func TestWriteBooks_CollisionSuffix(t *testing.T) {
	dir := t.TempDir()
	books := sampleBooks()

	WriteBooksToCSV(&books, dir)
	WriteBooksToCSV(&books, dir)

	if _, err := os.Stat(filepath.Join(dir, "books_1.csv")); os.IsNotExist(err) {
		t.Error("Expected books_1.csv on second write")
	}
}
