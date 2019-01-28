# go-math32

The [standard golang math package](https://github.com/golang/go/tree/release-branch.go1.11/src/math) converted to only use [IEEE 754 single-precision](https://en.wikipedia.org/wiki/IEEE_754#Basic_and_interchange_formats) (32-bits).

## Rationale and excluded functionality

This was made while experimenting with an unreleased project compiled using [TinyGo](https://github.com/aykevl/tinygo) that targets a 32-bit architecture.

Some files were deleted because I wasn't going to use them in the original project, and thus did not feel like porting them:

- acosh.go
- asinh.go
- atanh.go
- cbrt.go
- cmplx/cmath\_test.go
- expm1.go
- gamma.go
- lgamma.go
- nextafter.go
- pow10.go

They may be added at a later date.

# License

The [BSD 3-clause from upstream](https://golang.org/LICENSE).
