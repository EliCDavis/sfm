package meshroom_test

import (
	"fmt"
	"testing"

	"github.com/EliCDavis/sfm/meshroom"
	"github.com/EliCDavis/sfm/opensfm"
)

func TestLoad(t *testing.T) {
	cameras, err := meshroom.LoadCameras("C:/dev/sfm/logandrone2/MeshroomCache/StructureFromMotion/dd5f94dab4b71111a50d76ba6b97eb9478dc6cb6/cameras.sfm")

	if err != nil {
		panic(err)
	}

	if len(cameras.Poses) == 0 {
		t.Error("ass poses")
	}

	fmt.Printf("File contains %d poses", len(cameras.Poses))

	x, err := opensfm.LoadReconstruction("reconstruction.json")
	fmt.Printf("File contains %d reconstructions", len(x))
}
