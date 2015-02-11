package main

const Next = byte(0)
const Prev = byte(1)

func wl_list_init(list *foo){
	list->prev = list;
	list->next = list;
}
func wl_list_insert(list *foo, elm *foo){
	elm->prev = list;
	elm->next = list->next;
	list->next = elm;
	elm->next->prev = elm;
}

func wl_list_remove(elm *foo){
	elm->prev->next = elm->next;
	elm->next->prev = elm->prev;
	elm->next = NULL;
	elm->prev = NULL;
}

func wl_list_length(list *foo){
	struct wl_list *e;
	int count;

	count = 0;
	e = list->next;
	while (e != list) {
		e = e->next;
		count++;
	}

	return count;
}

func wl_list_empty(list *foo) int{
	return list->next == list;
}

func wl_list_insert_list(list *foo, struct wl_list *other){
	if (wl_list_empty(other))
		return;

	other->next->prev = list;
	other->prev->next = list->next;
	list->next->prev = other->prev;
	list->next = other->next;
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

