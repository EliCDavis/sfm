package colmap

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"

	"github.com/EliCDavis/bitlib"
	"github.com/EliCDavis/vector/vector2"
	"github.com/EliCDavis/vector/vector3"
	"github.com/EliCDavis/vector/vector4"
)

type ImagePoint struct {
	Id       int64
	Position vector2.Float64
}

type Image struct {
	Id          int
	CameraId    int
	Name        string
	Rotation    vector4.Float64
	Translation vector3.Float64
	Points      []ImagePoint
}

func ReadImagesBinary(in io.Reader) ([]Image, error) {
	reader := bitlib.NewReader(in, binary.LittleEndian)

	numImages := reader.UInt64()
	images := make([]Image, numImages)

	for i := uint64(0); i < numImages; i++ {
		img := Image{}
		img.Id = int(reader.Int32())

		qx := reader.Float64()
		qy := reader.Float64()
		qz := reader.Float64()
		qw := reader.Float64()
		img.Rotation = vector4.New(qx, qy, qz, qw)

		tx := reader.Float64()
		ty := reader.Float64()
		tz := reader.Float64()
		img.Translation = vector3.New(tx, ty, tz)

		img.CameraId = int(reader.Int32())

		// Read Name
		nameData := make([]byte, 0)
		currentCharacter := reader.Byte()
		for currentCharacter != 0 {
			nameData = append(nameData, currentCharacter)
			currentCharacter = reader.Byte()
		}
		img.Name = string(nameData)

		numPoints := reader.UInt64()
		points := make([]ImagePoint, numPoints)

		for i := uint64(0); i < numPoints; i++ {
			x := reader.Float64()
			y := reader.Float64()
			points[i] = ImagePoint{
				Id:       reader.Int64(),
				Position: vector2.New(x, y),
			}
		}

		img.Points = points
		images[i] = img
	}

	return images, reader.Error()
}

func LoadImagesBinary(filename string) ([]Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadImagesBinary(bufio.NewReader(f))
}
