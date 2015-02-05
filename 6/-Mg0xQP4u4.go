package main

func macro_pass(ptr *) (ret *) {
	return ptr
}

func main() {

	var foo int
	var bar string
	var baz byte
	
	foo = *macro_pass(&foo)
	bar = *macro_pass(&bar)
	baz = *macro_pass(&baz)
	

}

