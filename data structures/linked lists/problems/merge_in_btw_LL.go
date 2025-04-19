package problems

// You are given two linked lists: list1 and list2 of sizes n and m respectively.

// Remove list1's nodes from the ath node to the bth node, and put list2 in their place.

// Solution :

// Find prevA:
// Node just before position a.

// Find afterB:
// Node just after position b.

// Link prevA.Next to head of list2.

// Traverse to tail of list2.

// Link tail.Next to afterB.

// Return updated list1 head.
