package main

func CompareInt(a, b *int) int {
	return *a - *b
}

type RemoveMe int

func Max(compar func(*int, *int) int, items ...*int) {
	j := 0
	for i := range items {
		var l *int = items[i]
		var r *int = items[j]
		if compar(l, r) > 0 {
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

	var a = [](*int){&x, &y, &z}
	_ = a
	Max(CompareInt, a...)

}
