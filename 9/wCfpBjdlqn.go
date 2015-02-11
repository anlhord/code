package main

type foo struct {
	link [2]*foo
	foo int
}

func mklist(link *[2]*, item *) {

}

func main() {
	var x foo
	
	mklist(&x.link, &x)
	
	print("hello")
}

