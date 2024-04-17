package main

import "fmt"

// in place sorting
func cyclicSort(nums []int) []int {
	i := 0

	// technique: check number and check if they are in right index

	for i < len(nums) {
		idxOfCurrentElement := nums[i] - 1 // index of the current element where it should be placed

		if nums[i] != nums[idxOfCurrentElement] { // Check if the current element is not in its correct position.
			nums[i], nums[idxOfCurrentElement] = nums[idxOfCurrentElement], nums[i] // Swap the current element with the one at its correct position.
		} else {
			i++
		}
	}

	return nums
}

func main() {
	// unsorted from 1 to n
	arr := []int{1, 3, 4, 2, 5}

	result := cyclicSort(arr)
	fmt.Println("Sorted array is", result)

}
