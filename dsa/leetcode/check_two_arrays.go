// Given two arrays A and B of equal size N, the task is to find if given arrays are equal or not.
// Two arrays are said to be equal if both of them contain same set of elements, arrangements
// (or permutation) of elements may be different though.
// Note: If there are repetitions, then counts of repeated elements must also be same for two
// array to be equal.
// Example 1:
// Input:
// N = 5
// A[] = {1,2,5,4, 0)
// B[] = {2,4,5,0,1}
// Output: 1
// Explanation: Both the array can be
// rearranged to {0,1,2,4, 5}
// Example 2:
// Input:
// N = 3
// A[] = {1,2,5}
// B[J = {2,4, 15}
// Output: 0
// Explanation: A[J and B[] have only
// one common value.
// Your Task:
// Complete check() function which takes both the given array and their size as function
// arguments and returns true if the arrays are equal else returns false. The 0 and 1 printing is
// done by the driver code.
// Expected Time Complexity: O(N)
// Expected Auxilliary Space: O(N)
// Constraints:
// 1<=N<=107
// 1<-AM.B[1<=1018

package main

import "fmt"

func check(A []int, B []int) bool {
	if (len(A) != len(B)){
		return false;
	}

	elements := make(map[int]int)


	for _, elem := range A {
		if _, key := elements[elem]; !key {
			elements[elem] = 1;
		} else {
			elements[elem] += 1;
		}
	}

	for _, elem := range B {
		if _, key := elements[elem]; !key {
			return false;
		} else {
			elements[elem] -= 1
		}
	}

	for _, value := range elements {
		if value != 0 {
			return false
		}

	}

	return true

}

func main () {
	A := []int{2}
	B := []int{2}
	fmt.Println(check(A, B))
}