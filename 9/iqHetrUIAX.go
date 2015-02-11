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

func Compare2(cmp func(*,*)(int), l *,r *) {
	print(What(cmp(l,r)))
}

func What(n int) string {
	if n == 0 {
		return "same"
	} n > 0 {
		return "l > 0"
	}
}

func main() {
	s := []string{"foo","bar","zz","gg","aa","bb"}
	
	print(Compare2(String, &s[0], &s[1]))
}

