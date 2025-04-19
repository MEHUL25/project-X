package problems

// We can use a prefix sum approach:

// Keep a running sum while iterating through the linked list.

// Use a map to store the latest node seen at each sum value.

// If the same prefix sum appears twice, it means the values in between sum to zero (since the cumulative sum didn't change between those points).

// Remove the nodes in between by updating pointers.
