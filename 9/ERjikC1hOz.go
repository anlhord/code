package main

// all args should be same type pointers
func macro_mismatch(a, b, c, d, e, f *) {
	print("ARGUMENTS ARE OK")
}

func main() {

	var a,b,c,d,e,f int
	a = b
	c = d
	e = f

	var x byte
	_ = x

	// the compiler is smart. it checks all are same type (ints)
	// see http://our-gol-842.appspot.com/p/4mYZMJHFe3
	macro_mismatch(&a, &b, &c, &d, &e, &f)
}