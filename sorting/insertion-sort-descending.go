// Subprogram 3: descending insertion sort.
package main

// Perform the descending insertion sort.
func descSort(A []int) {
	// Visualize original array.
	visualize(A)

	// Insertion sort.
	for i := 1; i < len(A); i++ {
		// If the left element is smaller than the right element, switch position.
		for j := i; j > 0 && A[j-1] < A[j]; j-- {
			A[j], A[j-1] = A[j-1], A[j]

			// Visualize each step.
			visualize(A)
		}
	}
}
