package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 2, 3, 4}
	//1) создать отдельный список в котором будут содержать эти элементы но по одному
	unique := []int{}
	for i := 0; i < len(arr); i++ {
		exElement := arr[i]
		isExist := false

		for j := 0; j < len(unique); j++ {
			if exElement == unique[j] {
				isExist = true
			}
		}

		if !isExist {
			unique = append(unique, exElement)
		}
	}
	for i := 0; i < len(unique); i++ {
		counter := 0
		for j := 0; j < len(arr); j++ {
			if unique[i] == arr[j] {
				counter += 1
			}
		}
		fmt.Println(unique[i], counter)
	}
}
