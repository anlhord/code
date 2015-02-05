package main

func try1(a *, b *byte) {
	print("1IT ")
}

func main() {
	var xa, xb int
	xa = xb
	var xz byte
	_ = xz

	macro_reverse_fm1(&xa, &xz, try1)	// int wildcard
}

