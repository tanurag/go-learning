package main

import "fmt"
import "../algo"

func main() {
    array := []int{1,6,2,5,8,7,3,4}
    fmt.Println("Unsorted Array",array)
    algo.SelectionSort(array)
    fmt.Println("Sorted Array ",array)
}
