package standalone

import (
    "Booksgen/cmd/standalone/cover"
    b "Booksgen/internal/book"
    fw "Booksgen/internal/filewriter"
    u "Booksgen/internal/utils"
    s "Booksgen/pkg/style"
    "log"
    "os"
    "path/filepath"
    "strconv"
)

// Run starts standalone app, parses command-line arguments to determine output format (XML, JSON, CSV),
// output directory, and the number of books to generate. It generates the specified
// number of books and writes them to the chosen output format(s) in the specified directory.
func Run() {
    x := u.HasArg("-x") || u.HasArg("--xml")
    j := u.HasArg("-j") || u.HasArg("--json")
    c := u.HasArg("-c") || u.HasArg("--csv")
    o := u.HasArg("-o") || u.HasArg("--output")
    a := u.HasArg("-a") || u.HasArg("--amount")
    cov := u.HasArg("-cov") || u.HasArg("--cover")

    amount := 1
    dirPath := "./output/"

    if a {
        for i, arg := range os.Args {
            if arg == "-a" || arg == "--amount" {
                aPos := i
                var err error
                amount, err = strconv.Atoi(os.Args[aPos+1])
                if err != nil {
                    log.Fatalf("%sInvalid amount argument%s\n",
                        s.FG_RED, s.RESET)
                }
            }
        }
    }

    if o {
        for i, arg := range os.Args {
            if arg == "-o" || arg == "--output" {
                oPos := i
                dirPath = os.Args[oPos+1]
                oDir, err := filepath.Abs(dirPath)
                if err != nil {
                    log.Fatalf("%sError getting path to directory: %s\n%s%s\n",
                        s.FG_RED, oDir, err, s.RESET)
                }

                log.Printf("Path to output directory: %s", oDir)
            }
        }
    }

    books := b.GenerateBooks(uint32(amount))
    
    if cov {
        w := 400
        h := 650

        amount := 1
        dirPath := "./covoutput/"

        if u.HasArg("-cova") || u.HasArg("--covamount") {
            for i, arg := range os.Args {
                if arg == "-cova" || arg == "--covamount" {
                    var err error
                    amount, err = strconv.Atoi(os.Args[i+1])
                    if err != nil {
                        log.Fatalf("%sInvalid cover amount argument%s\n",
                            s.FG_RED, s.RESET)
                    }
                }
            }
        }

        if u.HasArg("-covo") || u.HasArg("--covoutput") {
            for i, arg := range os.Args {
                if arg == "-covo" || arg == "--covoutput" {
                    dirPath = os.Args[i+1]
                }
            }
        }

        if u.HasArg("-covw") || u.HasArg("--covwidth") {
            for i, arg := range os.Args {
                if arg == "-covw" || arg == "--covwidth" {
                    var err error
                    w, err = strconv.Atoi(os.Args[i+1])
                    if err != nil {
                        log.Fatalf("%sInvalid width argument%s\n",
                            s.BG_RED, s.RESET)
                    }
                }
            }
        }

        if u.HasArg("-covh") || u.HasArg("--covheight") {
            for i, arg := range os.Args {
                if arg == "-covh" || arg == "--covheight" {
                    var err error
                    h, err = strconv.Atoi(os.Args[i+1])
                    if err != nil {
                        log.Fatalf("%sInvalid height argument%s\n",
                            s.BG_RED, s.RESET)
                    }
                }
            }
        }

        cover.Generate(w, h, amount, dirPath)
    }

    if x {
        fw.WriteBooksToXML(&books, dirPath)
    }

    if j {
        fw.WriteBooksToJSON(&books, dirPath)
    }

    if c {
        fw.WriteBooksToCSV(&books, dirPath)
    }
}
