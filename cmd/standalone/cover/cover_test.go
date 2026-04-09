package cover

import (
	"Booksgen/res"
	"testing"

	"github.com/fogleman/gg"
)

func TestFonts_NonEmpty(t *testing.T) {
	if len(fonts) == 0 {
		t.Error("fonts slice must not be empty")
	}
}

func TestFonts_NoDuplicates(t *testing.T) {
	seen := make(map[string]bool)
	for _, f := range fonts {
		if seen[f] {
			t.Errorf("Duplicate font entry: %q", f)
		}
		seen[f] = true
	}
}

func TestFonts_AllEmbedded(t *testing.T) {
	for _, fontPath := range fonts {
		if _, err := res.Fonts.ReadFile(fontPath); err != nil {
			t.Errorf("font missing from embedded FS: %q: %v", fontPath, err)
		}
	}
}

func TestDrawCenteredText_DoesNotPanic(t *testing.T) {
	dc := gg.NewContext(400, 600)
	// No font loaded — gg uses its default; DrawStringWrapped must not panic.
	drawCenteredText(dc, "Test Title", 300, 400)
	drawCenteredText(dc, "", 100, 400)
}
