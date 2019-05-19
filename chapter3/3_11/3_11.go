// 问题：
// 假设一个单向链表的实现有一个表头结点,但是没有尾结点,并且只有一个指向表头结点的指针
// 写一个类,使之包括的方法可以
// a.返回链表的大小。
// b.打印链表。
// c.检測值x是否在链表中。
// d.如果值x没在链表中,则将其加入链表。
// e.如果值x包含在链表中,则删除这个值。

package main

import (
	"dsaa-go/chapter3/slist"
	"fmt"
)

const ListNum int = 5

func main() {
	L := slist.New()
	// L 升序排列
	for i := 0; i < ListNum; i++ {
		L.PushFront(i * 2)
	}

	// a 返回链表的大小
	fmt.Println(L.Len())
	fmt.Println("====")

	//b 打印链表
	L.Print()
	fmt.Println("====")

	//检测值X是否在链表中
	fmt.Println("pos:", L.IsElement(6))
	fmt.Println("====")

	//如果值X没在链表中，则将其加入链表
	L.HasElement(50)
	L.Print()
	fmt.Println("====")

	//如果值x包含在链表中,则删除这个值
	L.DelElement(6)
	L.Print()
}
