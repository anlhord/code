package main

import "mm"

func ConjugateCmplx(n *complex128) {
	*n = complex(real(*n), -imag(*n))
}
func ConjugateF64(n *float64) {
	// NO-OP
}


func main() {
	var a,b,c complex128 = 3-2i, 7+0i, 0-1i
	
	
	
	ConjugateCmplx(&a)
	ConjugateCmplx(&b)
	ConjugateCmplx(&c)
	
}
