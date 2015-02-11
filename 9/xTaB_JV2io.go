package main

const Next = byte(0)
const Prev = byte(1)

// creates a new list
func Mklist(set func(*,*,*), item *) {
	set(item, item, item)
}

// inserts item to list
func Insert(set func(*,*,*), get func(*, byte)(**), list *, item *) {
	set(item, list, get(item, Next))
	set(item, get(item, Prev), get(list, Next))
	set(list, get(list, Prev), item)
	set(get(item, Next), item, get(get(item, Next), Next))
}

// remove item from list.
func Remove(set func(*,*,*), get func(*, byte)(**), item *) {
	set(get(item, Prev), get(get(item, Prev), Prev), get(item, Next))
	set(get(item, Next), get(item, Prev), get(get(item, Next), Next))
	set(item, nil, nil)
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
	var x, y, z, w foo
	x = y
	z = w

	Mklist(set, &x)
	Insert(set, get, &x, &y)
	Remove(set, get)
	
	print("hello")
}

