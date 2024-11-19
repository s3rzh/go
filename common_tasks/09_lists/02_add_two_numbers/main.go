package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (it *ListNode) String() (result string) {
	for cur := it; cur != nil; cur = cur.Next {
		if cur == cur.Next {
			return "CYCLE DETECTED"
		}
		result += strconv.Itoa(cur.Val) + " "
	}
	return result
}

func main() {

	// Список: 2->4->3 Число:342
	// first := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// // Список: 5->6->4 Число:465
	// second := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// zero := &ListNode{0, nil}
	// Список: 9->9->9->9->9->9->9 Число: 9999999
	// third := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	// Список: 9->9->9->9 Число: 9999
	// fourth := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}

	third := &ListNode{9, &ListNode{9, nil}}
	fourth := &ListNode{9, &ListNode{}}

	// fmt.Println(addTwoNumbers(first, second)) // Список: 7->0->8  Число: 342+465=807
	// fmt.Println(addTwoNumbers(zero, zero))    // 0
	fmt.Println(addTwoNumbers(third, fourth)) // Список: 8->9->9->9->0->0->0->1 Число: 10009998
}

func addTwoNumbers(a, b *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	sum, carry := 0, 0
	for a != nil || b != nil || carry != 0 {
		sum = carry
		if a != nil {
			sum += a.Val
			a = a.Next
		}
		if b != nil {
			sum += b.Val
			b = b.Next
		}
		carry = sum / 10
		res.Next = &ListNode{sum % 10, nil}
		res = res.Next
	}

	return head.Next
}

// Задача
// На вход приходят 2 неотрицательных числа в виде связного списка. Цифры хранятся в обратном порядке. Необходимо сложить 2 числа и вернуть результат в виде связного списка.

// Решение
// В решении мы имитируем сложение числа в столбик. Проходимся по разрядам, складываем числа и, если полученное число больше 9, записываем оставшуюся часть в остаток carry. 
// В цикле проходимся до тех пор пока у нас есть остаток или есть узлы в списках a и b. 
// Результат возвращается в виде head.Next, потому что первый узел фиктивный и в нем хранится zero value - 0, его мы пропускаем.
