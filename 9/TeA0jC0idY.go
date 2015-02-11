package main

func String(a, b *string) int {
sa := *a
sb := *b
for i := 0; i < len(sa) && i < len(sb); i++ {
if sa[i] != sb[i] {
return int(sa[i] - sb[i])
}
}
return 0
}

func CompareInt(a, b *int) int {
	return *a - *b
}

func Max(compar func(*, *) int, items ...*int /*<-- delete int*/) int {
	j := 0
	for i := range items {
		if compar(items[i], items[j]) > 0 {
			j = i
		}
	}
	return j
}

func whatIs(j int) {
	print("biggest:#")
	print(j)
	print(". ")
}

func main() {
	var x, y, z int
	x = 35
	y = 64
	z = 96

	whatIs(Max(CompareInt, &x, &y, &z))

	whatIs(Max(CompareInt, &x, &y, &z))
}
