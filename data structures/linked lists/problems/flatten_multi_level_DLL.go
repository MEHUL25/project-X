package problems

// Use a stack to simulate the recursion.

// Traverse the list, and whenever you encounter a node with a Child, do the following:

// Push the current node's Next onto the stack (if it exists).

// Link the current node's Next to the Child.

// Set the child's Prev to current.

// Set the Child to nil (as it's now part of the main list).

// After finishing the child chain, pop the saved Next node from the stack and attach it to the current node.
