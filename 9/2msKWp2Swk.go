package main

func CompareInt(a, b *int) int {
	return *a - *b
}

func Max(compar func(*,*)int, items []*){
	j := 0
	for i := range items {
		if compar(items[i], items[j]) > 0 {
			j = i
		}
	}
	print("biggest:")
	print(j)
	print(". ")
}


func main() {
	var x, y, z int
	x = 35
	y = 64
	z = 96

	Max(CompareInt, &z)
	Max(CompareInt, &x)

}

