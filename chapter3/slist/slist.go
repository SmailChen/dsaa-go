// 该包参考和扩展了标准库的的 container/list

package slist

import (
	"dsaa-go/chapter3/utils"
	"errors"
	"fmt"
	"math"
)

type Element struct {
	next  *Element
	Value interface{}
}

type SList struct {
	root Element
	last *Element
	len  int
}

func (l *SList) Init() *SList {
	l.root.next = nil
	l.last = &l.root
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
	if l.root.next == nil {
		l.last = e
	}
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

// 头升序插入
func (l *SList) PushFrontInc(v interface{}, comparator utils.Comparator) *Element {
	q := &l.root //e的前驱节点
	for e := l.Front(); e != nil; q, e = e, e.Next() {
		if comparator(e.Value, v) == 1 {
			break
		}
	}
	return l.insertValue(v, q)

}

// 头降序插入
func (l *SList) PushFrontDec(v interface{}, comparator utils.Comparator) *Element {
	q := &l.root //e的前驱节点
	for e := l.Front(); e != nil; q, e = e, e.Next() {
		if comparator(e.Value, v) == -1 {
			break
		}
	}
	return l.insertValue(v, q)

}

func (l *SList) remove(index int) bool {
	if index > l.Len() || index < 1 {
		return false
	}
	n := l.PreElement(index)
	e := n.next
	n.next = e.next
	e.next = nil

	if index == l.Len() {
		l.last = n
	}
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

// 交互换任意元素 通过值
func (l *SList) Swap(p1, p2 int) *SList {
	var e1, e2 *Element
	for e, currentElement := 0, l.Front(); e1 == nil || e2 == nil; e, currentElement = e+1, currentElement.Next() {
		switch e {
		case p1:
			e1 = currentElement
		case p2:
			e2 = currentElement
		}
	}
	e1.Value, e2.Value = e2.Value, e1.Value
}

// 打印链表
func (l *SList) Print() {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

func (l *SList) rprint(e *Element) {
	if e != nil {
		l.rprint(e.Next())
		fmt.Printf("%v ", e.Value)
	}
}

// 反转打印链表
func (l *SList) Rprint() {
	l.rprint(l.Front())
}

// 检查值是否在链表
func (l *SList) IsElement(x interface{}, comparator utils.Comparator) (*Element, int) {
	index := 0
	for e := l.Front(); e != nil; e = e.Next() {
		index++
		if comparator(x, e.Value) == 0 {
			return e, index
		}
	}
	return nil, 0
}

// 检查值是否在链表,没有则加入
func (l *SList) HasElement(x interface{}, comparator utils.Comparator) *SList {

	if _, i := l.IsElement(x, comparator); i == 0 {
		l.PushFront(x)
	}
	return l
}

//检查值是否在链表,没有则加入 递增
func (l *SList) HasElementInc(x interface{}, comparator utils.Comparator) *SList {

	if _, i := l.IsElement(x, comparator); i == 0 {
		l.PushFrontInc(x, comparator)
	}
	return l
}

//检查值是否在链表,没有则加入 递减
func (l *SList) HasElementDec(x interface{}, comparator utils.Comparator) *SList {

	if _, i := l.IsElement(x, comparator); i == 0 {
		l.PushFrontDec(x, comparator)
	}
	return l
}

// 检查值是否在链表,有则删除
func (l *SList) DelElement(x interface{}, comparator utils.Comparator) bool {
	if _, index := l.IsElement(x, comparator); index > 0 {
		return l.Remove(index)
	}
	return false
}

// 返回链表最后一个元素
func (l *SList) LastElement() *Element {
	return l.last
}

// Check that the index is within bounds of the list
func (l *SList) withinRange(index int) bool {
	return index >= 0 && index < l.len
}

// 合并链表
func (l *SList) Splice(pos int, lst *SList) {
	// in reverse to keep passed order i.e. L1["a","b"], L2["c","d"] -> L2.Splice(1,L1) -> ["a","b","c",d"]
	if !l.withinRange(pos) || lst.Front() == nil {
		return
	}

	q := l.PreElement(pos)
	lst.last.next = q.Next()
	q.next = lst.Front()
	lst.Init()
}

// 自调整表
func (l *SList) Find(x interface{}, comparator utils.Comparator) error {
	q := &l.root //e的前驱节点
	for e := l.Front(); e != nil; q, e = e, e.Next() {
		if comparator(x, e.Value) == 0 {
			q.next = e.Next()
			e.next = l.root.Next()
			l.root.next = e
			return nil
		}
	}
	return errors.New("No find")
}
