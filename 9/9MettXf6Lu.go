package main

const nihongo = "日本語€"
const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	var foo = []rune(nihongo)
	var bar = []byte(nihongo)
	var baz = string(nihongo)

	print(foo)
	print(bar)
	print(baz)
}
