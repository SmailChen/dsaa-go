// 问题：
// 编写一个程序计算后缀表达式的值
package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func suffixOperation(str []string) float64 {
	L := list.New()
	var res float64

	for i := 0; i < len(str); i++ {
		if str[i] != "+" && str[i] != "-" && str[i] != "*" && str[i] != "/" {
			L.PushFront(str[i])
		} else {
			op1 := L.Front()
			op2 := L.Front().Next()
			L.Remove(op1)
			L.Remove(op2)

			switch str[i] {
			case "+":
				p1, err := strconv.ParseFloat(op1.Value.(string), 64)
				if err != nil {
					panic(err)
				}
				p2, err := strconv.ParseFloat(op2.Value.(string), 64)
				if err != nil {
					panic(err)
				}
				res = p1 + p2
			case "-":
				p1, err := strconv.ParseFloat(op1.Value.(string), 64)
				if err != nil {
					panic(err)
				}
				p2, err := strconv.ParseFloat(op2.Value.(string), 64)
				if err != nil {
					panic(err)
				}
				res = p1 - p2
			case "*":
				p1, err := strconv.ParseFloat(op1.Value.(string), 64)
				if err != nil {
					panic(err)
				}
				p2, err := strconv.ParseFloat(op2.Value.(string), 64)
				if err != nil {
					panic(err)
				}
				res = p1 * p2
			case "/":
				p1, err := strconv.ParseFloat(op1.Value.(string), 64)
				if err != nil {
					panic(err)
				}
				p2, err := strconv.ParseFloat(op2.Value.(string), 64)
				if err != nil {
					panic(err)
				}
				res = p1 / p2
			}
			L.PushFront(strconv.FormatFloat(res, 'f', 2, 32))
		}
	}
	res, err := strconv.ParseFloat(L.Front().Value.(string), 64)
	if err != nil {
		panic(err)
	}
	return res

}

func main() {

	//str := []string{"6", "5", "2", "3", "+", "8", "*", "+", "3", "+", "*"}
	str := []string{"4.99", "1.06", "*", "5.99", "+", "6.99", "1.06", "*", "+"}
	v := suffixOperation(str)
	fmt.Println(v)
}
