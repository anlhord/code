package main

func CompareInt(a, b *int) int {
	return *a - *b
}

type RemoveMe int

func Max(compar func(*int, *int) int, items ...*RemoveMe) {
	j := 0
	for i := range items {
		if compar((*RemoveMe)(items[i]), (*RemoveMe)(items[j])) > 0 {
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
