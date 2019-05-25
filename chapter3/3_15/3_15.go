// 问题：
// 给List类添加 splice操作。其声明如下:
// void splice( iterator position, List<T>& 1st);
// 删別除1st中的所有项,并将这些项放在List*this中的位置 positlon之前。1st和*this必
// 须是不同的表。所写的例程必须是常量时间的。

package main

import (
	"dsaa-go/chapter3/slist"
	"fmt"
)

func main() {
	data1 := []int{8, 7, 6}
	data2 := []int{3, 2, 1}
	L1 := slist.New()
	L2 := slist.New()
	for i := 0; i < len(data1); i++ {
		L1.PushFront(data1[i])
	}

	for i := 0; i < len(data2); i++ {
		L2.PushFront(data2[i])
	}

	//打印链表
	fmt.Printf("L1: ")
	L1.Print()

	fmt.Printf("\nL2: ")
	L2.Print()

	// 合并
	L1.Splice(1, L2)

	//打印链表
	fmt.Printf("\nL1: ")
	L1.Print()

	fmt.Printf("\nL2: ")
	L2.Print()

}
