package main


// all args should be same type pointers
func macro_mismatch(a, b, c, d, e, f *) {
}

func main() {


	var a,b,c,d,e,f int
	
	var x byte


	// change x to c to make it work
	macro_mismatch(&a, &b, &x, &d, &e, &f)
}

