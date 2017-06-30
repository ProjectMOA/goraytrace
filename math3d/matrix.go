package math3d

// Matrix represents a 4x4 matrix
type Matrix struct {
	a, b, c, d,
	e, f, g, h,
	i, j, k, l,
	m, n, o, p float64
}

// MultiplyPoint returns point multiplied by the matrix
func (mat *Matrix) MultiplyPoint(point *Vector3) *Vector3 {
	x := mat.a*point.X + mat.b*point.Y + mat.c*point.Z + mat.d
	y := mat.e*point.X + mat.f*point.Y + mat.g*point.Z + mat.h
	z := mat.i*point.X + mat.j*point.Y + mat.k*point.Z + mat.l
	h := mat.m*point.X + mat.n*point.Y + mat.o*point.Z + mat.p
	return &Vector3{X: x / h, Y: y / h, Z: z / h}
}

// MultiplyVector returns vector multiplied by the matrix
func (mat *Matrix) MultiplyVector(vector *Vector3) *Vector3 {
	x := mat.a*vector.X + mat.b*vector.Y + mat.c*vector.Z
	y := mat.e*vector.X + mat.f*vector.Y + mat.g*vector.Z
	z := mat.i*vector.X + mat.j*vector.Y + mat.k*vector.Z
	return &Vector3{X: x, Y: y, Z: z}
}

// ComposeMatrix composes the two matrices multiplying them
func (mat *Matrix) ComposeMatrix(mat2 *Matrix) *Matrix {
	a := mat.a*mat2.a + mat.b*mat2.e + mat.c*mat2.i + mat.d*mat2.m
	b := mat.a*mat2.b + mat.b*mat2.f + mat.c*mat2.j + mat.d*mat2.n
	c := mat.a*mat2.c + mat.b*mat2.g + mat.c*mat2.k + mat.d*mat2.o
	d := mat.a*mat2.d + mat.b*mat2.h + mat.c*mat2.l + mat.d*mat2.p
	e := mat.e*mat2.a + mat.f*mat2.e + mat.g*mat2.i + mat.h*mat2.m
	f := mat.e*mat2.b + mat.f*mat2.f + mat.g*mat2.j + mat.h*mat2.n
	g := mat.e*mat2.c + mat.f*mat2.g + mat.g*mat2.k + mat.h*mat2.o
	h := mat.e*mat2.d + mat.f*mat2.h + mat.g*mat2.l + mat.h*mat2.p
	i := mat.i*mat2.a + mat.j*mat2.e + mat.k*mat2.i + mat.l*mat2.m
	j := mat.i*mat2.b + mat.j*mat2.f + mat.k*mat2.j + mat.l*mat2.n
	k := mat.i*mat2.c + mat.j*mat2.g + mat.k*mat2.k + mat.l*mat2.o
	l := mat.i*mat2.d + mat.j*mat2.h + mat.k*mat2.l + mat.l*mat2.p
	m := mat.m*mat2.a + mat.n*mat2.e + mat.o*mat2.i + mat.p*mat2.m
	n := mat.m*mat2.b + mat.n*mat2.f + mat.o*mat2.j + mat.p*mat2.n
	o := mat.m*mat2.c + mat.n*mat2.g + mat.o*mat2.k + mat.p*mat2.o
	p := mat.m*mat2.d + mat.n*mat2.h + mat.o*mat2.l + mat.p*mat2.p
	return &Matrix{a: a, b: b, c: c, d: d,
		e: e, f: f, g: g, h: h,
		i: i, j: j, k: k, l: l,
		m: m, n: n, o: o, p: p}
}
