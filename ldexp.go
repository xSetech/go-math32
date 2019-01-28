// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go-math32

// Ldexp is the inverse of Frexp.
// It returns frac × 2**exp.
//
// Special cases are:
//	Ldexp(±0, exp) = ±0
//	Ldexp(±Inf, exp) = ±Inf
//	Ldexp(NaN, exp) = NaN
func Ldexp(frac float32, exp int) float32

func ldexp(frac float32, exp int) float32 {
	// special cases
	switch {
	case frac == 0:
		return frac // correctly return -0
	case IsInf(frac, 0) || IsNaN(frac):
		return frac
	}
	frac, e := normalize(frac)
	exp += e
	x := Float32bits(frac)
	exp += int(x>>shift)&mask - bias
	if exp < -150 { //
		return Copysign(0, frac) // underflow
	}
	if exp > 127 { // overflow
		if frac < 0 {
			return Inf(-1)
		}
		return Inf(1)
	}
	var m float32 = 1
	if exp < -126 { // denormal
		exp += 24
		m = 1.0 / (1 << 24) // 2**-24
	}
	x &^= 0x7F800000 // mask << shift
	x |= uint32(exp+bias) << shift
	return m * Float32frombits(x)
}
