// 问题：
// 给定一个链表L和另一个链表P,它们包含以升序排列的整数。操作 printouts(L,P)将打印L
// 中那些由P所指定的位置上的元素。例如,如果P=1,3,4,6,那么L中的第1、3、4和6个
// 元素被打印出来。写出过程 printlots(L,P)。只可以使用公有的STL容器操作。该过程的运行
// 时间是多少?

package main

import (
	"container/list"
	"fmt"
	"time"
)

const ListNum int = 1e5

// 时间复杂度O(N)
func printLots(L *list.List, P *list.List) {

	var pos int = 0
	j := L.Front()
	start := time.Now()
	for i := P.Front(); i != nil; i = i.Next() {
		for ; pos <= i.Value.(int); j = j.Next() {
			if j == nil {
				return
			}
			if pos == i.Value.(int) {
				fmt.Println(j.Value)
			}
			pos++
		}
	}
	fmt.Println(time.Since(start))

}

/*
// 时间复杂度O(N^2)
func printLots(L *list.List, P *list.List) {

	start := time.Now()
	for i := P.Front(); i != nil; i = i.Next() {
		var pos int = 0
		for j := L.Front(); j != nil; j = j.Next() {
			if pos == i.Value.(int) {
				fmt.Println(j.Value)
			}
			pos++
		}
	}
	fmt.Println(time.Since(start))

}
*/
func main() {

	L := list.New()
	P := list.New()

	// L 升序排列
	for i := 0; i < ListNum; i++ {
		L.PushBack(i * 2)
	}
	// P 升序排列
	for i := 0; i < ListNum/2; i++ {
		P.PushBack(i + 2)
	}
	printLots(L, P)
}
