package cover

import (
	b "Booksgen/internal/book"
	fw "Booksgen/internal/filewriter"
	s "Booksgen/pkg/style"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"path/filepath"

	"github.com/fogleman/gg"
)

var fonts = [...]string{
	"res/fonts/129_kosmos/129 KOSMOS.ttf",
	"res/fonts/big_scratch_brush/BigScratchBrush.ttf",
	"res/fonts/bold_pixels/BoldPixels.ttf",
	"res/fonts/cossette_texte/CossetteTexte-Regular.ttf",
	"res/fonts/kosolapa_script/KosolapaScript-Regular.ttf",
	"res/fonts/manufacturing_consent/fonts/ttf/ManufacturingConsent-Regular.ttf",
	"res/fonts/southgetto/SOUTHGHETTO.ttf",
	"res/fonts/super_malibu/Super Malibu.ttf",
}

func Generate(width, height, amount int, dirPath string) {
	fw.CreateDir(dirPath)
	books := b.GenerateBooks(uint32(amount))

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

func loadFont(dc *gg.Context, fontPath string, size float64) {
	err := dc.LoadFontFace(fontPath, size)
	if err != nil {
		log.Fatalf("%sError loading font: %s, %s%s\n",
			s.FG_RED, fontPath, err, s.RESET)
	}
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
