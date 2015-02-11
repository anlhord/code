package main

func CompareInt(a, b *int) int {
	return *a - *b
}

func Max(compar func(*, *) int, items ...) int {
	j := 0
	for i := range items {
		if compar(&items[i], &items[j]) >= 0 {
			j = i
		}
	}
	return j
}

func whatIs(j int) {
	print("biggest:#")
	print(j+1)
	print(". ")
}

func main() {
	whatIs(Max(CompareInt, 7, 5, 9))
	whatIs(Max(CompareInt, 8, 4, 2))
	whatIs(Max(CompareInt, 5, 3, 1))
	
}
