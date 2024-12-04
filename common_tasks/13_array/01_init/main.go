package main

func main() {
	var data0 [0]int               // [] неприменимая кострукция, но го позволяет так сделать нуелвые массивы (без элементов)
 	var data1 [5]int               // [0 0 0 0 0]
	var data2 [2][5]int            // [[0 0 0 0 0] [0 0 0 0 0]]
	data3 := [...]int{1, 2, 3}     // [1 2 3]
	data4 := [5]int{1, 2, 3}       // [1 2 3 0 0]
	data5 := [5]int{3: 4}          // [0 0 0 4 0]
	data6 := [5]int{2: 5, 6, 1: 7} // [0 7 5 6 0]
	data7 := [5]int{}              // [0 0 0 0 0] // даже с пустым литералом {} все равно массив будет проинициализирован zero-value
}

// Массив это стркура данных у которой есть адресс начала, длина и кол-во элементов (те это эменты, который обьеденились в одну область памяти, есть начало, размер элемента и длина)

// При передачи массива в функции - они копируются (те если это долшой массив - то будет создан такой же массив в памяти)

// Операции с массивами: чтение, запись, len(arr), cap(arr) (да, можно и она всегда будет равна len!), &arr (те получения указателя на массив) и arr[1:4] получения среза из массива
// результат len(arr) вычисляется на этапе компиляции, те в коде поставляется константа (значение берется из таблицы символов - не нашёл про это ничего!!!)

// Массивы можно сравнивать между собой операторами == и != (не нельзя другиеми, например <, > <=, >= будет ошибка компиляции)

// Размер массива можно указать через константу (через переменную НЕЛЬЗЯ) const N = 5; var arr [N]int

// Обращение к любому элементу массива - это константная величина O(1)

// В момент роста стека горутины (начинает от 2кб) массив может переехать (те скопироваться в другую область памяти)

// Го алоцирует массив в стеке ЕСЛИ его размер меньше и равен (<=) 10 Мб, если больше - будет на хипе, так же на хипе будет если на массив кто-то ссылается за пределами функции.
