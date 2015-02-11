package main


func wl_list_init(struct wl_list *list){
	list->prev = list;
	list->next = list;
}
func wl_list_insert(struct wl_list *list, struct wl_list *elm){
	elm->prev = list;
	elm->next = list->next;
	list->next = elm;
	elm->next->prev = elm;
}

func wl_list_remove(struct wl_list *elm){
	elm->prev->next = elm->next;
	elm->next->prev = elm->prev;
	elm->next = NULL;
	elm->prev = NULL;
}

func wl_list_length(struct wl_list *list){
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

func wl_list_empty(struct wl_list *list) int{
	return list->next == list;
}

func wl_list_insert_list(struct wl_list *list, struct wl_list *other){
	if (wl_list_empty(other))
		return;

	other->next->prev = list;
	other->prev->next = list->next;
	list->next->prev = other->prev;
	list->next = other->next;
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

