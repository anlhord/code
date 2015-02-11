package main

type foo struct {
	next *foo
	prev *foo
	foo int
}

func main() {
	print("hello")
}

