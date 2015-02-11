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

func Make(set func(*,*,*), list *){
	set(list, list, list)
}

func Insert(set func(*,*,*), get func(*, byte)(**), list *, elm *){
	setprev(set, get, elm, list)
//	elm.prev = list;
	setnext(set, get, elm, get(list, Next))
//	elm.next = list.next;
	setnext(set, get, list, elm)
//	list.next = elm;
	setprev(set, get, get(elm, Next), elm)
//	elm.next.prev = elm;
}

func Remove(set func(*,*,*), get func(*, byte)(**), elm *){
	setnext(set, get, get(elm, Prev), get(elm, Next))
//	elm.prev.next = elm.next;
	setprev(set, get, get(elm, Next), get(elm, Prev))
//	elm.next.prev = elm.prev;
	set(elm, nil, nil)
//	elm.next = nil;
//	elm.prev = nil;
}

func Len(list *foo) int {
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

func Empty(get func(*, byte)(**), list *foo) bool {
	return list.next == list;
}

func Append(set func(*,*,*), get func(*, byte)(**), list *foo, other *foo){
	if (Empty(get, other)) {
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


	Make(set, &x)
	
	print(Empty(get, &x))
	print(" ")
	Insert(set, get, &x, &y)

	print(Empty(get, &x))
	print(" ")
	Insert(set, get, &x, &z)
	

	print(Empty(get, &x))
	print(" ")
	Insert(set, get, &x, &w)


	print(Empty(get, &x))
	print(" ")
	Remove(set, get, &y)
	
	print("helloworld")
}

