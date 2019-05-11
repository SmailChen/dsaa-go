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

	l := intersectionList(P, L)

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
