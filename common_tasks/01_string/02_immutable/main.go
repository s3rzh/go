package main

// что выведет программа и как её исправить?

func main() {
	println(f())
}

func f() string {
	s := "Test"
	s[0] = 'R'
	return s
}
