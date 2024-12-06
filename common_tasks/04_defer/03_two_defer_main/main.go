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
