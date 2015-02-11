package main

func (*) Foo(i int) {
	print("FOO")
}

func (*int) Foo(i int) {
	print("FOO")
}

func main() {
	var i int
	(&i).Foo(3)
}

