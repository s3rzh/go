package main

func main() {
	x := 10
	defer println(x)
	x = 20
	defer println(x)
}

// output:
// 20 
// 10

// те в данном случае х запоминается на момент определения defer (а не на момент вызова!)
