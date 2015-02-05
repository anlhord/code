package main

func printnum(num *int) {
	print("NUM IS:")
	print(*num)
	print("\n")
}

func macro_visit(visitor func(*), obj *) {
	visitor(obj)
}

func main() {
	var num int = 9586
	_ = num

	macro_visit(printnum, &num)
}

