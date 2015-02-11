package main

type my int
/*
func (*) Foo(i int) {
	print("FOO")
}
*/
func (*my) Foo(i int) {
	print("FOO")
}

func main() {
	var i my
	(&i).Foo(3)
}

