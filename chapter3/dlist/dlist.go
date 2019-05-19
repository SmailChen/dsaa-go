// 该包扩展了标准库的的 container/list

package dlist

import "fmt"

// Element is an element of a linked DList.
type Element struct {
	// Next and previous pointers in the doubly-linked DList of elements.
	// To simplify the implementation, internally a DList l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// DList element (l.Back()) and the previous element of the first DList
	// element (l.Front()).
	next, prev *Element

	// The DList to which this element belongs.
	DList *DList

	// The value stored with this element.
	Value interface{}
}

// Next returns the next DList element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.DList != nil && p != &e.DList.root {
		return p
	}
	return nil
}

// Prev returns the previous DList element or nil.
func (e *Element) Prev() *Element {
	if p := e.prev; e.DList != nil && p != &e.DList.root {
		return p
	}
	return nil
}

// DList represents a doubly linked DList.
// The zero value for DList is an empty DList ready to use.
type DList struct {
	root Element // sentinel DList element, only &root, root.prev, and root.next are used
	len  int     // current DList length excluding (this) sentinel element
}

// Init initializes or clears DList l.
func (l *DList) Init() *DList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New returns an initialized DList.
func New() *DList { return new(DList).Init() }

// Len returns the number of elements of DList l.
// The complexity is O(1).
func (l *DList) Len() int { return l.len }

// Front returns the first element of DList l or nil if the DList is empty.
func (l *DList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of DList l or nil if the DList is empty.
func (l *DList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero DList value.
func (l *DList) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *DList) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.DList = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *DList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove removes e from its DList, decrements l.len, and returns e.
func (l *DList) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.DList = nil
	l.len--
	return e
}

// move moves e to next to at and returns e.
func (l *DList) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e

	return e
}

// Remove removes e from l if e is an element of DList l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *DList) Remove(e *Element) interface{} {
	if e.DList == l {
		// if e.DList == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of DList l and returns e.
func (l *DList) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of DList l and returns e.
func (l *DList) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the DList is not modified.
// The mark must not be nil.
func (l *DList) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.DList != l {
		return nil
	}
	// see comment in DList.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the DList is not modified.
// The mark must not be nil.
func (l *DList) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.DList != l {
		return nil
	}
	// see comment in DList.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of DList l.
// If e is not an element of l, the DList is not modified.
// The element must not be nil.
func (l *DList) MoveToFront(e *Element) {
	if e.DList != l || l.root.next == e {
		return
	}
	// see comment in DList.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of DList l.
// If e is not an element of l, the DList is not modified.
// The element must not be nil.
func (l *DList) MoveToBack(e *Element) {
	if e.DList != l || l.root.prev == e {
		return
	}
	// see comment in DList.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the DList is not modified.
// The element and mark must not be nil.
func (l *DList) MoveBefore(e, mark *Element) {
	if e.DList != l || e == mark || mark.DList != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the DList is not modified.
// The element and mark must not be nil.
func (l *DList) MoveAfter(e, mark *Element) {
	if e.DList != l || e == mark || mark.DList != l {
		return
	}
	l.move(e, mark)
}

// PushBackList inserts a copy of an other DList at the back of DList l.
// The lists l and other may be the same. They must not be nil.
func (l *DList) PushBackList(other *DList) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of an other DList at the front of DList l.
// The lists l and other may be the same. They must not be nil.
func (l *DList) PushFrontList(other *DList) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// 打印链表
func (l *DList) Print() {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (l *DList) min(p1, p2 int) int {
	if p1 < p2 {
		return p1
	} else {
		return p2
	}
}

func (l *DList) max(p1, p2 int) int {
	if p1 > p2 {
		return p1
	} else {
		return p2
	}
}

// 交互换相邻元素
func (l *DList) ExAdjoinElement(p1, p2 int) *DList {

	e := &l.root
	for i := 0; i < l.min(p1, p2); i++ {
		e = e.next //较小位置的节点
	}
	n := e.next //较小位置节点的后继
	e.prev.next = n
	n.prev = e.prev

	e.next = n.next
	n.next.prev = e

	n.next = e
	e.prev = n

	return l
}
