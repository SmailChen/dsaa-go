// 问题：
// Josephus问题( Josephus problem)是下面的游戏:N个人编号1~N,围坐成一个圆
// 圈。从1号人开始传递一个热土豆。经过M次传递后拿着热土豆的人被清除离座,
// 围坐的圆圈缩紧,由坐在被清除人后面的人拿起热土豆继续进行游戏。最后剩下的
// 人获胜。因此,如果M=0和N=5,则游戏人依序被清除,5号游戏人获胜。如果
// M=1和N=5,那么被清除的人的顺序是2,4,1,5。
// a.编写一个程序解决在M和N为一般的值下的 Josephus问题,应使所编程序尽可
// 能地高效率,要确保各个单元都能被清除。
// b.这个程序的运行时间是多少?
// c.如果M=1,这个程序的运行时间又是多少?对于N的一些大值(N>100000)
// delete例程如何影响该程序的速度?

package main

import (
	"container/ring"
	"fmt"
)

// 时间复杂度O(n)
func josephus(r *ring.Ring, M int, N int) int {

	for r.Len() > 1 {
		r = r.Move(M - 1)
		r1 := r.Unlink(1)
		fmt.Println(r1.Value)
		r = r.Move(1)
	}
	return r.Value.(int)
}

func main() {
	r := ring.New(5)
	// 初始化 环形链表
	for i := 1; i <= r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	winner := josephus(r, 1, r.Len())
	fmt.Println("winner:", winner)
}
