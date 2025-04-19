package problems

// Given a binary tree

// struct Node {
//   int val;
//   Node *left;
//   Node *right;
//   Node *next;
// }
// Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

// Initially, all next pointers are set to NULL.

// Use a queue for level order traversal.

// At each level:

// Use a prev pointer to track the previous node.

// Connect prev.Next = current.

// Add the current node’s children to the queue.

// After finishing each level, the last node's .Next remains nil by default.
