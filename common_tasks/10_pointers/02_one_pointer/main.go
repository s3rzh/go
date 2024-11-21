package main

// Что выведет код?

func main() {
	v := 5
	p := &v
	println(*p)

	changePointer(p)
	println(*p)
}

func changePointer(p *int) {
	v := 3
	p = &v
}

// Output: 5 5
