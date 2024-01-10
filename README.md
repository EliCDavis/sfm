# SFM

Structs for interacting with reconstruction data from different SFM programs

## Meshroom

Interact with Meshroom's `cameras.sfm` file output from the `StructureFromMotion` step.

```golang
package example

import (
	"fmt"

	"github.com/EliCDavis/sfm/meshroom"
)

func TestLoad(t *testing.T) {
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

func TestLoad(t *testing.T) {
	reconstructions, err := opensfm.LoadReconstruction("MyOpenSFMProject/reconstruction.json")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File contains %d reconstructions", len(reconstructions))
}
```