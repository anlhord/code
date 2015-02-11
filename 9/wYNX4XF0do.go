package main

// This is the list implementation

const Next = byte(0)
const Prev = byte(1)

func setnext(set func(*,*,*), get func(*, byte)(**), elm, to *) {
	t := get(elm, Prev)
	set(elm, t, to)
}

func setprev(set func(*,*,*), get func(*, byte)(**), elm, to *) {
	t := get(elm, Next)
	set(elm, to, t)
}

func wl_list_make(set func(*,*,*), list *){
	setprev(set, list)
	setnext(set, list)
}

func wl_list_insert(set func(*,*,*), get func(*, byte)(**), list *foo, elm *foo){
	elm.prev = list;
	elm.next = list.next;
	list.next = elm;
	elm.next.prev = elm;
}

func wl_list_remove(set func(*,*,*), get func(*, byte)(**), elm *foo){
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

func wl_list_empty(get func(*, byte)(**), list *foo) bool {
	return list.next == list;
}

func wl_list_insert_list(set func(*,*,*), get func(*, byte)(**), list *foo, other *foo){
	if (wl_list_empty(get, other)) {
		return;
	}

	other.next.prev = list;
	other.prev.next = list.next;
	list.next.prev = other.prev;
	list.next = other.next;
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

//. the main

func main() {
	var x, y, z, w, q, p foo
	x = y
	z = w
	p = q


	wl_list_make(set, &x)
	
	print(wl_list_empty(get, &x))
	print(" ")
	wl_list_insert(set, get, &x, &y)

	print(wl_list_empty(get, &x))
		print(" ")
	wl_list_insert(set, get, &x, &z)
	

	print(wl_list_empty(get, &x))
		print(" ")
	wl_list_insert(set, get, &x, &w)


	print(wl_list_empty(get, &x))
		print(" ")
	wl_list_remove(set, get, &y)
	
	print("helloworld")
}

