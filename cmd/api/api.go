package api

import (
	b "Booksgen/internal/book"
	s "Booksgen/pkg/style"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Run initializes and starts the HTTP server for the Booksgen API.
// It uses the flag package to configure the server's port, IP address,
// and optional SQLite database logging. The server exposes endpoints for
// generating book data in JSON, XML, and CSV formats. If logging is enabled,
// each request is logged to the SQLite database.
func Run() {
	fs := flag.NewFlagSet("booksgen", flag.ExitOnError)

	// Mode flag — registered so the parser doesn't error when invoked via main.
	_ = fs.Bool("api", false, "")

	var port int
	fs.IntVar(&port, "p", 8080, "server port")
	fs.IntVar(&port, "port", 8080, "server port")

	var ip string
	fs.StringVar(&ip, "i", "", "server IP address")
	fs.StringVar(&ip, "ip", "", "server IP address")

	var logdb bool
	fs.BoolVar(&logdb, "log", false, "enable SQLite request logging")

	var dbPath string
	fs.StringVar(&dbPath, "db", "./requests.db", "SQLite database path")
	fs.StringVar(&dbPath, "dbpath", "./requests.db", "SQLite database path")

	fs.Parse(os.Args[1:])

	if logdb {
		log.Printf("%s!Logging to sqlite db enabled!%s\n", s.FG_YELLOW, s.RESET)
		initDB(dbPath)
		fullDbPath, err := filepath.Abs(dbPath)
		if err != nil {
			log.Fatalf("%sError getting path to database: %s\n%s%s\n",
				s.FG_RED, fullDbPath, err, s.RESET)
		}
		log.Printf("Database path: %s", fullDbPath)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/books/json", jsonHandler(logdb))
	mux.HandleFunc("/books/xml", xmlHandler(logdb))
	mux.HandleFunc("/books/csv", csvHandler(logdb))

	log.Printf("%sServer is running on: %s:%d%s\n", s.FG_GREEN, ip, port, s.RESET)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), mux))
}

// jsonHandler returns an HTTP handler function that generates a JSON response containing
// a specified amount of books. The handler writes the generated books as JSON to the response.
// If logdb is true, the request is also logged to the database with the endpoint "/books/json".
func jsonHandler(logdb bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := parseAmount(r)
		w.Header().Set("Content-Type", "application/json")
		books := b.GenerateBooks(a)

		w.Write(b.BooksToJSON(&books))

		ip := getRealIP(r)
		log.Printf("JSON requested from: %s\nAmount: %d", ip, a)
		if logdb {
			logRequest(ip, "/books/json", int(a))
		}
	}
}

// xmlHandler returns an HTTP handler function that generates an XML response containing
// a specified amount of books. The handler writes the generated books as XML to the response.
// If logdb is true, the request is also logged to the database with the endpoint "/books/xml".
func xmlHandler(logdb bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := parseAmount(r)
		w.Header().Set("Content-Type", "application/xml")
		books := b.GenerateBooks(a)

		w.Write(b.BooksToXML(&books))

		ip := getRealIP(r)
		log.Printf("XML requested from: %s\nAmount: %d", ip, a)
		if logdb {
			logRequest(ip, "/books/xml", int(a))
		}
	}
}

// csvHandler returns an HTTP handler function that generates an CSV response containing
// a specified amount of books. The handler writes the generated books as CSV to the response.
// If logdb is true, the request is also logged to the database with the endpoint "/books/csv".
func csvHandler(logdb bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := parseAmount(r)
		w.Header().Set("Content-Type", "text/csv")
		books := b.GenerateBooks(a)

		w.Write(b.BooksToCSV(&books))

		ip := getRealIP(r)
		log.Printf("CSV requested from: %s\nAmount: %d", ip, a)
		if logdb {
			logRequest(ip, "/books/csv", int(a))
		}
	}
}

const maxAmount = 10_000

// parseAmount extracts the "amount" query parameter from the given HTTP request,
// attempts to convert it to a positive integer, and returns its value as uint32.
// If the parameter is missing, invalid, or less than or equal to zero, it defaults to 1.
// Values above maxAmount are capped to prevent excessive resource usage.
func parseAmount(r *http.Request) uint32 {
	q := r.URL.Query().Get("amount")
	amount, err := strconv.Atoi(q)

	if err != nil || amount <= 0 {
		return 1
	}

	if amount > maxAmount {
		return maxAmount
	}

	return uint32(amount)
}

// getIP extracts the IP address from the given HTTP request.
// It attempts to split the remote address into host and port using net.SplitHostPort.
// If splitting fails, it returns the original remote address.
// Otherwise, it returns the extracted IP address.
func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return ip
}

// getRealIP extracts the client's real IP address from the HTTP request.
// It first checks the "X-Forwarded-For" header (taking only the first entry,
// since proxies append their address to the list), then the "X-Real-IP" header,
// and finally falls back to using the remote address from the request.
func getRealIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// X-Forwarded-For may contain a comma-separated chain; the first is the client.
		if i := strings.Index(xff, ","); i != -1 {
			return strings.TrimSpace(xff[:i])
		}
		return xff
	}

	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	// fallback
	return getIP(r)
}
