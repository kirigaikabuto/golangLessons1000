package main

import "fmt"

//https://acmp.ru/index.asp?main=task&id_task=642

func main() {
	n := 6
	limit := 18
	cars := []int{5, 10, 1, 2, 1, 20}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if cars[i] < cars[j] {
				cars[i], cars[j] = cars[j], cars[i]
			}
		}
	}
	sumi := 0
	counter := 0
	for i := 0; i < n; i++ {
		if sumi+cars[i] >= limit {
			break
		}
		sumi = sumi + cars[i]
		counter += 1
	}
	fmt.Println(counter)
}
