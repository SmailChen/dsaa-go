// 问题:
// 给定两个排序后的表L1和L2.写出一个程序仅使用基本的表操作来计算L1 ∪ L2。
package main

import (
	"container/list"
	"fmt"
)

const ListNum int = 20

// 时间复杂度O(n)
func unionList(L1 *list.List, L2 *list.List) *list.List {

	if L1.Len() == 0 {
		return L2
	}
	if L2.Len() == 0 {
		return L1
	}

	MaxNum := 0
	if L1.Len() > L2.Len() {
		MaxNum = L1.Len()
	} else {
		MaxNum = L2.Len()
	}
	L := list.New()
	e1 := L1.Front()
	e2 := L2.Front()

	for i := 0; i < MaxNum; i++ {
		// 元素相等，追加到L，然后L1和L2，同时指向下一个元素
		if e1.Value.(int) == e2.Value.(int) {
			L.PushBack(e1.Value)
			e2 = e2.Next()
			e1 = e1.Next()
		}
		// 较小的元素追加到L，同时指向下一个元素
		if e1.Value.(int) < e2.Value.(int) {
			L.PushBack(e1.Value)
			e1 = e1.Next()
		} else if e1.Value.(int) > e2.Value.(int) {
			L.PushBack(e2.Value)
			e2 = e2.Next()
		}

		// 追加较长链表的剩余元素到L
		if e1 == nil || e2 == nil {
			if e1 == nil {
				for ; e2 != nil; e2 = e2.Next() {
					L.PushBack(e2.Value)
				}
			}
			if e2 == nil {
				for ; e1 != nil; e1 = e1.Next() {
					L.PushBack(e1.Value)
				}
			}
			return L
		}

	}
	return nil
}

func main() {
	L1 := list.New()
	L2 := list.New()

	// L1 升序排列
	for i := 0; i < ListNum; i++ {
		L1.PushBack(i * 2)
	}
	// L2 升序排列
	for i := 0; i < ListNum/2; i++ {
		L2.PushBack(i + 2)
	}

	l := unionList(L1, L2)

	// 打印结果
	for e := L1.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	for e := L2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	for e := l.Front(); e != nil; e = e.Next() {
		if e == nil {
			break
		}
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

}
