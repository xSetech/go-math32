// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go-math32

// Copysign returns a value with the magnitude
// of x and the sign of y.
func Copysign(x, y float32) float32 {
	const sign = 1 << 31
	return Float32frombits(Float32bits(x)&^sign | Float32bits(y)&sign)
}
