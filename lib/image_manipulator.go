package lib

import (
	"image"
	"image/color"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

func AddTextToImage(img *image.NRGBA, text string) {
	addLabelHelper(img, 20, 30, text)
}

func addLabelHelper(img *image.NRGBA, x, y int, label string) {
	col := color.RGBA{0, 0, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}