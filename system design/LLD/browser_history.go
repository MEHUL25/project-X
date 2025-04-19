package problems

type Node struct {
	url  string
	prev *Node
	next *Node
}

type BrowserHistory struct {
	current *Node
}

func Constructor(homepage string) BrowserHistory {
	node := &Node{url: homepage}
	return BrowserHistory{current: node}
}

func (this *BrowserHistory) Visit(url string) {
	newNode := &Node{url: url}
	this.current.next = newNode
	newNode.prev = this.current
	this.current = newNode
}

func (this *BrowserHistory) Back(steps int) string {
	for steps > 0 && this.current.prev != nil {
		this.current = this.current.prev
		steps--
	}
	return this.current.url
}

func (this *BrowserHistory) Forward(steps int) string {
	for steps > 0 && this.current.next != nil {
		this.current = this.current.next
		steps--
	}
	return this.current.url
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
