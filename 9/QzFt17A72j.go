package main

func (*) Foo(i int) {
	print("FOO")
}

func main() {
	var i,j,k,l int
	i=j
	k=l
	&i.Foo(3)
}

