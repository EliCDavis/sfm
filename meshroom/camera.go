package meshroom

import (
	"io"

	"github.com/EliCDavis/sfm"
)

type Version []string

type View struct {
	ViewID      string `json:"viewId"`
	PoseID      string `json:"poseId"`
	FrameID     string `json:"frameId"`
	IntrinsicID string `json:"intrinsicId"`
	ResectionID string `json:"resectionId"`
	Path        string `json:"path"`
	Width       string `json:"width"`
	Height      string `json:"height"`
	Metadata    struct {
		ResolutionUnit string `json:"ResolutionUnit"`
		XResolution    string `json:"XResolution"`
		YResolution    string `json:"YResolution"`
		OiioColorSpace string `json:"oiio:ColorSpace"`
	} `json:"metadata"`
}

type Intrinsic struct {
	IntrinsicID                  string   `json:"intrinsicId"`
	Width                        string   `json:"width"`
	Height                       string   `json:"height"`
	SensorWidth                  string   `json:"sensorWidth"`
	SensorHeight                 string   `json:"sensorHeight"`
	SerialNumber                 string   `json:"serialNumber"`
	Type                         string   `json:"type"`
	InitializationMode           string   `json:"initializationMode"`
	InitialFocalLength           string   `json:"initialFocalLength"`
	FocalLength                  string   `json:"focalLength"`
	PixelRatio                   string   `json:"pixelRatio"`
	PixelRatioLocked             string   `json:"pixelRatioLocked"`
	PrincipalPoint               []string `json:"principalPoint"`
	DistortionInitializationMode string   `json:"distortionInitializationMode"`
	DistortionParams             []string `json:"distortionParams"`
	UndistortionOffset           []string `json:"undistortionOffset"`
	UndistortionParams           string   `json:"undistortionParams"`
	Locked                       string   `json:"locked"`
}

type Pose struct {
	PoseID string `json:"poseId"`
	Pose   struct {
		Transform struct {
			Rotation []string `json:"rotation"`
			Center   []string `json:"center"`
		} `json:"transform"`
		Locked string `json:"locked"`
	} `json:"pose"`
}

type Cameras struct {
	Version         Version     `json:"version"`
	FeaturesFolders []string    `json:"featuresFolders"`
	MatchesFolders  []string    `json:"matchesFolders"`
	Views           []View      `json:"views"`
	Intrinsics      []Intrinsic `json:"intrinsics"`
	Poses           []Pose      `json:"poses"`
}

func ReadCameras(in io.Reader) (Cameras, error) {
	return sfm.ReadJSON[Cameras](in)
}

func LoadCameras(filename string) (Cameras, error) {
	return sfm.LoadJSONFile[Cameras](filename)
}
