package editimg

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/clusters"
)

func labToHex(point []float64) string {
	c := colorful.Lab(point[0], point[1], point[2])
	c = c.Clamped()

	return c.Hex()
}

func GetColorsOfCenters(clusters []clusters.Cluster) []string{
	hexColors := make([]string, len(clusters))
	for k, v := range clusters {
		point := v.Center.Coordinates()
		hex := labToHex(point)
		hexColors[k] = hex
	}
	return hexColors
}
