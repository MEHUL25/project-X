package problems

// Find the next smaller element for each element in an array. For each element, you need to find the next element to the right that is smaller than the current element.

// Let's look at an example to understand it better. Consider an array A of integers:

// A = [4, 2, 14, 7, 1, 9]

// For each element, we want to find the immediate next smallest element. If there is no smaller element, we can use some kind of sentinel value, like -1 or None, to indicate that. Here's how we would solve this for array A:

// The next smallest element after 4 is 2.
// The next smallest element after 2 is 1.
// For 14, the next smallest element is 7.
// For 7, the next smallest element is 1.
// For 1, there is no smaller element, so we might say -1 or None.
// Likewise, for 9, the next smaller element is also -1 or None because there is no element after it, and thus no smaller element.
// Thus, after solving the problem for array A, we'd have the result:

// [2, 1, 7, 1, -1, -1]

// Task
// Given an integer n, the length of an array
// In the next line you are given n integers, elements of the array.
// Print the next smallest value for each element of the array. If there does not exist a next smaller element, print -1.

// ===============================================================

// Approach :
// We iterate from the end of the array.

// The stack maintains the potential candidates for the next smaller element.

// For each element, we discard larger or equal elements from the stack (since they can't be the next smaller).

// The next element on the top of the stack (if any) is the answer.
