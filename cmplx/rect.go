// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

import "go-math32"

// Rect returns the complex number x with polar coordinates r, θ.
func Rect(r, θ float32) complex128 {
	s, c := math.Sincos(θ)
	return complex(r*c, r*s)
}
