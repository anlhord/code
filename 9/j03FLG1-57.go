package main


func ConjugateCmplx(n *complex128) {
	*n = complex(real(*n), -imag(*n))
//	print(*n)
//	print(" ")
}
func ConjugateF64(n *float64) {
	//  NO-OP
}

func CTransposeF64_2_2(conj func(*), mat [2][2]float64) {
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

func CTransposeC128_2_2(conj func(*), mat [2][2]complex128) {
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
//	var a,b,c complex128 = 3-2i, 7+0i, 0+1i	
//	ConjugateCmplx(&a)
//	ConjugateCmplx(&b)
//	ConjugateCmplx(&c)
	
	var test1 = [2][2]complex128 {{1, 1-1i},{-2+1i, 0-1i}}
	CTransposeC128_2_2(ConjugateCmplx, test1)
	print(test1[0][0])
	print(test1[0][1])
	print(test1[1][0])
	print(test1[1][1])
	
}
