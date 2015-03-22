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
//
func Insert(set func(*,*,*), get func(*, byte)(**), list *foo, elm *foo){
//	setprev(set, get, elm, list)
	elm.prev = list;
//	setnext(set, get, elm, get(list, Next))
	elm.next = list.next;
//	setnext(set, get, list, elm)
	list.next = elm;
//	setprev(set, get, get(elm, Next), elm)
	elm.next.prev = elm;
}

func Removed(set func(*,*,*), get func(*, byte)(**), elm *){
	setnext(set, get, get(elm, Prev), get(elm, Next))
//	elm.prev.next = elm.next;
	setprev(set, get, get(elm, Next), get(elm, Prev))
//	elm.next.prev = elm.prev;
	set(elm, nil, nil)
//	elm.next = nil;
//	elm.prev = nil;
}

// toto je zle
func rmv(set func(*,*,*), get func(*, byte)(**), elm *){
	setnext(set, get, get(elm, Prev), get(elm, Next))
	setprev(set, get, get(elm, Next), get(elm, Prev))
}

func Remove(set func(*,*,*), get func(*, byte)(**), elm *foo){

//	setnext(set, get, get(elm, Prev), get(elm, Next))
//	setprev(set, get, get(elm, Next), get(elm, Prev))

	elm.prev.next = elm.next
	elm.next.prev = elm.prev

	set(elm, nil, nil)
//	elm.next = nil
//	elm.prev = nil
}

func Empty(get func(*, byte)(**), list *foo) bool {
	return list.next == list;
}

func Append(set func(*,*,*), get func(*, byte)(**), list *foo, other *foo){
	if (Empty(get, other)) {
		return;
	}
//	setprev(set, get, get(elm, Next), get(elm, Prev))
	other.next.prev = list;
//	setnext(set, get, get(elm, Prev), get(elm, Next))
	other.prev.next = list.next;
//	setprev(set, get, get(list, Next), get(other, Prev))
	list.next.prev = other.prev;
//	setnext(set, get, list, get(other, Next))
	list.next = other.next;
}

func wl_list_dump(list *foo) int {
	var e *foo
	var count int

	count = 0
	e = list.next;
	for (e != list) {
		print(e.foo)
		e = e.next;
		count++;
	}

	return count;
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

//
func main() {
	var listhead, x, y, z, w, q, p foo
	x = y
	z = w
	p = q
	
	y.foo = 1
	z.foo = 2
	w.foo = 3

	Make(set, &listhead)
	
	print(Empty(get, &listhead))
	print(" ")
	
	Insert(set, get, &listhead, &y)

	print(Empty(get, &listhead))
	print(" ")
	
	Insert(set, get, &listhead, &z)
	

	print(Empty(get, &listhead))
	print(" ")
	
	Insert(set, get, &listhead, &w)

	wl_list_dump(&listhead)

	print(Empty(get, &listhead))
	print(" ")

	Remove(set, get, &y)
	
	print(Empty(get, &listhead))
	print(" ")

	Remove(set, get, &w)
	
	print(Empty(get, &listhead))
	print(" ")

	Remove(set, get, &z)
	
	print(Empty(get, &listhead))
	print(" ")
	
	
	print("helloworld")
}

