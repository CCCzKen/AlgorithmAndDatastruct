package main

import (
	"fmt"
	"time"
)

func BubbleSort(values []int) {
	var tmp int
	length := len(values)

	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if values[j] < values[j-1] {
				tmp = values[j-1]
				values[j-1] = values[j]
				values[j] = tmp
			}
		}
	}
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", values[i])
	}
	fmt.Println()
}

func BubbleSort2(values []int) {
	var tmp int
	length := len(values)

	for i := 0; i < length ; i++ {
		for j := length - 1; j > i; j-- {
			if values[j] < values[j-1] {
				tmp = values[j-1]
				values[j-1] = values[j]
				values[j] = tmp
			}
		}
	}

	for i := 0; i < length; i++ {
		fmt.Printf("%d ", values[i])
	}
	fmt.Println()

}

func BubbleSort3(values []int) {
	var tmp int
	length := len(values)

	for i := 0; i < length; i++ {
		for j := length - 1; j < i; j-- {
			if values[j] < values[j-1] {
				tmp = values[j-1]
				values[j] = values[j-1]
				values[j-1] = tmp
			}
		}
	}
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", values[i])
	}
	fmt.Println()
}

func BubbleSort4(values []int) {
	var tmp int
	length := len(values)

	for i := 0; i < length; i++ {
		for j := length - 1; j < i; j-- {
			if values[j] < values[j-1] {
				tmp = values[j-1]
				values[j-1] = values[j]
				values[j] = tmp
			}
		}
	}
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", values[i])
	}
	fmt.Println()
}

func main() {
	values := []int{513, 517, 520, 9, 11, 530, 20, 533, 24, 538, 27, 28, 176, 549, 552, 41, 554, 555, 558, 48, 562, 563, 52, 54, 57, 59, 575, 64, 779, 68, 75, 590, 601, 92, 605, 94, 608, 99, 612, 614, 102, 104, 225, 189, 627, 629, 633, 635, 126, 640, 107, 132, 645, 648, 139, 141, 144, 536, 660, 366, 964, 666, 674, 170, 171, 685, 688, 182, 183, 184, 185, 187, 701, 191, 194, 198, 204, 716, 718, 723, 725, 215, 732, 224, 737, 226, 228, 741, 231, 750, 754, 249, 762, 42, 896, 259, 260, 774, 267, 386, 271, 277, 284, 803, 807, 808, 299, 303, 817, 478, 311, 318, 319, 736, 834, 836, 325, 839, 328, 334, 340, 341, 345, 350, 866, 356, 873, 362, 875, 876, 878, 883, 884, 374, 378, 380, 384, 898, 390, 396, 397, 398, 399, 914, 403, 917, 411, 927, 928, 931, 420, 935, 424, 430, 431, 950, 439, 442, 960, 961, 452, 453, 456, 969, 262, 975, 977, 978, 472, 985, 989, 990, 992, 995, 579, 502, 504, 505}
	start := time.Now().UnixNano()
	BubbleSort(values)
	BubbleSort2(values)
	BubbleSort3(values)
	BubbleSort4(values)
	fmt.Println(time.Now().UnixNano() - start)
}
