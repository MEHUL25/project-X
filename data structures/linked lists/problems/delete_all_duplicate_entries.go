package problems

// Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

// Use a dummy node before the head to simplify edge cases.

// Traverse with two pointers:

// prev: points to the last node in the result list.

// current: used to scan the list.

// If a duplicate is detected (current.Val == current.Next.Val), skip all nodes with that value.

// If not a duplicate, connect prev.Next to current.
