// 问题：
// 编写仅用一个数组而实现两个栈的例程。除非数组的每一个单元都被使用,否则栈例程不能有溢H声明。

package main

import (
	"errors"
	"fmt"
)

const Max = 3

type SharedStack struct {
	top1  int
	top2  int
	space [Max]interface{}
}

func (s *SharedStack) Init() {
	s.top1 = -1
	s.top2 = len(s.space)
}

// 入栈
func (s *SharedStack) Push(v interface{}, stack int) error {
	if stack != 1 && stack != 2 {
		return errors.New(`parameter "stack" is 1 or 2`)
	}
	if s.top1 >= s.top2-1 {
		return errors.New("SharedStack is no space!")
	}

	if stack == 1 {
		s.top1++
		s.space[s.top1] = v

	} else {
		s.top2--
		s.space[s.top2] = v
	}
	return nil
}

// 出栈
func (s *SharedStack) Pop(i int) (interface{}, error) {
	if i != 1 && i != 2 {
		return nil, errors.New(`parameter "stack" is 1 or 2`)
	}

	var ret interface{}
	if i == 1 {
		if s.top1 == -1 {
			return nil, errors.New("stack1 is empty")
		}
		ret = s.space[s.top1]
		s.space[s.top1] = nil
		s.top1--
	} else {
		if s.top2 == len(s.space) {
			return nil, errors.New("stack2 is empty")
		}
		ret = s.space[s.top2]
		s.space[s.top2] = nil
		s.top2++
	}
	return ret, nil
}

func main() {
	var s SharedStack
	s.Init()

	s.Push(1, 1)
	s.Push(100, 2)
	s.Push(100, 2)
	fmt.Println(s.Push(100, 2))
	for i := 0; i < len(s.space); i++ {
		fmt.Println(s.space[i])
	}
}
