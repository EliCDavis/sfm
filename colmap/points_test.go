package colmap_test

import (
	"fmt"
	"testing"

	"github.com/EliCDavis/sfm/colmap"
)

func TestLoad(t *testing.T) {
	points, err := colmap.LoadCamerasBinary("C:/dev/sfm/chris2023/sparse/0/cameras.bin")

	if err != nil {
		panic(err)
	}

	if len(points) == 0 {
		t.Error("ass poses")
	}

	fmt.Printf("File contains %d poses", len(points))
}
