package main

// This is the list implementation

const Next = byte(0)
const Prev = byte(1)

func setnext(set func(*,*,byte), get func(*, byte)(**), elm, to *) {
	set(elm, to, Next)
}

func setprev(set func(*,*,byte), get func(*, byte)(**), elm, to *) {
	set(elm, to, Prev)
}

func Make(set func(*,*,byte), list *){
	set(list, list, 0)
	set(list, list, 1)
	
}

//
func Insert(set func(*,*,byte), get func(*, byte)(**), list *foo, elm *foo){
//	setprev(set, get, elm, list)
	elm.prev = list;
//	setnext(set, get, elm, get(list, Next))
	elm.next = list.next;
//	setnext(set, get, list, elm)
	list.next = elm;
//	setprev(set, get, get(elm, Next), elm)
	elm.next.prev = elm;
}

func Remove(set func(*,*,byte), get func(*, byte)(**), elm *){
	setnext(set, get, get(elm, Prev), get(elm, Next))
	setprev(set, get, get(elm, Next), get(elm, Prev))
	set(elm, nil, 0)
	set(elm, nil, 1)
}

func rmv(set func(*,*,byte), get func(*, byte)(**), elm *){
	setnext(set, get, get(elm, Prev), get(elm, Next))
	setprev(set, get, get(elm, Next), get(elm, Prev))
}

func Removexx(set func(*,*,byte), get func(*, byte)(**), elm *foo){





	print(elm.next)
	print(elm.prev.next)	// this becomes nil
	print(elm.next.prev)	// this becomes nil
	
	print(" ")

//	setnextdbg(set, get, get(elm, Prev), get(elm, Next))
	elm.prev.next = elm.next
	
	print(elm.next)
	print(elm.prev.next)
	print(elm.next.prev)

	print(" ")
//	setprev(set, get, get(elm, Next), get(elm, Prev))
//
	elm.next.prev = elm.prev
	
	print(elm.prev.next)
	print(elm.next.prev)
	
	print(" ")
	


	set(elm, nil, 1)

	set(elm, nil, 0)
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

