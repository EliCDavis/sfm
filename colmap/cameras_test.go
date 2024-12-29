package colmap_test

import (
	"testing"

	"github.com/EliCDavis/sfm/colmap"
	"github.com/stretchr/testify/assert"
)

func TestCameras(t *testing.T) {
	tests := map[string]struct {
		model         colmap.CameraModel
		parameters    int
		displayString string
	}{
		"SIMPLE_PINHOLE": {
			model:         colmap.SIMPLE_PINHOLE,
			parameters:    3,
			displayString: "Simple Pinhole",
		},
		"PINHOLE": {
			model:         colmap.PINHOLE,
			parameters:    4,
			displayString: "Pinhole",
		},
		"SIMPLE_RADIAL": {
			model:         colmap.SIMPLE_RADIAL,
			parameters:    4,
			displayString: "Simple Radial",
		},
		"RADIAL": {
			model:         colmap.RADIAL,
			parameters:    5,
			displayString: "Radial",
		},
		"OPENCV": {
			model:         colmap.OPENCV,
			parameters:    8,
			displayString: "OpenCV",
		},
		"OPENCV_FISHEYE": {
			model:         colmap.OPENCV_FISHEYE,
			parameters:    8,
			displayString: "OpenCV Fisheye",
		},
		"FULL_OPENCV": {
			model:         colmap.FULL_OPENCV,
			parameters:    12,
			displayString: "Full OpenCV",
		},
		"FOV": {
			model:         colmap.FOV,
			parameters:    5,
			displayString: "FOV",
		},
		"SIMPLE_RADIAL_FISHEYE": {
			model:         colmap.SIMPLE_RADIAL_FISHEYE,
			parameters:    4,
			displayString: "Simple Radial Fisheye",
		},
		"RADIAL_FISHEYE": {
			model:         colmap.RADIAL_FISHEYE,
			parameters:    5,
			displayString: "Radial Fisheye",
		},
		"THIN_PRISM_FISHEYE": {
			model:         colmap.THIN_PRISM_FISHEYE,
			parameters:    12,
			displayString: "Thin Prism Fisheye",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.displayString, tc.model.String())
			assert.Equal(t, tc.parameters, tc.model.NumParameters())
		})
	}
}
