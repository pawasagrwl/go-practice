// Union of Two Sorted Arrays
// Accuracy: 13.34% Submissions: 15 Points: 80
// Given two sorted arrays. The task is to find the union of these two arrays.
// Note: Union of two arrays can be defined as the distinct elements in the two arrays
// Example 1:
// Input:
// N = 5, M = 3
// arr1[1 = {1 2 3 4 5}
// arr2[1 = {1 2 3}
// Output: 1 2 3 4 5
// Explanation: Union of the arrays is:
// {1 2 3 4 5}
// â€ Example 2:
// Input:
// N = 5, M = 5
// arr1[1 = {2 2 3 4 5}
// arr2[] = {1 1 2 3 4)
// Output: 2 3 4 5 1
// Explanation: The Union of the arrays is:
// {2 3 4 5 1}
// Your Task:
// You don't need to read input or print anything. Your task is to complete the
// function findUnion () which takes two arrays arr1 [I, arr2[] and their sizes n and m
// (respectively) as inputs and returns an array containing the union of the two
// arrays. Maintain the order of elements as they appear in the array. Elements of first array
// should get added first then elements of second array.
// Expected Time Complexity: 0(n + m).
// Expected Auxiliary Space: O(n + m).
// Constraints:
// 1 <=N. M <=105
// 1 <= arrlil. brrlil <= 106

package main

import "fmt"


func findUnion (arr1 []int, arr2 []int) []int {
	elems := make(map[int]bool)
	var arr []int
	for _, elem := range arr1 {
		if !elems[elem] {
			arr = append(arr, elem)
			elems[elem] = true;
		}
	}

	for _, elem := range arr2 {
		if !elems[elem] {
			arr = append(arr, elem)
			elems[elem] = true;
		}
	}
	return arr
}
func main () {
	arr1 := []int{1,2,3,4,5}
	arr2 := []int{4,5,5,7,8}
	fmt.Println(findUnion(arr1, arr2))
}