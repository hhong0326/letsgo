package main

import "fmt"

func main() {
	// 한정된 범위일 때 가장 효율높은 알고리즘: Counting Sort
	// limit reach: 0 ~ 10
	arr := []int{2, 4, 3, 6, 8, 7, 5, 1, 3, 2, 6, 8, 2, 6, 6, 9, 0, 5, 1, 10}

	var counts [11]int

	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}

	// sorted := make([]int, 0, len(arr))

	// for 문 2개 사용 방법

	// for i := 1; i < len(counts); i++ {
	// 	for j := 0; j < counts[i]; j++ {
	// 		sorted = append(sorted, i)
	// 	}
	// }

	// 누적합을 이용한 방법

	sorted := make([]int, len(arr))

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	for i := 0; i < len(arr); i++ {
		sorted[counts[arr[i]]-1] = arr[i]
		counts[arr[i]]--
	}

	// 결과 출력
	fmt.Println(sorted)
}
