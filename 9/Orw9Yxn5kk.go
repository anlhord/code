package main


func ConjugateCmplx(n *complex128) {
	*n = complex(real(*n), -imag(*n))
//	print(*n)
//	print(" ")
}
func ConjugateF64(n *float64) {
	//  NO-OP
}

func CTransposeF64(conj func(*), mat [2][2]float64) {
	t := b
	
}

func CTransposeF64(conj func(*), mat [2][2]complex128) {
	t := b
	b = 
}

func main() {
//	var a,b,c complex128 = 3-2i, 7+0i, 0+1i	
//	ConjugateCmplx(&a)
//	ConjugateCmplx(&b)
//	ConjugateCmplx(&c)
	
	var test1 = [2][2]complex128 {{1, 1-1i},{-2+1i, 0-1i}}
}
