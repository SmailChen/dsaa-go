// 该包参考和扩展了标准库的的 container/list

package slist

import (
	"fmt"
	"math"
	"reflect"
)

type Element struct {
	next  *Element
	Value interface{}
}

type SList struct {
	root Element
	len  int
}

func (l *SList) Init() *SList {
	l.root.next = nil
	l.len = 0
	return l
}

func New() *SList { return new(SList).Init() }

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if p := e.next; p != nil {
		return p
	}
	return nil
}

// Front returns the first element of list l or nil if the list is empty.
func (l *SList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// insert inserts e after at, increments l.len, and returns e.
func (l *SList) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.next = n
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *SList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *SList) PushFront(v interface{}) *Element {

	return l.insertValue(v, &l.root)
}

func (l *SList) remove(index int) bool {
	if index > l.Len() || index < 1 {
		return false
	}
	n := l.PreElement(index)
	e := n.next
	n.next = e.next
	e.next = nil
	l.len--
	return true
}

// 删除 index位置的节点
func (l *SList) Remove(index int) bool {
	return l.remove(index)
}

func (l *SList) Len() int { return l.len }

// 获取index位置的前驱节点
func (l *SList) preElement(start *Element, startPos int, index int) *Element {
	e := start
	for i := startPos; i < index-1; i++ {
		e = e.Next()
	}
	return e
}

func (l *SList) PreElement(index int) *Element {
	return l.preElement(&l.root, 0, index)
}

func (l *SList) min(p1, p2 int) int {
	if p1 < p2 {
		return p1
	} else {
		return p2
	}
}

func (l *SList) max(p1, p2 int) int {
	if p1 > p2 {
		return p1
	} else {
		return p2
	}
}

// 交互换相邻元素
func (l *SList) ExAdjoinElement(p1, p2 int) *SList {
	if p1 == p2 || math.Abs(float64(p1-p2)) > 1 || p1 == 0 || p2 == 0 || p1 > l.len || p2 > l.len {
		return l
	}
	e := l.PreElement(l.min(p1, p2)) //较小位置的节点的前驱
	n := e.Next()
	n1 := e.Next().Next()
	e.next = n1
	n.next = n1.Next()
	n1.next = n

	return l

}

// 交互换任意元素
func (l *SList) ExchangeElement(p1, p2 int) *SList {

	if math.Abs(float64(p1-p2)) <= 1 || p1 == 0 || p2 == 0 || p1 > l.len || p2 > l.len {
		return l.ExAdjoinElement(p1, p2)
	}
	e := l.PreElement(l.min(p1, p2)) //较小位置的节点的前驱
	n1 := e
	n := n1.Next()
	n2 := n1.Next().Next()

	e = l.preElement(n, l.min(p1, p2), l.max(p1, p2)) //较大位置的节点的前驱
	m1 := e
	m := m1.Next()
	m2 := m1.Next().Next()

	n1.next = m
	n.next = m2
	m1.next = n
	m.next = n2

	return l

}

// 打印链表
func (l *SList) Print() {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (l *SList) rprint(e *Element) {
	if e != nil {
		l.rprint(e.Next())
		fmt.Println(e.Value)
	}
}

// 反转打印链表
func (l *SList) Rprint() {
	l.rprint(l.Front())
}

// 检查值是否在链表
func (l *SList) IsElement(x interface{}) int {
	index := 0
	for e := l.Front(); e != nil; e = e.Next() {
		index++
		if reflect.ValueOf(e.Value).Int() == reflect.ValueOf(x).Int() {
			return index
		}
	}
	return 0
}

// 检查值是否在链表,没有则加入
func (l *SList) HasElement(x interface{}) *SList {

	if l.IsElement(x) == 0 {
		l.PushFront(x)
	}
	return l
}

// 检查值是否在链表,有则删除
func (l *SList) DelElement(x interface{}) bool {
	if index := l.IsElement(x); index > 0 {
		return l.Remove(index)
	}
	return false
}
