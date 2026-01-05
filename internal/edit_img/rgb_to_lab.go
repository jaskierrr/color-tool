package editimg

import (
	"image"
	"image/color"

	e "github.com/jaskierrr/color-tool/internal/entities"
	"github.com/lucasb-eyer/go-colorful"
)

func rgbToLab(c color.Color) colorful.Color {
	r, g, b, _ := c.RGBA()
	f := func(color uint32) float64 {return float64(color) / 65535.0}
	LabColor := colorful.Color{R: f(r), G: f(g), B: f(b)}
	return LabColor 
}

func GetLabColorPoints(img image.Image) []e.LabPoint {
	colorPoints := []e.LabPoint{}
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			color := rgbToLab(c)
			l, a, b := color.Lab()
			point := e.ToLabPoint(l, a, b)
			colorPoints = append(colorPoints, point)
		}
	}
	return colorPoints
}
