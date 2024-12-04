package main

func main() {
	var data1 [5]int               // [0 0 0 0 0]
	var data2 [2][5]int            // [[0 0 0 0 0] [0 0 0 0 0]]
	data3 := [...]int{1, 2, 3}     // [1 2 3]
	data4 := [5]int{1, 2, 3}       // [1 2 3 0 0]
	data5 := [5]int{3: 4}          // [0 0 0 4 0]
	data6 := [5]int{2: 5, 6, 1: 7} // [0 7 5 6 0]
}
