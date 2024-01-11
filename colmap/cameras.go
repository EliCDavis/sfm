package colmap

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/EliCDavis/bitlib"
)

type CameraModel int

const (
	SIMPLE_PINHOLE CameraModel = iota
	PINHOLE
	SIMPLE_RADIAL
	RADIAL
	OPENCV
	OPENCV_FISHEYE
	FULL_OPENCV
	FOV
	SIMPLE_RADIAL_FISHEYE
	RADIAL_FISHEYE
	THIN_PRISM_FISHEYE
)

func (cm CameraModel) NumParameters() int {
	switch cm {
	case SIMPLE_PINHOLE:
		return 3

	case PINHOLE:
		return 4

	case SIMPLE_RADIAL:
		return 4

	case RADIAL:
		return 5

	case OPENCV:
		return 8

	case OPENCV_FISHEYE:
		return 8

	case FULL_OPENCV:
		return 12

	case FOV:
		return 5

	case SIMPLE_RADIAL_FISHEYE:
		return 4

	case RADIAL_FISHEYE:
		return 5

	case THIN_PRISM_FISHEYE:
		return 12
	}
	panic(fmt.Errorf("unimplemented Camera Model Parameter Count %d", cm))
}

type Camera struct {
	ID     int
	Model  CameraModel
	Width  uint64
	Height uint64
	Params []float64
}

func ReadCamerasBinary(filename string) ([]Camera, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bitlib.NewReader(f, binary.LittleEndian)

	numCameras := reader.UInt64()
	cameras := make([]Camera, numCameras)

	for i := uint64(0); i < numCameras; i++ {
		c := Camera{}

		c.ID = int(reader.Int32())
		c.Model = CameraModel(reader.Int32())
		c.Width = reader.UInt64()
		c.Height = reader.UInt64()

		paramCount := c.Model.NumParameters()
		c.Params = make([]float64, paramCount)
		for pI := 0; pI < paramCount; pI++ {
			c.Params[pI] = reader.Float64()
		}

		cameras[i] = c
	}

	return cameras, reader.Error()
}
