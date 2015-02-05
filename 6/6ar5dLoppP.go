package main

func printnum(num *int) {
	print("NUM IS:")
	print(*num)

}

func macro_visit(visitor func(*), obj *) {
	visitor(obj)
}

func main() {
	var num int = 9586
	_ = num
	print("hi")
//	macro_visit(printnum, &num)
}

