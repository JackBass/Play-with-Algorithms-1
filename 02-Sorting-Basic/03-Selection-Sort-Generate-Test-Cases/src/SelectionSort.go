package main

import (
	"fmt"
	"practice/Play-with-Algorithms/02-Sorting-Basic/03-Selection-Sort-Generate-Test-Cases/src/helper"
)

func SelectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		// 寻找 [i, n) 区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

// slice 是引用类型
func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func main() {
	// 测试排序算法辅助函数
	N := 20000
	arr := helper.GenerateRandomArray(20000, 0, N)
	SelectionSort(arr, len(arr))

	fmt.Println(arr)
}
