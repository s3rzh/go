package main

func main() {
	x := 1
  
	defer func() {
		x += 1
		println("a:", x)
	}()

	x = 10

	defer func() {
		x += 1
		println("b:", x)
	}()
}

// output:

// b: 11
// a: 12

