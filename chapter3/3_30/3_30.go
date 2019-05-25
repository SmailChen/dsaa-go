// 问题：
// a.写出一个自调整表的数组实现。在自调整表( self-adjusting list)中,所有的插入操作都发生
// 在表的前端。自调整表添加了一个find操作,当一个元素由圧ind访问的时候,该元素就被
// 移到表的前端,而其他元素的相对顺序保持不变
// b.写出“-个自调整表的链表实现。
// c.假设每一个元素被访问的概率P都是固定的。说明最高访问概率的元素应该在表的最前面。

package main

import (
	"dsaa-go/chapter3/slist"
	"dsaa-go/chapter3/utils"
	"fmt"
)

func main() {
	data := []int{3, 2, 5, 6, 0, 7, 9, 1, 23, 2, 3}
	L := slist.New()
	for i := 0; i < len(data); i++ {
		L.PushFront(data[i])
	}
	L.Print()
	err := L.Find(7, utils.IntComparator)
	fmt.Println(err)
	L.Print()

	err = L.Find(23, utils.IntComparator)
	fmt.Println(err)
	L.Print()

}
