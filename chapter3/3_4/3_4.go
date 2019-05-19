// 问题:
// 给定两个排序后的表L1和L2.写出一个程序仅使用基本的表操作来计算L1 ∩ L2。
package main

import (
	"container/list"
	"fmt"
)

const ListNum int = 20

// 时间复杂度O(n)
func intersectionList(L1 *list.List, L2 *list.List) *list.List {
	if L1.Len() == 0 || L2.Len() == 0 {
		return nil
	}
	L := list.New()
	e1 := L1.Front()
	e2 := L2.Front()
	for {

		if e1.Value.(int) == e2.Value.(int) {
			L.PushBack(e1.Value)
			e2 = e2.Next()
			e1 = e1.Next()
			if e2 == nil || e1 == nil {
				return L
			}
			continue
		}

		if e1.Value.(int) < e2.Value.(int) {
			e1 = e1.Next()
		} else {
			e2 = e2.Next()
		}

		if e1 == nil || e2 == nil {
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

	l := intersectionList(L1, L2)

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
