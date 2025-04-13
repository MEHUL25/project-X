package list

type Node struct {
	val  int
	next *Node
}

func simpleNodeExample() {

	var head *Node = nil
	node1 := Node{val: 1}

	node2 := Node{val: 2}
	node1.next = &node2

	node3 := Node{val: 3}
	node2.next = &node3

	node4 := Node{val: 4}
	node3.next = &node4

	node5 := Node{val: 5}
	node4.next = &node5

	head = &node1

	for head != nil {
		println(head.val)
		head = head.next
	}
}
