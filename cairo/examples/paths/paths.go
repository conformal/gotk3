package main

import (
	"math"
	"github.com/conformal/gotk3/cairo"
)

func main() {
	surf := cairo.ImageSurfaceCreate(cairo.FORMAT_ARGB32, 256, 256)

	cr := cairo.Create(surf)

	// blank the canvas
	cr.SetSourceRGB(1, 1, 1)
	cr.Paint()

	// straight line
	cr.SetSourceRGB(0.3, 0.2, 0.8)
	cr.MoveTo(32, 32)
	cr.LineTo(224, 224)
	cr.Stroke()

	// curved line, with the inside of the curve filled
	cr.SetSourceRGB(0.8, 0.2, 0.8)
	cr.MoveTo(32, 224)
	cr.CurveTo(64, 64, 192, 192, 224, 32)
	cr.Fill()

	// square
	cr.SetSourceRGB(0.75, 0, 0)
	cr.Rectangle(64, 64, 128, 128)
	cr.Fill()

	// circle
	cr.SetSourceRGB(0.3, 0.8, 0.3)
	cr.Arc(128, 128, 64, 0, math.Pi * 2)
	cr.Stroke()

	// text
	cr.SetSourceRGB(0.3, 0.8, 0.8)
	cr.MoveTo(20, 20)
	cr.SelectFontFace("sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	cr.SetFontSize(16)
	cr.ShowText("Hello World!")

	surf.WriteToPNG("paths.png")
}
