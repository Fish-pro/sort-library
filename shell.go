/*
Copyright 2021
author: York

希尔排序
原理：
时间复杂度：
空间复杂度：
*/

package sort_go

func ShellSort(array []int) []int {
	for h := len(array) / 2; h > 0; h = h / 2 {
		for i := h; i < len(array); i++ {
			tmp := array[i]
			var j int
			for j = i - h; j >= 0; j -= h {
				if tmp < array[j] {
					array[j+h] = array[j]
				} else {
					break
				}
			}
			array[j+h] = tmp
		}
	}
	return array
}
