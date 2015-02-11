package main

var ConjugateFloat func(*) = nil

func ConjugateCmplx128(n *complex128) {
	*n = complex(real(*n), -imag(*n))
}

func ConjugateTranspose(conj func(*), mat *[2][2]) {

	{ // swap pattern
	t := (*mat)[0][1]
	(*mat)[0][1] = (*mat)[1][0]
	(*mat)[1][0] = t
	}

	if conj != nil {
	conj(&(*mat)[0][0])
	conj(&(*mat)[0][1])
	conj(&(*mat)[1][0])
	conj(&(*mat)[1][1])
	}
}
