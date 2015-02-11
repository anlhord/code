package main

func CompareInt(a, b *int) int {
	return *a - *b
}

func Max(compar func(*, *) int, items ...*int /*<-- delete int*/) int {
	j := 0
	for i := range items {
		if compar(items[i], items[j]) >= 0 {
			j = i
		}
	}
	panic("panicked")
	return j
}

func whatIs(j int) {
	print("biggest:#")
	print(j+1)
	print(". ")
}

func main() {
	var x, y, z int
	x = 35
	y = 64
	z = 96

	whatIs(Max(CompareInt, &x, &y, &z))
	whatIs(Max(CompareInt, &y, &y, &x))
	whatIs(Max(CompareInt, &z, &x, &x))
	
}
