package colmap

// All derived from here
// https://github.com/colmap/colmap/blob/8e7093d22b324ce71c70d368d852f0ad91743808/src/colmap/sensor/models.h

// ============================================================================

type SimplePinholeCamera Camera

func (spc SimplePinholeCamera) FocalLength() float64 { return spc.Params[0] }
func (spc SimplePinholeCamera) Cx() float64          { return spc.Params[1] }
func (spc SimplePinholeCamera) Cy() float64          { return spc.Params[2] }

// ============================================================================

type PinholeCamera Camera

func (spc PinholeCamera) Fx() float64 { return spc.Params[0] }
func (spc PinholeCamera) Fy() float64 { return spc.Params[1] }
func (spc PinholeCamera) Cx() float64 { return spc.Params[2] }
func (spc PinholeCamera) Cy() float64 { return spc.Params[3] }

// ============================================================================

type SimpleRadialCamera Camera

func (spc SimpleRadialCamera) FocalLength() float64 { return spc.Params[0] }
func (spc SimpleRadialCamera) Cx() float64          { return spc.Params[1] }
func (spc SimpleRadialCamera) Cy() float64          { return spc.Params[2] }
func (spc SimpleRadialCamera) K() float64           { return spc.Params[3] }

// ============================================================================

type RadialCamera Camera

func (spc RadialCamera) FocalLength() float64 { return spc.Params[0] }
func (spc RadialCamera) Cx() float64          { return spc.Params[1] }
func (spc RadialCamera) Cy() float64          { return spc.Params[2] }
func (spc RadialCamera) K1() float64          { return spc.Params[3] }
func (spc RadialCamera) K2() float64          { return spc.Params[4] }

// ============================================================================

type OpenCVCamera Camera

func (spc OpenCVCamera) Fx() float64 { return spc.Params[0] }
func (spc OpenCVCamera) Fy() float64 { return spc.Params[1] }
func (spc OpenCVCamera) Cx() float64 { return spc.Params[2] }
func (spc OpenCVCamera) Cy() float64 { return spc.Params[3] }
func (spc OpenCVCamera) K1() float64 { return spc.Params[4] }
func (spc OpenCVCamera) K2() float64 { return spc.Params[5] }
func (spc OpenCVCamera) P1() float64 { return spc.Params[6] }
func (spc OpenCVCamera) P2() float64 { return spc.Params[7] }

// ============================================================================

type OpenCVFishEyeCamera Camera

func (spc OpenCVFishEyeCamera) Fx() float64 { return spc.Params[0] }
func (spc OpenCVFishEyeCamera) Fy() float64 { return spc.Params[1] }
func (spc OpenCVFishEyeCamera) Cx() float64 { return spc.Params[2] }
func (spc OpenCVFishEyeCamera) Cy() float64 { return spc.Params[3] }
func (spc OpenCVFishEyeCamera) K1() float64 { return spc.Params[4] }
func (spc OpenCVFishEyeCamera) K2() float64 { return spc.Params[5] }
func (spc OpenCVFishEyeCamera) K3() float64 { return spc.Params[6] }
func (spc OpenCVFishEyeCamera) K4() float64 { return spc.Params[7] }

// ============================================================================

type FullOpenCVCamera Camera

func (spc FullOpenCVCamera) Fx() float64 { return spc.Params[0] }
func (spc FullOpenCVCamera) Fy() float64 { return spc.Params[1] }
func (spc FullOpenCVCamera) Cx() float64 { return spc.Params[2] }
func (spc FullOpenCVCamera) Cy() float64 { return spc.Params[3] }
func (spc FullOpenCVCamera) K1() float64 { return spc.Params[4] }
func (spc FullOpenCVCamera) K2() float64 { return spc.Params[5] }
func (spc FullOpenCVCamera) P1() float64 { return spc.Params[6] }
func (spc FullOpenCVCamera) P2() float64 { return spc.Params[7] }
func (spc FullOpenCVCamera) K3() float64 { return spc.Params[8] }
func (spc FullOpenCVCamera) K4() float64 { return spc.Params[9] }
func (spc FullOpenCVCamera) K5() float64 { return spc.Params[10] }
func (spc FullOpenCVCamera) K6() float64 { return spc.Params[11] }

// ============================================================================

type FovCamera Camera

func (spc FovCamera) Fx() float64    { return spc.Params[0] }
func (spc FovCamera) Fy() float64    { return spc.Params[1] }
func (spc FovCamera) Cx() float64    { return spc.Params[2] }
func (spc FovCamera) Cy() float64    { return spc.Params[3] }
func (spc FovCamera) Omega() float64 { return spc.Params[4] }

// ============================================================================

type SimpleRadialFisheyeCamera Camera

func (spc SimpleRadialFisheyeCamera) FocalLength() float64 { return spc.Params[0] }
func (spc SimpleRadialFisheyeCamera) Cx() float64          { return spc.Params[1] }
func (spc SimpleRadialFisheyeCamera) Cy() float64          { return spc.Params[2] }
func (spc SimpleRadialFisheyeCamera) K() float64           { return spc.Params[3] }

// ============================================================================

type RadialFisheyeCamera Camera

func (spc RadialFisheyeCamera) FocalLength() float64 { return spc.Params[0] }
func (spc RadialFisheyeCamera) Cx() float64          { return spc.Params[1] }
func (spc RadialFisheyeCamera) Cy() float64          { return spc.Params[2] }
func (spc RadialFisheyeCamera) K1() float64          { return spc.Params[3] }
func (spc RadialFisheyeCamera) K2() float64          { return spc.Params[4] }

// ============================================================================

type ThinPrismFisheyeCamera Camera

func (spc ThinPrismFisheyeCamera) Fx() float64  { return spc.Params[0] }
func (spc ThinPrismFisheyeCamera) Fy() float64  { return spc.Params[1] }
func (spc ThinPrismFisheyeCamera) Cx() float64  { return spc.Params[2] }
func (spc ThinPrismFisheyeCamera) Cy() float64  { return spc.Params[3] }
func (spc ThinPrismFisheyeCamera) K1() float64  { return spc.Params[4] }
func (spc ThinPrismFisheyeCamera) K2() float64  { return spc.Params[5] }
func (spc ThinPrismFisheyeCamera) P1() float64  { return spc.Params[6] }
func (spc ThinPrismFisheyeCamera) P2() float64  { return spc.Params[7] }
func (spc ThinPrismFisheyeCamera) K3() float64  { return spc.Params[8] }
func (spc ThinPrismFisheyeCamera) K4() float64  { return spc.Params[9] }
func (spc ThinPrismFisheyeCamera) Sx1() float64 { return spc.Params[10] }
func (spc ThinPrismFisheyeCamera) Sy1() float64 { return spc.Params[11] }
