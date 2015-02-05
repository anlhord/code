package main

func Int(a, b *int) int {
	return *a - *b
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

func Max(cmp func(*,*)(int), i *...) {
	print(What(cmp(l,r)))
}

func Compare2(cmp func(*,*)(int), l *,r *) {
	print(What(cmp(l,r)))
}

func What(n int) string {
	if n == 0 {
		return "same"
	} else if n > 0 {
		return "l > r"
	}
	return "l < r"
}

func main() {
	s := []string{"foo","bar","zz","gg","aa","bb"}
	
	Compare2(String, &s[0], &s[1])
}

