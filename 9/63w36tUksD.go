package main

type foo struct {
	prev *foo
	next *foo
	foo int
}

func mklist(prev, next **, item *) {
	*prev = item
	*next = item
}

func main() {
	var x foo

	mklist(&x.prev, &x.next, &x)

	print("hello")
}

