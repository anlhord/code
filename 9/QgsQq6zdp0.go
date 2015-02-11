package main

func CompareInt(a, b *int) int {
	return *a - *b
}

func Max(compar func(*,*)int, foo *, bar *){
	if compar(foo,bar) >0 {
		print("foo is bigger\n")
	}
	print("foo's not bigger\n")
}


func main() {
	var x, y, z int
	x = 35
	y = 64
	z = 96

	Max(CompareInt, &x, &z)

}

