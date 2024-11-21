package main



func main() {
	v := 5
	p := &v
	println(*p)

	changePointer(p)
	println(*p)
}

func changePointer(p *int) {
	v := 3
	*p = v // разыменовываем и копируем
}

// Output: 5 3
