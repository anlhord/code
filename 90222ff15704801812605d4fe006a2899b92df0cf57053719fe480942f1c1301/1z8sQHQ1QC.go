package main

type foo struct {
	next *foo
	prev *foo
	foo int
}

func mklist(next, prev **, item *)

func main() {
	var x foo
	
	mklist(&(x.next), &(x.prev), &x)
	
	print("hello")
}

