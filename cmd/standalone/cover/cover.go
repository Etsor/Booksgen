package cover

import (
	"Booksgen/internal/book"
	fw "Booksgen/internal/filewriter"
	"Booksgen/res"
	s "Booksgen/pkg/style"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

// Paths are relative to the embedded FS root (res/).
var fonts = [...]string{
	"fonts/129_kosmos/129 KOSMOS.ttf",
	"fonts/big_scratch_brush/BigScratchBrush.ttf",
	"fonts/bold_pixels/BoldPixels.ttf",
	"fonts/cossette_texte/CossetteTexte-Regular.ttf",
	"fonts/kosolapa_script/KosolapaScript-Regular.ttf",
	"fonts/manufacturing_consent/fonts/ttf/ManufacturingConsent-Regular.ttf",
	"fonts/southgetto/SOUTHGHETTO.ttf",
	"fonts/super_malibu/Super Malibu.ttf",
}

func Generate(width, height, amount int, dirPath string) {
	fw.CreateDir(dirPath)
	books := book.GenerateBooks(uint32(amount))

	for i := range amount {
		dc := gg.NewContext(width, height)

		dc.SetColor(
			color.RGBA{
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				255},
		)

		dc.Clear()

		dc.SetRGB(0, 0, 0)
		font := fonts[rand.Intn(len(fonts))]

		dc.DrawRectangle(0, 0, float64(width), float64(height/8))
		dc.DrawRectangle(0, float64(height)-float64(height/8), float64(width), float64(height/8))
		dc.SetRGB(0, 0, 0)
		dc.Fill()

		// AUTHOR
		loadFont(dc, font, 32)
		drawCenteredText(dc, books[i].Author,
			float64(height/2)-float64(height/4),
			float64(width))

		// TITLE
		loadFont(dc, font, 48)
		drawCenteredText(dc, books[i].Title,
			float64(height/2)-float64(height/8),
			float64(width))

		// GENRE
		loadFont(dc, font, 32)
		drawCenteredText(dc, books[i].Genre,
			float64(height/2)+float64(height/8),
			float64(width))

		// PUBLISHER
		loadFont(dc, font, 16)
		drawCenteredText(dc, books[i].Publisher,
			float64(height/2)+float64(height/4),
			float64(width))

		file := filepath.Join(dirPath, fmt.Sprintf("cover_%d.png", i+1))
		dc.SavePNG(file)

		log.Printf("Generated cover: %s\nUsed font: %s\n\n", file, font)
	}
}

// loadFont reads a TTF from the embedded FS and sets it on the drawing context.
// This is equivalent to gg.LoadFontFace but works regardless of the working directory.
func loadFont(dc *gg.Context, fontPath string, size float64) {
	data, err := res.Fonts.ReadFile(fontPath)
	if err != nil {
		log.Fatalf("%sError reading embedded font %s: %s%s\n",
			s.FG_RED, fontPath, err, s.RESET)
	}

	f, err := truetype.Parse(data)
	if err != nil {
		log.Fatalf("%sError parsing embedded font %s: %s%s\n",
			s.FG_RED, fontPath, err, s.RESET)
	}

	dc.SetFontFace(truetype.NewFace(f, &truetype.Options{Size: size}))
}

func drawCenteredText(dc *gg.Context, text string, offset_y, width float64) {
	dc.DrawStringWrapped(
		text,
		float64(width/2), offset_y,
		0.5, 0.5,
		float64(width)-40,
		1.5,
		gg.AlignCenter,
	)
}
