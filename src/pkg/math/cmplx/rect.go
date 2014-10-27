// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

import "math"

// Rect returns the complex number x with polar coordinates r, θ.

// Rect 返回极坐标形式 (r, θ) 的复数 x。
func Rect(r, θ float64) complex128 {
	s, c := math.Sincos(θ)
	return complex(r*c, r*s)
}