package standalone

import (
	"Booksgen/cmd/standalone/cover"
	b "Booksgen/internal/book"
	fw "Booksgen/internal/filewriter"
	s "Booksgen/pkg/style"
	"flag"
	"log"
	"os"
	"path/filepath"
)

// Run starts the standalone app. It uses the flag package to parse command-line
// arguments for output format (XML, JSON, CSV), output directory, book count,
// and optional cover generation.
func Run() {
	fs := flag.NewFlagSet("booksgen", flag.ExitOnError)

	// Mode flags — registered so the parser doesn't error when invoked via main.
	_ = fs.Bool("st", false, "")
	_ = fs.Bool("standalone", false, "")

	var exportXML, exportJSON, exportCSV, genCover bool
	fs.BoolVar(&exportXML, "x", false, "export XML")
	fs.BoolVar(&exportXML, "xml", false, "export XML")
	fs.BoolVar(&exportJSON, "j", false, "export JSON")
	fs.BoolVar(&exportJSON, "json", false, "export JSON")
	fs.BoolVar(&exportCSV, "c", false, "export CSV")
	fs.BoolVar(&exportCSV, "csv", false, "export CSV")
	fs.BoolVar(&genCover, "cov", false, "generate covers")
	fs.BoolVar(&genCover, "cover", false, "generate covers")

	var amount int
	fs.IntVar(&amount, "a", 1, "number of books to generate")
	fs.IntVar(&amount, "amount", 1, "number of books to generate")

	var dirPath string
	fs.StringVar(&dirPath, "o", "./output/", "output directory")
	fs.StringVar(&dirPath, "output", "./output/", "output directory")

	var covAmount, covW, covH int
	fs.IntVar(&covAmount, "cova", 1, "number of covers")
	fs.IntVar(&covAmount, "covamount", 1, "number of covers")
	fs.IntVar(&covW, "covw", 400, "cover width in pixels")
	fs.IntVar(&covW, "covwidth", 400, "cover width in pixels")
	fs.IntVar(&covH, "covh", 650, "cover height in pixels")
	fs.IntVar(&covH, "covheight", 650, "cover height in pixels")

	var covDir string
	fs.StringVar(&covDir, "covo", "./covoutput/", "cover output directory")
	fs.StringVar(&covDir, "covoutput", "./covoutput/", "cover output directory")

	fs.Parse(os.Args[1:])

	oDir, err := filepath.Abs(dirPath)
	if err != nil {
		log.Fatalf("%sError getting path to directory: %s\n%s%s\n",
			s.FG_RED, oDir, err, s.RESET)
	}
	log.Printf("Path to output directory: %s", oDir)

	books := b.GenerateBooks(uint32(amount))

	if genCover {
		cover.Generate(covW, covH, covAmount, covDir)
	}

	if exportXML {
		fw.WriteBooksToXML(&books, dirPath)
	}
	if exportJSON {
		fw.WriteBooksToJSON(&books, dirPath)
	}
	if exportCSV {
		fw.WriteBooksToCSV(&books, dirPath)
	}
}
