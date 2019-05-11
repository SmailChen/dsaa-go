package main

import "container/list"

func listExchange(L *list.List, pos1 int, pos2 int) bool {
	if L.Len() == 0 || pos1 < 0 || pos2 < 0 || pos1 > L.Len() || pos2 > L.Len() {
		return false
	}
	if pos1 < pos2 {
		p := 0
		var tmpE list.Element

		for e := L.Front(); e != nil; e = e.Next() {
			if p == pos1 {

			}
		}
	}

}

func mian() {

}
