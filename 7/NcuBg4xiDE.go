package main

func macro_pass(ptr *) (ret *) {
	return ptr
}

func main() {

	var foo int
	var bar string
	var baz byte
	
	a := macro_pass(&foo)
	b := macro_pass(&bar)
	c := macro_pass(&baz)
	
	foo = a
	bar = b
	baz = c
	
}

