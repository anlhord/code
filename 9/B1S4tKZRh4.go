package main

type foo struct {
	bar *int
	baz *byte
}

type bar struct {
	
}

func generizer(x bar) {
	print("hello")
}

func main() {

	var x int = 7
	var y byte = 5

	var f foo
	f.bar = &x
	f.baz = &y

	generizer(f)
}

