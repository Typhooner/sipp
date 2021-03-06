// Copyright Raul Vera 2015-2020

// Package sgrad provides facilities for the computation and rendering of
// a finite-difference gradient image from a source SippImage and a 2x2 kernel.
package sgrad

import (
	"image"
	"math"
)

import (
	. "github.com/Causticity/sipp/scomplex"
	. "github.com/Causticity/sipp/simage"
)

// TODO: The below could all be reimplemented to use only floating-point
// arithmetic, with a conversion to complex only at the end. As it is now it
// goes back and forth unnecessarily. It could be done with integer arithmetic
// by adding a function to simage that returns the pixel value as an int, but
// that probably won't save much. This is all optimisation and should be done
// only with proper profiling and a specific performance target. Although
// integer-only arithmetic might also be useful to debug numerical issues.

// SippGradKernels are defined in the same way as images are stored in memory,
// i.e. in row-major order from the top-left corner down.
type SippGradKernel [2][2]complex128

var defaultKernel = SippGradKernel{
	{-1 + 0i, 0 + 1i},
	{0 - 1i, 1 + 0i},
}

func byKernel(kern SippGradKernel, pix, right, below, belowRight float64) complex128 {
	return kern[0][0]*complex(pix, 0) +
		kern[0][1]*complex(right, 0) +
		kern[1][0]*complex(below, 0) +
		kern[1][1]*complex(belowRight, 0)
}

// Use a 2x2 kernel to create a finite-differences complex gradient image, one
// pixel narrower and shorter than the original. We'd rather reduce the size of
// the output image than arbitrarily wrap around or extend the source image, as
// any such procedure could introduce errors into the statistics.
func FdgradKernel(src SippImage, kern SippGradKernel) (grad *ComplexImage) {
	// Create the dst image from the bounds of the src
	srect := src.Bounds()
	grad = new(ComplexImage)
	grad.Rect = image.Rect(0, 0, srect.Dx()-1, srect.Dy()-1)
	grad.Pix = make([]complex128, grad.Rect.Dx()*grad.Rect.Dy())
	grad.MaxMod = 0

	dsti := 0
	for y := 0; y < grad.Rect.Dy(); y++ {
		for x := 0; x < grad.Rect.Dx(); x++ {
			val := byKernel(kern, src.Val(x, y),
				src.Val(x+1, y), src.Val(x, y+1), src.Val(x+1, y+1))
			grad.Pix[dsti] = val
			dsti++
			modsq := real(val)*real(val) + imag(val)*imag(val)
			// store the maximum squared value, then take the root afterwards
			if modsq > grad.MaxMod {
				grad.MaxMod = modsq
			}
		}
	}
	grad.MaxMod = math.Sqrt(grad.MaxMod)

	return
}

// Use a default kernel to compute a finite-differences gradient. See
// FdgradKernel for details.
func Fdgrad(src SippImage) (grad *ComplexImage) {
	return FdgradKernel(src, defaultKernel)
}
