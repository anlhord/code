package main

type foo struct {
	prev *foo
	next *foo
	foo int
}

func get(l, prev, next *foo) {
	l.prev = prev
	l.next = next
}

func set(l *foo) (prev **foo, next **foo) {
	return &l.prev, &l.next
}

func mklist(set func(*,*,*), item *) {
	set(item, item, item)
}

func main() {
	var x foo

	mklist(set, &x)

	print("hello")
}

