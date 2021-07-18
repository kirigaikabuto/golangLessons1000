package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	r := 4
	//c := 4
	//1 2 3 4
	//5 6 7 8
	//9 10 11 12
	//13 14 15 16
	//17 18 19 20
	global := [][]int{}
	temp := []int{}
	for i := 0; i < len(arr); i++ {
		if (i)%r == 0 && i != 0 {
			temp = append(temp, arr[i])
			global = append(global, temp)
			temp = []int{}
		} else {
			temp = append(temp, arr[i])
		}
	}
	for i := 0; i < len(global); i++ {
		fmt.Println(global[i])
	}
}
