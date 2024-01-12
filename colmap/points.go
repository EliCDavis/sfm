package colmap

import (
	"bufio"
	"encoding/binary"
	"image/color"
	"io"
	"os"

	"github.com/EliCDavis/bitlib"
	"github.com/EliCDavis/vector/vector3"
)

// Everything transcribed from
// https://github.com/colmap/colmap/blob/main/scripts/python/read_write_model.py
// https://docs.python.org/3/library/struct.html

type Point3DTrack struct {
	ImageID   int
	Point2DID int
}

type Point3D struct {
	ID       uint64
	Position vector3.Float64
	Color    color.RGBA
	Error    float64
	Tracks   []Point3DTrack // as (IMAGE_ID, POINT2D_IDX)
}

func ReadPoints3DBinary(in io.Reader) ([]Point3D, error) {
	reader := bitlib.NewReader(in, binary.LittleEndian)

	numPoints := reader.UInt64()
	points := make([]Point3D, numPoints)

	for i := uint64(0); i < numPoints; i++ {
		p := Point3D{}

		p.ID = reader.UInt64()

		p.Position = vector3.New(
			reader.Float64(),
			reader.Float64(),
			reader.Float64(),
		)

		c := color.RGBA{A: 255}
		c.R = reader.Byte()
		c.G = reader.Byte()
		c.B = reader.Byte()
		p.Color = c

		p.Error = reader.Float64()

		numTracks := reader.UInt64()
		p.Tracks = make([]Point3DTrack, numTracks)
		for trackIndex := uint64(0); trackIndex < numTracks; trackIndex++ {
			track := Point3DTrack{}
			track.ImageID = int(reader.Int32())
			track.Point2DID = int(reader.Int32())
			p.Tracks[trackIndex] = track
		}

		points[i] = p
	}

	return points, reader.Error()
}

func LoadPoints3DBinary(filename string) ([]Point3D, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadPoints3DBinary(bufio.NewReader(f))
}
