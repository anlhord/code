package main

const Next = 0
const Prev = 1

// creates a new list
func mklist(set func(*,*,*), item *) {
	set(item, item, item)
}

// inserts item to list
insert(set func(*,*,*), get func(*)(**,**), list *, item *)
{
/*

	lext = list->next;	//get
	
	item->prev = list;	//set
	item->next = lext;
	
	iext = item->next;	//get
	lrev = list->prev;	//get
	
	list->prev = lrev;	//set
	list->next = item;

	iixt = iext->next;	/get

	iext->prev = item;
	iext->next = iixt;
*/

	lprev, lnext := get(list)
	set(item, list, lnext)
	_, enext := get(elm)
	set(list, list, lnext)
	
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

func get(l *foo, i int) (link **foo) {
	switch i {
	case 0: return &l.prev
	default: return &l.next
	}
}

func main() {
	var x foo

	mklist(set, &x)

	print("hello")
}

