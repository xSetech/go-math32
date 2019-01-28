// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go-math32

// Modf returns integer and fractional floating-point numbers
// that sum to f. Both values have the same sign as f.
//
// Special cases are:
//	Modf(±Inf) = ±Inf, NaN
//	Modf(NaN) = NaN, NaN
func Modf(f float32) (int float32, frac float32)

func modf(f float32) (int float32, frac float32) {
	if f < 1 {
		switch {
		case f < 0:
			int, frac = Modf(-f)
			return -int, -frac
		case f == 0:
			return f, f // Return -0, -0 when f == -0
		}
		return 0, f
	}

	x := Float32bits(f)
	e := uint(x>>shift)&mask - bias

	// Keep the top 12+e bits, the integer part; clear the rest.
	if e < 32-9 {
		x &^= 1<<(32-9-e) - 1
	}
	int = Float32frombits(x)
	frac = f - int
	return
}
