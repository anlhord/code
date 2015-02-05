package main

func main(a, b *) {
	if a < b {
		return a
	}
	return b
}

func main() {
	min(1, 2)
	min('a', 'b')
	min("a", "b")
}
