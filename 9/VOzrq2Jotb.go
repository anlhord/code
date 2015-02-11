package main

// creates a new list
func mklist(set func(*,*,*), item *) {
	set(item, item, item)
}

// inserts item to list
wl_list_insert(set func(*,*,*), get func(*)(**,**), struct wl_list *list, *item)
{
elm->prev = list;
elm->next = list->next;
list->next = elm;
elm->next->prev = elm;
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

func get(l *foo) (prev **foo, next **foo) {
	return &l.prev, &l.next
}

func main() {
	var x foo

	mklist(set, &x)

	print("hello")
}

