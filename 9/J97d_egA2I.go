package main

const Next = byte(0)
const Prev = byte(1)

func SetNext() {}

func wl_list_init(list *foo){
	list.prev = list;
	list.next = list;
}

func wl_list_insert(list *foo, elm *foo){
	elm.prev = list;
	elm.next = list.next;
	list.next = elm;
	elm.next.prev = elm;
}

func wl_list_remove(elm *foo){
	elm.prev.next = elm.next;
	elm.next.prev = elm.prev;
	elm.next = nil;
	elm.prev = nil;
}

func wl_list_length(list *foo) int {
	var e *foo
	var count int;

	count = 0;
	e = list.next;
	for (e != list) {
		e = e.next;
		count++;
	}

	return count;
}

func wl_list_empty(list *foo) bool {
	return list.next == list;
}

func wl_list_insert_list(list *foo, other *foo){
	if (wl_list_empty(other)) {
		return;
	}

	other.next.prev = list;
	other.prev.next = list.next;
	list.next.prev = other.prev;
	list.next = other.next;
}

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

// is list empty?
func Empty(get func(*, byte)(**), list *) bool {
	la := get(list, Next)
	lb := get(la, Prev)
	lc := get(lb, Prev)
	return lb == lc;
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
	var x, y, z, w, q, p foo
	x = y
	z = w
	p = q

	Mklist(set, &x)
	
	print(Empty(get, &x))
	Insert(set, get, &x, &y)

	print(Empty(get, &x))
	Insert(set, get, &x, &z)
	

	print(Empty(get, &x))
	Insert(set, get, &x, &w)


	print(Empty(get, &x))
	Remove(set, get, &y)
	
	print("hello")
}

