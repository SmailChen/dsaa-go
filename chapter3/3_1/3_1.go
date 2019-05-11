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
