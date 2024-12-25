package colmap_test

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/EliCDavis/sfm/colmap"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestRenderDepthmap(t *testing.T) {
	depthmap, err := colmap.LoadDepthmap("depthmap.bin")
	check(err)

	maxV := depthmap.MaxValue() / 100
	img := image.NewRGBA(image.Rect(0, 0, depthmap.Width, depthmap.Height))
	for x := 0; x < depthmap.Width; x++ {
		for y := 0; y < depthmap.Height; y++ {
			v := byte(min((depthmap.Value(x, y)/maxV), 1) * 255)
			img.Set(x, y, color.RGBA{
				R: v,
				G: v,
				B: v,
				A: 255,
			})
		}
	}

	f, err := os.Create("depth.png")
	check(err)
	defer f.Close()
	check(png.Encode(f, img))
}

func TestRenderNormalmap(t *testing.T) {
	normalMap, err := colmap.LoadNormalmap("normal.bin")
	check(err)

	img := image.NewRGBA(image.Rect(0, 0, normalMap.Width, normalMap.Height))
	for x := 0; x < normalMap.Width; x++ {
		for y := 0; y < normalMap.Height; y++ {
			v := normalMap.Value(x, y)

			// https://github.com/colmap/colmap/blob/e924f0f825c6033a0cf2b89fa79c585236d4bf3a/src/colmap/mvs/normal_map.cc#L112
			img.Set(x, y, color.RGBA{
				R: uint8((v.X() + 1) * 0.5 * 255),
				G: uint8((v.Y() + 1) * 0.5 * 255),
				B: uint8(v.Z() * -255.),
				A: 255,
			})
		}
	}

	f, err := os.Create("normals.png")
	check(err)
	defer f.Close()
	writer := bufio.NewWriter(f)
	check(png.Encode(writer, img))
	check(writer.Flush())
}
