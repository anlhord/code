package main

// all args should be same type pointers
func macro_mismatch(a, b, c, d, e, f *) {
}

func main() {

	var a,b,c,d,e,f int
	a = b
	c = d
	e = f

	var x byte
	_ = x

	// the compiler is smart. it checks all are same type (ints)
	
	macro_mismatch(&a, &b, &x, &d, &e, &f)
}