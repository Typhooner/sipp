// Copyright Raul Vera 2016

// Tests for package scomplex

package scomplex

import (
	"image"
	"reflect"
	"testing"
)

import (
	. "github.com/Causticity/sipp/simage"
	. "github.com/Causticity/sipp/sipptesting"
	. "github.com/Causticity/sipp/sipptesting/sipptestcore"
)

var shiftedPic = []complex128{
	1, -2, 3, -4,
	-5, 6, -7, 8,
	9, -10, 11, -12,
	-13, 14, -15, 16,
}

var scaledShiftedPic = []uint8{
	131, 106, 148, 90,
	82, 172, 65, 189,
	197, 41, 213, 24,
	16, 238, 0, 255,
}

var smallZeroPic = []uint8{
	0, 0, 0, 0,
	0, 0, 0, 0,
	0, 0, 0, 0,
	0, 0, 0, 0,
}

var GrayZero = image.Gray{
	smallZeroPic,
	4,
	image.Rectangle{image.Point{0, 0}, image.Point{4, 4}},
}

var SgrayZero *SippGray

func TestFromComplex(t *testing.T) {
	cpx := FromComplexArray(CosxCosyTinyGrad, 19)
	if !reflect.DeepEqual(cpx.Pix, CosxCosyTinyGrad) {
		t.Error("Error: Complex image array differs from the one constructed from")
	}
	rect := image.Rect(0, 0, 19, 19)
	if !reflect.DeepEqual(cpx.Rect, rect) {
		t.Errorf("Error: Complex image rect incorrect, expected %v, got %v\n",
			rect, cpx.Rect)
	}
	if cpx.MaxMod != CosxCosyTinyGradMaxMod {
		t.Errorf("Error: Incorrect max modulus. Expected: %v, got %v",
			CosxCosyTinyGradMaxMod, cpx.MaxMod)
	}
}

func TestComplex(t *testing.T) {
	SgrayZero = new(SippGray)
	SgrayZero.Gray = &GrayZero

	comp := ToShiftedComplex(Sgray)
	if !reflect.DeepEqual(shiftedPic, comp.Pix) {
		t.Error("Shifted complex doesn't match Gray!")
	}

	comp = ToShiftedComplex(Sgray16)
	if !reflect.DeepEqual(shiftedPic, comp.Pix) {
		t.Error("Shifted complex doesn't match Gray16!")
	}

	re, im := comp.Render()

	if !reflect.DeepEqual(re.Pix(), scaledShiftedPic) {
		t.Error("real unexpected")
	}

	if !reflect.DeepEqual(im, SgrayZero) {
		t.Error("imaginary not zero")
	}
}
