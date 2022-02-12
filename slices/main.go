package main

import (
	_ "fmt"
	_ "math/rand"
	_ "os"
	_ "sort"
	_ "strconv"
	_ "time"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// 	rand.Seed(time.Now().UnixNano())

	// 	max, _ := strconv.Atoi(os.Args[1])

	// 	var uniques []int

	// loop:
	// 	for len(uniques) < max {
	// 		n := rand.Intn(max) + 1
	// 		fmt.Printf("%d ", n)

	// 		for _, unique := range uniques {
	// 			if unique == n {
	// 				continue loop
	// 			}
	// 		}
	// 		uniques = append(uniques, n)
	// 	}
	// 	fmt.Println("\n\n uniques: ", uniques)

	// 	sort.Ints(uniques)

	// 	fmt.Println("\n\n Sorted: ", uniques)

	//Full slice expressions

	s.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	odds := append(nums[:2:2], 5, 7) //limit the capacity
	//Unless you want to override the old slice,use full slice expression
	evens := append(nums[2:4], 6, 8)
	s.Show("nums", nums)
	s.Show("odds", odds)
	s.Show("evens", evens)

}
