package main

type my int

func Bar(i *my) {
	print("BAR")
}
func Bar(i *) {
	print("BAR")
}

func main() {
	var i my
	(&i).Foo(3)
	Bar(&i)
}

