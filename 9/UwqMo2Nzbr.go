package main

const Next = byte(0)
const Prev = byte(1)

// creates a new list
func mklist(set func(*,*,*), item *) {
	set(item, item, item)
}

// inserts item to list
func insert(set func(*,*,*), get func(*, byte)(**), list *, item *) {
	set(item, list, get(Next, item))
	set(item, get(Prev, item), get(Next, list))
	set(list, get(Prev, list), item)
	set(get(Next, item), item, get(Next, get(Next, item)))
}

// the custom struct

type foo struct {
	prev *foo
	next *foo
	foo int
}

func set(l, prev, next *foo) {
	l.prev = prev
	l.next = next
}

func get(l *foo, i byte) (link **foo) {
	if i == Next {
		return &l.next
	}
	return &l.prev
}

// the main

func main() {
	var x foo

	mklist(set, &x)

	print("hello")
}

