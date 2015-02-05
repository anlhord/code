package main

func Int(a, b *int) int {
	return *a - *b
}

func String(a, b *string) {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			continue
		}
		
	}
}

func main() {
	print("genericsppp")
}

