package main

func ConjugateCmplx(n *complex128) {
	*n = complex(real(*n), -imag(*n))
}
func ConjugateF64(n *float64) {
	//  NO-OP
}

func ConjugateTranspose2x2(conj func(*), mat [2][2]) {
	{ // swap pattern
	t := mat[0][1]
	mat[0][1] = mat[1][0]
	mat[1][0] = t
	}

	conj(&mat[0][0])
	conj(&mat[0][1])
	conj(&mat[1][0])
	conj(&mat[1][1])
}

func main() {
	var test1 = [2][2]complex128 {{1, 1-1i},{-2+1i, 0-1i}}
	ConjugateTranspose2x2(ConjugateCmplx, test1)
}