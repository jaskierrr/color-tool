// Package clustering provides functions for calculate cluster centers
package clustering

import (
	"fmt"

	e "github.com/jaskierrr/color-tool/internal/entities"
	c "github.com/muesli/clusters"
	"github.com/muesli/kmeans"
	"github.com/jaskierrr/color-tool/internal/edit_img"
)

func GetClusters(points []e.LabPoint, clusterCount int) (clusters c.Clusters, err error) {
	obs := make(c.Observations, len(points))
	for k, v := range points {
		obs[k] = c.Coordinates{v.L, v.A, v.B}
	}
	km := kmeans.New()
	clusters, err = km.Partition(obs, clusterCount)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("clusters: %+v", clusters)
	// fmt.Print("\n------------------------------\n")
	hexColors := editimg.GetColorsOfCenters(clusters)
	fmt.Printf("hexColors: %+v", hexColors)
	// fmt.Print("\n------------------------------\n")
	// centers := clusters.CentersInDimension(0)
	// fmt.Printf("centers: %+v", centers)
	return clusters, nil
}

 // [#1a3252 #ff6100 #ba4491 #ffbc10 #ff071a #74640e #4b715c #00000e]
 // [#7fffc5 #0094ff #ff00dc #ff006e #4cff00 #ffe97f #ff6a00 #4800ff]
 // [#5e81ac #a3be8c #c5aec5 #ebcb8b #bf616a #8ab4c0 #2e3440 #d08770]

