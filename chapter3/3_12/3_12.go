// 问题：
// 重复练习3.11 保持单向链表总是处于排序状态

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
		L.PushFrontInc(data[i], utils.IntComparator)
	}

	// a 返回链表的大小
	fmt.Println("len:", L.Len())

	//b 打印链表
	fmt.Printf("Print: ")
	L.Print()

	//c 检测值X是否在链表中
	_, pos := L.IsElement(6, utils.IntComparator)
	fmt.Println("pos:", pos)
	fmt.Println()

	//d 如果值X没在链表中，则将其加入链表
	L.HasElementInc(10, utils.IntComparator)
	fmt.Printf("inc Print: ")
	L.Print()

	//e 如果值x包含在链表中,则删除这个值
	L.DelElement(6, utils.IntComparator)
	fmt.Printf("\ndel Print: ")
	L.Print()
}
