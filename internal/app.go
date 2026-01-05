// Package app start this application
package app

import (
	editimg "github.com/jaskierrr/color-tool/internal/edit_img"
	"github.com/jaskierrr/color-tool/internal/clustering"
)

func Run() error {
	img, err := editimg.OpenImg()
	if err != nil {
		return err
	}

	points := editimg.GetLabColorPoints(img)
	_, _ = clustering.GetClusters(points, 8)

	return nil
}
