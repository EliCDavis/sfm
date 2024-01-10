package opensfm

import "github.com/EliCDavis/sfm"

// ReconstructionJson represents the reconstruction.json output by OpenSFM
type ReconstructionJsonSchema []ReconstructionSchema

type ReconstructionSchema struct {
	Cameras      map[string]CameraSchema `json:"cameras"`
	Shots        map[string]ShotSchema   `json:"shots"`
	Points       map[string]PointSchema  `json:"points"`
	Biases       map[string]BiasSchema   `json:"biases"`
	RigCameras   map[string]RigCamera    `json:"rig_cameras"`
	RigInstances map[string]RigInstance  `json:"rig_instances"`
	ReferenceLLA ReferenceLLA            `json:"reference_lla"`
}

type CameraSchema struct {
	ProjectionType string  `json:"projection_type"`
	Width          int     `json:"width"`
	Height         int     `json:"height"`
	Focal          float64 `json:"focal"`
	K1             float64 `json:"k1"`
	K2             float64 `json:"k2"`
}

type ShotSchema struct {
	Rotation    []float64 `json:"rotation"`
	Translation []float64 `json:"translation"`
	Orientation int       `json:"orientation"`
	Camera      string    `json:"camera"`
	CaptureTime float64   `json:"capture_time"`
	Scale       float64   `json:"scale"`
}

type PointSchema struct {
	Color       []float64 `json:"color"` // 0-255 float
	Coordinates []float64 `json:"coordinates"`
}

type BiasSchema struct {
	Rotation    []float64 `json:"rotation"`
	Translation []float64 `json:"translation"`
	Scale       float64   `json:"scale"`
}

type RigCamera struct {
	Rotation    []float64 `json:"rotation"`
	Translation []float64 `json:"translation"`
}

type RigInstance struct {
	Rotation     []float64         `json:"rotation"`
	Translation  []float64         `json:"translation"`
	RigCameraIDs map[string]string `json:"rig_camera_ids"`
}

type ReferenceLLA struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}

func LoadReconstruction(filename string) (ReconstructionJsonSchema, error) {
	return sfm.LoadJSONFile[ReconstructionJsonSchema](filename)
}
