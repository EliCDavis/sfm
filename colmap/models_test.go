package colmap_test

import (
	"testing"

	"github.com/EliCDavis/sfm/colmap"
	"github.com/stretchr/testify/assert"
)

func TestSimplePinholeCamera(t *testing.T) {
	cam := colmap.SimplePinholeCamera(colmap.Camera{
		Params: []float64{1, 2, 3},
	})

	assert.Equal(t, 1., cam.FocalLength())
	assert.Equal(t, 2., cam.Cx())
	assert.Equal(t, 3., cam.Cy())
}

func TestPinholeCamera(t *testing.T) {
	cam := colmap.PinholeCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4},
	})

	assert.Equal(t, 1., cam.Fx())
	assert.Equal(t, 2., cam.Fy())
	assert.Equal(t, 3., cam.Cx())
	assert.Equal(t, 4., cam.Cy())
}

func TestSimpleRadialCamera(t *testing.T) {
	cam := colmap.SimpleRadialCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4},
	})

	assert.Equal(t, 1., cam.FocalLength())
	assert.Equal(t, 2., cam.Cx())
	assert.Equal(t, 3., cam.Cy())
	assert.Equal(t, 4., cam.K())
}

func TestRadialCamera(t *testing.T) {
	cam := colmap.RadialCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4, 5},
	})

	assert.Equal(t, 1., cam.FocalLength())
	assert.Equal(t, 2., cam.Cx())
	assert.Equal(t, 3., cam.Cy())
	assert.Equal(t, 4., cam.K1())
	assert.Equal(t, 5., cam.K2())
}

func TestOpenCVCamera(t *testing.T) {
	cam := colmap.OpenCVCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4, 5, 6, 7, 8},
	})

	assert.Equal(t, 1., cam.Fx())
	assert.Equal(t, 2., cam.Fy())
	assert.Equal(t, 3., cam.Cx())
	assert.Equal(t, 4., cam.Cy())
	assert.Equal(t, 5., cam.K1())
	assert.Equal(t, 6., cam.K2())
	assert.Equal(t, 7., cam.P1())
	assert.Equal(t, 8., cam.P2())
}

func TestOpenCVFishEyeCamera(t *testing.T) {
	cam := colmap.OpenCVFishEyeCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4, 5, 6, 7, 8},
	})

	assert.Equal(t, 1., cam.Fx())
	assert.Equal(t, 2., cam.Fy())
	assert.Equal(t, 3., cam.Cx())
	assert.Equal(t, 4., cam.Cy())
	assert.Equal(t, 5., cam.K1())
	assert.Equal(t, 6., cam.K2())
	assert.Equal(t, 7., cam.K3())
	assert.Equal(t, 8., cam.K4())
}

func TestFullOpenCVCamera(t *testing.T) {
	cam := colmap.FullOpenCVCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	})

	assert.Equal(t, 1., cam.Fx())
	assert.Equal(t, 2., cam.Fy())
	assert.Equal(t, 3., cam.Cx())
	assert.Equal(t, 4., cam.Cy())
	assert.Equal(t, 5., cam.K1())
	assert.Equal(t, 6., cam.K2())
	assert.Equal(t, 7., cam.P1())
	assert.Equal(t, 8., cam.P2())
	assert.Equal(t, 9., cam.K3())
	assert.Equal(t, 10., cam.K4())
	assert.Equal(t, 11., cam.K5())
	assert.Equal(t, 12., cam.K6())
}

func TestFovCamera(t *testing.T) {
	cam := colmap.FovCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4, 5},
	})

	assert.Equal(t, 1., cam.Fx())
	assert.Equal(t, 2., cam.Fy())
	assert.Equal(t, 3., cam.Cx())
	assert.Equal(t, 4., cam.Cy())
	assert.Equal(t, 5., cam.Omega())
}

func TestSimpleRadialFisheyeCamera(t *testing.T) {
	cam := colmap.SimpleRadialFisheyeCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4},
	})

	assert.Equal(t, 1., cam.FocalLength())
	assert.Equal(t, 2., cam.Cx())
	assert.Equal(t, 3., cam.Cy())
	assert.Equal(t, 4., cam.K())
}

func TestRadialFisheyeCamera(t *testing.T) {
	cam := colmap.RadialFisheyeCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4, 5},
	})

	assert.Equal(t, 1., cam.FocalLength())
	assert.Equal(t, 2., cam.Cx())
	assert.Equal(t, 3., cam.Cy())
	assert.Equal(t, 4., cam.K1())
	assert.Equal(t, 5., cam.K2())
}

func TestThinPrismFisheyeCamera(t *testing.T) {
	cam := colmap.ThinPrismFisheyeCamera(colmap.Camera{
		Params: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	})

	assert.Equal(t, 1., cam.Fx())
	assert.Equal(t, 2., cam.Fy())
	assert.Equal(t, 3., cam.Cx())
	assert.Equal(t, 4., cam.Cy())
	assert.Equal(t, 5., cam.K1())
	assert.Equal(t, 6., cam.K2())
	assert.Equal(t, 7., cam.P1())
	assert.Equal(t, 8., cam.P2())
	assert.Equal(t, 9., cam.K3())
	assert.Equal(t, 10., cam.K4())
	assert.Equal(t, 11., cam.Sx1())
	assert.Equal(t, 12., cam.Sy1())
}
