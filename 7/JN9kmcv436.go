package main

type foo struct {
	next *foo
	prev *foo
	foo int
}

func list_init(next **, prev **, item *) {
	*next = item
	*prev = item
}

func main() {
	var x foo

	list_init(&x.next, &x.prev, &x)
	print("hi")
}

