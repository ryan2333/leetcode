package main

import "fmt"

func rotate1(matrix [][]int) {
	var (
		tmp [][]int
	)
	tmp = make([][]int, len(matrix))
	copy(tmp, matrix)
	for i := 0; i < len(tmp); i++ {
		el := []int{}
		for j := len(tmp[i]) - 1; j >= 0; j-- {
			el = append(el, tmp[j][i])
		}
		fmt.Println("el: ", el)
		matrix[i] = el
	}
	fmt.Println("matrix: ", matrix)
	//fmt.Printf("%p %p\n", tmp, matrix)
}

func rotate(matrix [][]int) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := 0; j < size/2; j++ {
			matrix[i][j], matrix[i][size-1-j] = matrix[i][size-1-j], matrix[i][j]
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			matrix[i][j], matrix[size-1-j][size-1-i] = matrix[size-1-j][size-1-i], matrix[i][j]
		}
	}

	fmt.Println("m: ", matrix)
}

func main() {
	var n1, n2 [][]int
	n1 = [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	n2 = [][]int{[]int{5, 1, 9, 11}, []int{2, 4, 8, 10}, []int{13, 3, 6, 7}, []int{15, 14, 12, 16}}
	rotate(n1)
	rotate(n2)
}
