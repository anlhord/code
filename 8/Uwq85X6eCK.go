package main

func Int(a, b *int) int {
	return *a - *b
}

func String(a, b *string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if (*a)[i] != (*b)[i] {
			return (*a)[i] - (*b)[i]
		}
	}
	return 0
}

func main() {
	print("genericsppp")
}

