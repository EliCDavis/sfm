# SFM

Structs for interacting with reconstruction data from different SFM programs

* COLMAP
	* Cameras
	* Points
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