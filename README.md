# SFM

Structs for interacting with reconstruction data from different SFM programs

* COLMAP
	* Cameras
	* Points
	* Images
	* Depthmap
	* Normalmap
* OpenSFM
	* reconstruction.json
* Meshroom
	* cameras.sfm

## COLMAP

### Point Data

Interact with COLMAPS's `points3D.bin` file output under the sparse reconstruction data.

```golang
package example

import (
	"fmt"

	"github.com/EliCDavis/sfm/colmap"
)

func main() {
	points, err := colmap.LoadPoints3DBinary("ColmapProject/sparse/0/points3D.bin")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File contains %d points", len(points))
}
```

### Cameras

Interact with COLMAPS's `cameras.bin` file output under the sparse reconstruction data.

```golang
package example

import (
	"fmt"

	"github.com/EliCDavis/sfm/colmap"
)

func main() {
	cameras, err := colmap.LoadCamerasBinary("ColmapProject/sparse/0/cameras.bin")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File contains %d cameras", len(cameras))
}
```

### Images

Interact with COLMAPS's `images.bin` file output under the sparse reconstruction data.

```golang
package example

import (
	"fmt"

	"github.com/EliCDavis/sfm/colmap"
)

func main() {
	images, err := colmap.LoadImagesBinary("ColmapProject/sparse/0/images.bin")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File contains %d images", len(images))
}
```

### Depth maps

Create a Image representing the depthmap data

```golang
package main

import (
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

func main() {
	depthmap, err := colmap.LoadDepthmap("depthmap.bin")
	check(err)

	maxV := depthmap.MaxValue()
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
```

### Normal maps

Create a Image representing the normalmap data


```golang
package main

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

func main() {
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
```

## Meshroom

Interact with Meshroom's `cameras.sfm` file output from the `StructureFromMotion` step.

```golang
package example

import (
	"fmt"

	"github.com/EliCDavis/sfm/meshroom"
)

func main() {
	cameras, err := meshroom.LoadCameras("MeshroomProject/MeshroomCache/StructureFromMotion/abc123/cameras.sfm")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File contains %d Poses", len(cameras.Poses))
}
```

## OpenSFM

Interact with OpenSFM's `reconstruction.json` file output from the `reconstruct` step.

```golang
package example

import (
	"fmt"

	"github.com/EliCDavis/sfm/opensfm"
)

func main() {
	reconstructions, err := opensfm.LoadReconstruction("MyOpenSFMProject/reconstruction.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File contains %d reconstructions", len(reconstructions))
}
```