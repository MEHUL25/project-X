package lld

// Design a queue that supports push and pop operations in the front, middle, and back.

// Solution :

// We can split the queue into two deques:

// left: holds the first half

// right: holds the second half

// We balance them so that:

// left has equal or 1 more element than right
// i.e., len(left) >= len(right) and len(left) - len(right) ≤ 1

// This allows middle operations to target the end of left.
