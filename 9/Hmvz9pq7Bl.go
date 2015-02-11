package main

const Next = byte(0)
const Prev = byte(1)


func Make(set func(*, byte, *), list *){
	set(list, Prev, list)
	set(list, Next, list)
}
func Insert(get func(*foo, byte) (*foo), set func(*foo, byte, *foo), list *foo, elm *foo){
	elm.prev = list;
	elm.next = get(list, Next)
	list.next = elm;
	elm.next.prev = elm;
}

func Remove(get func(*foo, byte) (*foo), set func(*foo, byte, *foo), elm *foo){
	elm.prev.next = get(elm,Next)
	elm.next.prev = get(elm,Prev)
	elm.next = nil;
	elm.prev = nil;
}

func Len(list *foo) int {
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

func Empty(get func(*foo, byte) (*foo), list *foo) bool{
	return list.next == list;
}

func InsertList(get func(*foo, byte) (*foo), set func(*foo, byte, *foo), list *foo, other *foo){
	if (Empty(get, other)) {
		return;
}
	other.next.prev = list;
	other.prev.next = list.next;
	list.next.prev = other.prev;
	list.next = other.next;
}

// for testing

func Dump(list *foo) int {
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

func set(l *foo, what byte, to *foo) {
	if what == Next {
		l.next = to
	} else {
		l.prev = to
	}
}

func get(l *foo, what byte) (link *foo) {
	if what == Next {
		return l.next
	}
	return l.prev
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
	
//	Insert(get, set, &listhead, &y)

	print(Empty(get, &listhead))
	print(" ")
	
	Insert(get, set, &listhead, &z)

	print(Empty(get, &listhead))
	print(" ")
	
	Insert(get, set, &listhead, &w)

	Dump(&listhead)

	print(Empty(get, &listhead))
	print(" ")

//	Remove(get, set, &y)
	
//	print(Empty(get, &listhead))
// 	print(" ")

	Remove(get, set, &w)	// this damages it
	
	print(Empty(get, &listhead))
	print(" ")
	

	
	if z.prev.next == nil {
		print("ELM PREV NEXT IS NULL")
		return
	}

	Remove(get, set, &z)
	
	print(Empty(get, &listhead))
	print(" ")
	
	
	print("helloworld")
}