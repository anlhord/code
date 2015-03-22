package main


func Int(a, b *int) int {
	return *a - *b
}

func Pstr(i *string) {
	print(*i)
}
func Pint(i *int) {
	print(*i)
}
func String(a, b *string) int {
	sa := *a
	sb := *b
	for i := 0; i < len(sa) && i < len(sb); i++ {
		if sa[i] != sb[i] {
			return int(sa[i] - sb[i])
		}
	}
	return 0
}

func Compare2(cmp func(*,*)(int), pnt func(*), l *,r *) {
	pnt(l)
	print(What(cmp(l,r)))
	pnt(r)
}

func What(n int) string {
	if n == 0 {
		return "="
	} else if n > 0 {
		return ">"
	}
	return "<"
}

func main() {
	s := []string{"foo","bar","zz","gg","aa","bb"}
	
	Compare2(String, &s[0], &s[1])
	Compare2(String, &s[2], &s[3])
	Compare2(String, &s[4], &s[5])
	Compare2(String, &s[6], &s[7])
	
}
