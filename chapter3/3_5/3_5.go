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

	l := unionList(P, L)

	// 打印结果
	for e := L.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	for e := P.Front(); e != nil; e = e.Next() {
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
