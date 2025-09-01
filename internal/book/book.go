package book

import (
    "Booksgen/pkg/style"
    "encoding/json"
    "encoding/xml"
    "fmt"
    "log"
    "math/rand"
    "runtime"
    "strings"
    "sync"
)

type Book struct {
    XMLName   xml.Name `xml:"book" json:"-"`
    ISBN      string   `xml:"isbn" json:"isbn"`
    Title     string   `xml:"title" json:"title"`
    Author    string   `xml:"author" json:"author"`
    Genre     string   `xml:"genre" json:"genre"`
    Publisher string   `xml:"publisher" json:"publisher"`
    Year      uint16   `xml:"year" json:"year"`
    Pages     uint16   `xml:"pages" json:"pages"`
}

type Books struct {
    XMLName xml.Name `xml:"books" json:"-"`
    Items   []Book   `xml:"book" json:"books"`
}

func generateBook() Book {
    var b Book

    for range 13 {
        b.ISBN += fmt.Sprint(rand.Intn(10))
    }

    b.Title = Words[rand.Intn(len(Words))] + " " +
        Words[rand.Intn(len(Words))]

    b.Author = Firnames[rand.Intn(len(Firnames))] + " " +
        Surnames[rand.Intn(len(Surnames))]

    b.Genre = Genres[rand.Intn(len(Genres))]
    b.Publisher = Publishers[rand.Intn(len(Publishers))]
    b.Year = uint16(rand.Int31n(2025))
    b.Pages = uint16(rand.Int31n(2048))

    return b
}

// GenerateBooks creates and returns a slice of Book structs of the specified amount.
// It utilizes concurrent workers, equal to the number of CPU cores, to generate books in parallel.
func GenerateBooks(amount uint32) []Book {
    var wg sync.WaitGroup

    numWorkers := runtime.NumCPU()
    jobs := make(chan uint32, amount)

    books := make([]Book, amount)

    log.Printf("Generating books using %d CPUs..", numWorkers)
    for range numWorkers {
        wg.Go(func() {
            for i := range jobs {
                books[i] = generateBook()
            }
        })
    }

    for i := range amount {
        jobs <- i
    }

    close(jobs)

    wg.Wait()

    log.Printf("%s%sBooks generated%s",
        style.BG_GREEN, style.FG_BLACK, style.RESET)

    return books
}

// BooksToCSV serializes a slice of Book structs into an indented CSV format
// and returns the result as a byte slice.
// If serialization fails, it logs a fatal error and exits.
func BooksToCSV(books *[]Book) []byte {
    var sb strings.Builder
    sb.WriteString("isbn,title,author,genre,publisher,year,pages\n")

    for _, b := range *books {
        fmt.Fprintf(&sb, "%s,%s,%s,%s,%s,%d,%d\n",
            b.ISBN, b.Title, b.Author, b.Genre, b.Publisher, b.Year, b.Pages)
    }

    return []byte(sb.String())
}

// BooksToJSON serializes a slice of Book structs into an indented JSON format
// and returns the result as a byte slice.
// If serialization fails, it logs a fatal error and exits.
func BooksToJSON(books *[]Book) []byte {
    wrapper := Books{Items: *books}
    data, err := json.MarshalIndent(wrapper, "", "  ")
    if err != nil {
        log.Fatalf("%sError occured while JSON serialization: %s\n",
            style.FG_RED, style.RESET)
        panic(err)
    }

    return data
}

// BooksToXML serializes a slice of Book structs into an indented XML format
// and returns the result as a byte slice.
// If serialization fails, it logs a fatal error and exits.
func BooksToXML(books *[]Book) []byte {
    wrapper := Books{Items: *books}
    data, err := xml.MarshalIndent(wrapper, "", "    ")
    if err != nil {
        log.Fatalf("%sError occured while XML serialization: %s%s",
            style.FG_RED, err, style.RESET)
    }

    return []byte(xml.Header + string(data))
}
