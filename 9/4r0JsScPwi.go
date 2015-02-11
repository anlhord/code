package main

func main() {
	print("generics")
}

func ConjugateTranspose2x2(conj func(*), mat [2][2]) {
	{ // swap pattern
	t := mat[0][1]
	mat[0][1] = mat[1][0]
	mat[1][0] = t
	}

	conj(&mat[0][0])
	conj(&mat[0][1])
	conj(&mat[1][0])
	conj(&mat[1][1])
}