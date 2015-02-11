package main

const Next = byte(0)
const Prev = byte(1)

func Make(list *foo){
	list.prev = list;
	list.next = list;
}
func Insert(list *foo, elm *foo){
	elm.prev = list;
	elm.next = list.next;
	list.next = elm;
	elm.next.prev = elm;
}

func Remove(elm *foo){
	elm.prev.next = elm.next;
	elm.next.prev = elm.prev;
	elm.next = nil;
	elm.prev = nil;
}

func Len(list *foo){
	var e *foo;
	var count int;

	count = 0;
	e = list.next;
	for (e != list) {
		e = e.next;
		count++;
	}

	return count;
}

func Empty(list *foo) int{
	return list.next == list;
}

func InsertList(list *foo, other *foo){
	if (Empty(other)) {
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

func set(l, to *foo, i byte) {
	if i == Next {
		l.next = to
	} else {
		l.prev = to
	}
}

func get(l *foo, i byte) (link **foo) {
	if i == Next {
		return &l.next
	}
	return &l.prev
}

func main() {
	var listhead, x, y, z, w, q, p foo
	x = y
	z = w
	p = q
	
	y.foo = 1
	z.foo = 2
	w.foo = 3

	Make(set, &listhead)
	
//	print(Empty(get, &listhead))
//	print(" ")
	
//	Insert(set, get, &listhead, &y)

	print(Empty(get, &listhead))
	print(" ")
	
	Insert(set, get, &listhead, &z)

	print(Empty(get, &listhead))
	print(" ")
	
	Insert(set, get, &listhead, &w)

	wl_list_dump(&listhead)

	print(Empty(get, &listhead))
	print(" ")

//	Remove(set, get, &y)
	
//	print(Empty(get, &listhead))
// 	print(" ")

	Remove(set, get, &w)	// this damages it
	
	print(Empty(get, &listhead))
	print(" ")
	

	
	if z.prev.next == nil {
		print("ELM PREV NEXT IS NULL")
		return
	}

	Remove(set, get, &z)
	
	print(Empty(get, &listhead))
	print(" ")
	
	
	print("helloworld")
}

