package main

type foo struct {
	list [2]*foo
	foo int
}

func list_init(list [2]*, item *) {
	list[0] = item
	list[1] = item
}

func main() {
	var x foo

	list_init(x.list, &x)
	print("hi")
}
