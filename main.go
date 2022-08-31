package main

import (
	"fmt"
	"sort"
)

type Op int

const (
	_ Op = iota
	Right
	Left
)

func BinSearch(arr []int, search_val int) []int {
	var res []int

	low := 0
	hight := len(arr) - 1

	for low < hight {
		m := low + (hight-low)/2
		if m < 1 {
			return res
		}

		if search_val > arr[m] {
			low += 1
		} else if search_val < arr[m] {
			low -= 1
		} else {
			res = append(res, m)
			r := checkN(arr, m, search_val, Right)
			res = append(res, r...)
			l := checkN(arr, m, search_val, Left)
			res = append(l, res...)
			return res
		}
	}

	return res
}

func checkN(arr []int, pos int, search_val int, op Op) []int {
	var buf []int
	hight := len(arr) - 1

	var compare func(l int, r int) bool
	var continPred bool
	var offset func(pos *int)

	if op == Right {
		continPred = pos < hight
		offset = func(pos *int) { *pos += 1 }
		compare = func(l int, r int) bool {
			return l < r
		}
	}

	if op == Left {
		continPred = pos > 0
		offset = func(pos *int) { *pos -= 1 }
		compare = func(l int, r int) bool {
			return l > r
		}
	}

	for continPred {
		offset(&pos)
		if compare(search_val, arr[pos]) {
			return buf
		}

		buf = append(buf, pos)
	}

	return buf
}

func main() {
	data_arr := []int{5, 8, 77, 1, 4, 25, 13, 54, 22}
	sort.Ints(data_arr)

	var search_value int

	fmt.Print("Input search value >>")
	fmt.Scanln(&search_value)

	index := BinSearch(data_arr, search_value)
	if len(index) == 0 {
		fmt.Println("Element is not in the array")
		return
	}

	fmt.Println(data_arr)
	fmt.Println(index)

}
