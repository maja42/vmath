# vmath [![GoDoc](https://godoc.org/github.com/maja42/vmath?status.svg)](https://godoc.org/github.com/maja42/vmath)

vmath is a standalone vector math library for go, supporting both `float32` and `int` types.

Matrices are stored in column major order.


This library provides 2D, 3D and 4D vector and matrix types with an extensive set of operations and utility functions. \
Vectors are both provided for `float32` and `int`.
 
Additional support for quaternions is also provided.
 
vmath aims to provide all functionality needed for graphics development (eg. using *OpenGL*) in a highly performant matter.
It therefore also offers basic functionality related to spatial calculations. 

Related functionality from the standard `math` package that go only provides for `float64` is available for the `float32` type in this package.
 
## Related projects

This library is inspired by [glm](https://glm.g-truc.net/0.9.9/index.html), which is typically used by C or C++ projects.

**Alternative go packages:**

[MathGL](https://github.com/go-gl/mathgl) provides similar functionality for `float32` and `float64` types,
but does not support `int`. \
Besides `int` vectors, this package is missing some features like geometric utility functions, `float32` versions for functionality available in the `math` package, 
and a few methods that you might or might not need.

[Azul3D](https://github.com/azul3d) is a game engine written in go, that contains its own math functions in the sub-package *github.com/azul3d/engine/lmath*. \
However, Azul3D uses `float64` and is also missing some key features (like 2x2 matrices and `int` vector support). 
\
Furthermore, the subpackage is not intended to be used as a standalone library and Azul3D is also no longer maintained.  

## Contributions

Feel free to submit bug reports or pull requests for new features, examples or unit tests.
