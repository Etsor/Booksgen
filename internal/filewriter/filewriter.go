package filewriter

import (
    b "Booksgen/internal/book"
    s "Booksgen/pkg/style"
    "bufio"
    "log"
    "os"
    "path/filepath"
    "strconv"
)

// WriteBooksToCSV exports a slice of Book objects to an CSV file in the specified directory.
// It creates the directory if it does not exist, generates the CSV file, and writes the serialized
// CSV data to the file. If an error occurs during file creation, it logs a fatal error and exits.
func WriteBooksToCSV(books *[]b.Book, dirPath string) {
    CreateDir(dirPath)
    f, err := CreateFile(dirPath, "books", ".csv")
    if err != nil {
        log.Fatalf("%sError occured while exporting to CSV: %s%s\n",
            s.FG_RED, err, s.RESET)
    }

    defer f.Close()

    w := bufio.NewWriter(f)
    w.Write(b.BooksToCSV(books))

    log.Printf("Written to CSV\n")
    w.Flush()
}

// WriteBooksToJSON exports a slice of Book objects to an JSON file in the specified directory.
// It creates the directory if it does not exist, generates the JSON file, and writes the serialized
// JSON data to the file. If an error occurs during file creation, it logs a fatal error and exits.
func WriteBooksToJSON(books *[]b.Book, dirPath string) {
    CreateDir(dirPath)
    f, err := CreateFile(dirPath, "books", ".json")
    if err != nil {
        log.Fatalf("%sError occured while exporting to JSON: %s%s\n",
            s.FG_RED, err, s.RESET)
    }

    defer f.Close()

    w := bufio.NewWriter(f)
    w.Write(b.BooksToJSON(books))

    log.Printf("Written to JSON\n")
    w.Flush()
}

// WriteBooksToXML exports a slice of Book objects to an XML file in the specified directory.
// It creates the directory if it does not exist, generates the XML file, and writes the serialized
// XML data to the file. If an error occurs during file creation, it logs a fatal error and exits.
func WriteBooksToXML(books *[]b.Book, dirPath string) {
    CreateDir(dirPath)
    f, err := CreateFile(dirPath, "books", ".xml")
    if err != nil {
        log.Fatalf("%sError occured while exporting to XML: %s%s\n",
            s.FG_RED, err, s.RESET)
    }

    defer f.Close()

    w := bufio.NewWriter(f)
    w.Write(b.BooksToXML(books))

    log.Printf("Written to XML\n")
    w.Flush()
}

// createDir creates a directory at the specified path, including any necessary parents.
// If an error occurs during directory creation, the function logs a fatal error and terminates the program.
func CreateDir(dirPath string) {
    err := os.MkdirAll(dirPath, os.ModePerm)
    if err != nil {
        log.Fatalf("%sError creating dir: %s%s\n",
            s.FG_RED, err, s.RESET)
    }
}

// createFile attempts to create a new file in the specified directory with the given base name and extension.
// If a file with the same name already exists, it appends an incrementing number to the base name (e.g., "file_1.txt")
// until an unused filename is found. Returns the created file and any error encountered during the process.
func CreateFile(dirPath, baseName, ext string) (*os.File, error) {
    i := 0
    fullPath := ""

    for {
        if i == 0 {
            fullPath = filepath.Join(dirPath, baseName+ext)
        } else {
            fullPath = filepath.Join(dirPath, baseName+"_"+strconv.Itoa(i)+ext)
        }

        _, err := os.Stat(fullPath)
        if os.IsNotExist(err) {
            return os.Create(fullPath)
        }
        if err != nil {
            return nil, err
        }

        i++
    }
}
