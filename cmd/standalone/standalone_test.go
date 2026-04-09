package standalone

import (
	"os"
	"testing"
)

// setArgs temporarily replaces os.Args and returns a restore function.
func setArgs(args []string) func() {
	orig := os.Args
	os.Args = args
	return func() { os.Args = orig }
}

func TestRun_NoFormatFlags(t *testing.T) {
	// With no -x/-j/-c flags, Run generates books but writes no files.
	// This should complete without panicking.
	restore := setArgs([]string{"booksgen"})
	defer restore()

	Run()
}

func TestRun_WithAmount(t *testing.T) {
	restore := setArgs([]string{"booksgen", "-a", "3"})
	defer restore()

	Run()
}

func TestRun_WithAmountLong(t *testing.T) {
	restore := setArgs([]string{"booksgen", "--amount", "2"})
	defer restore()

	Run()
}

func TestRun_WithOutputDir(t *testing.T) {
	dir := t.TempDir()
	restore := setArgs([]string{"booksgen", "-o", dir})
	defer restore()

	// No format flags — nothing written, but path parsing must not panic.
	Run()
}

func TestRun_WriteCSV(t *testing.T) {
	dir := t.TempDir()
	restore := setArgs([]string{"booksgen", "-c", "-a", "2", "-o", dir})
	defer restore()

	Run()

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir error: %v", err)
	}
	if len(entries) == 0 {
		t.Error("Expected a CSV file to be written in output dir")
	}
}

func TestRun_WriteJSON(t *testing.T) {
	dir := t.TempDir()
	restore := setArgs([]string{"booksgen", "-j", "-a", "1", "-o", dir})
	defer restore()

	Run()

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir error: %v", err)
	}
	if len(entries) == 0 {
		t.Error("Expected a JSON file to be written in output dir")
	}
}

func TestRun_WriteXML(t *testing.T) {
	dir := t.TempDir()
	restore := setArgs([]string{"booksgen", "-x", "-a", "1", "-o", dir})
	defer restore()

	Run()

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir error: %v", err)
	}
	if len(entries) == 0 {
		t.Error("Expected an XML file to be written in output dir")
	}
}
