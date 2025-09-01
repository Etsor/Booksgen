package api

import (
	b "Booksgen/internal/book"
	u "Booksgen/internal/utils"
	s "Booksgen/pkg/style"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Run initializes and starts the HTTP server for the Booksgen API.
// It parses command-line arguments to configure the server's port, IP address,
// and optional SQLite database logging. The server exposes endpoints for
// generating book data in JSON, XML, and CSV formats. If logging is enabled,
// each request is logged to the SQLite database.
func Run() {
	port := 8080
	ip := ""

	logdb := u.HasArg("--log")
	if logdb {
		logDB()
	}

	if u.HasArg("-p") || u.HasArg("--port") {
		for i, arg := range os.Args {
			if arg == "-p" || arg == "--port" {
				pPos := i
				var err error
				port, err = strconv.Atoi(os.Args[pPos+1])
				if err != nil {
					log.Fatalf("%sInvalid port%s\n",
						s.FG_RED, s.RESET)
				}
			}
		}
	}

	if u.HasArg("-i") || u.HasArg("--ip") {
		for i, arg := range os.Args {
			if arg == "-i" || arg == "--ip" {
				ipPos := i
				ip = os.Args[ipPos+1]
			}
		}
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/books/json", jsonHandler(logdb))
	mux.HandleFunc("/books/xml", xmlHandler(logdb))
	mux.HandleFunc("/books/csv", csvHandler(logdb))

	log.Printf("%sServer is running on: %s:%d%s\n",
		s.FG_GREEN, ip, port, s.RESET)

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

// logDB initializes logging to a SQLite database. It determines the database path from
// command-line arguments ("-db" or "--dbpath"), falling back to a default path if not provided.
func logDB() {
	dbPath := "./requests.db"
	if u.HasArg("-db") || u.HasArg("--dbpath") {
		for i, arg := range os.Args {
			if arg == "-db" || arg == "--dbpath" {
				dbPos := i
				dbPath = os.Args[dbPos+1]
			}
		}
	}
	log.Printf("%s!Logging to sqlite db enabled!%s\n",
		s.FG_YELLOW, s.RESET)

	initDB(dbPath)

	fullDbPath, err := filepath.Abs(dbPath)
	if err != nil {
		log.Fatalf("%sError getting path to database: %s\n%s%s\n",
			s.FG_RED, fullDbPath, err, s.RESET)
	}
	log.Printf("Database path: %s", fullDbPath)

	defer db.Close()
}

// parseAmount extracts the "amount" query parameter from the given HTTP request,
// attempts to convert it to a positive integer, and returns its value as uint32.
// If the parameter is missing, invalid, or less than or equal to zero, it defaults to 1.
func parseAmount(r *http.Request) uint32 {
	q := r.URL.Query().Get("amount")
	amount, err := strconv.Atoi(q)

	if err != nil || amount <= 0 {
		return 1
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
// It first checks the "X-Forwarded-For" header, then the "X-Real-IP" header,
// and finally falls back to using the remote address from the request.
// This is useful when the server is behind a proxy or load balancer.
func getRealIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip
	}

	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}
	// fallback
	return getIP(r)
}
