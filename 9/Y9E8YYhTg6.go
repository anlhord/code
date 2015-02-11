package main

func ConjugateCmplx(n *complex128) {
	print(real(*n))
	print(imag(*n))
}
func ConjugateF64(n *float64) {
	// NO-OP
}


func main() {
	var a,b,c complex128 = 3-2i, 7, -i
	
	
	
	ConjugateCmplx(&a)
	ConjugateCmplx(&b)
	ConjugateCmplx(&c)
	
}
