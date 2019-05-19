// 问题：
// 通过只调整链(而不是数据)来交换两个相邻的元素,使用:
// a.单向链表。
// b.双向链表

package main

import (
	"dsaa-go/chapter3/slist"
	"fmt"
)

const ListNum int = 10

func main() {
	L := slist.New()
	// L 升序排列
	for i := 0; i < ListNum; i++ {
		L.PushFront(i * 2)
	}

	L.Print()
	//L = L.ExchangeElement(3, 6)
	L = L.ExAdjoinElement(3, 4)
	fmt.Println("====")
	L.Print()

}
