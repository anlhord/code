package main

func CompareInt(a, b *int) int {
	return *a - *b
}

func Max(compar func(*,*)int, items ...*){
	j := 0
	for i := range 
	if compar(foo,bar) >0 {
		print("foo is bigger ")
	} else {
		print("foo's not bigger ")
	}
}


func main() {
	var x, y, z int
	x = 35
	y = 64
	z = 96

	Max(CompareInt, &x, &z)
	Max(CompareInt, &y, &x)

}

