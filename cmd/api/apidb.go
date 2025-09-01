package api

import (
    s "Booksgen/pkg/style"
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB(dbPath string) {
    var err error
    db, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Fatalf("%sError opening DB: %s%s",
            s.FG_RED, err, s.RESET)
    }

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS requests (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            ip TEXT,
            endpoint TEXT,
            amount INTEGER,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)

    if err != nil {
        log.Fatalf("%sError creating table: %s%s",
            s.FG_RED, err, s.RESET)
    }
}

func logRequest(ip, endpoint string, amount int) {
    _, err := db.Exec(
        "INSERT INTO requests (ip, endpoint, amount) VALUES (?, ?, ?)",
        ip, endpoint, amount,
    )

    if err != nil {
        log.Printf("%sError inserting request: %s%s",
            s.FG_RED, err, s.RESET)
    }
}
