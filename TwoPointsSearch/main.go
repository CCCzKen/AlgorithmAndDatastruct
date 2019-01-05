package main

import "fmt"

func find(_list []int, value int) int {
	low := 0
	high := len(_list)

	for {
		if low > high {
			break
		}
		mid := (high + low) / 2
		if _list[mid] == value {
			return mid
		} else if _list[mid] < value{
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	_list := []int{3, 10, 55, 57, 60, 68, 99, 102, 103, 107}
	fmt.Println(find(_list, 107))
	
}
