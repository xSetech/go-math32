// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go-math32

// Log10 returns the decimal logarithm of x.
// The special cases are the same as for Log.
func Log10(x float32) float32

func log10(x float32) float32 {
	return Log(x) * (1 / Ln10)
}

// Log2 returns the binary logarithm of x.
// The special cases are the same as for Log.
func Log2(x float32) float32

func log2(x float32) float32 {
	frac, exp := Frexp(x)
	// Make sure exact powers of two give an exact answer.
	// Don't depend on Log(0.5)*(1/Ln2)+exp being exactly exp-1.
	if frac == 0.5 {
		return float32(exp - 1)
	}
	return Log(frac)*(1/Ln2) + float32(exp)
}
