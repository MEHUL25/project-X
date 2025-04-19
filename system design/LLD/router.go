package lld

import "sort"

type Router struct {
	memoryLimit int
	queue       []Packet
	packetSet   map[[3]int]struct{}
	destMap     map[int][]int // map from destination to sorted timestamps
}
type Packet struct {
	source      int
	destination int
	timestamp   int
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit: memoryLimit,
		queue:       make([]Packet, 0),
		packetSet:   make(map[[3]int]struct{}),
		destMap:     make(map[int][]int),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	key := [3]int{source, destination, timestamp}
	if _, exists := this.packetSet[key]; exists {
		return false
	}

	// If memory limit reached, remove the oldest packet
	if len(this.queue) >= this.memoryLimit {
		oldest := this.queue[0]
		this.queue = this.queue[1:]
		oldestKey := [3]int{oldest.source, oldest.destination, oldest.timestamp}
		delete(this.packetSet, oldestKey)

		// Remove timestamp from destination map
		times := this.destMap[oldest.destination]
		for i, t := range times {
			if t == oldest.timestamp {
				this.destMap[oldest.destination] = append(times[:i], times[i+1:]...)
				break
			}
		}
	}

	// Add new packet
	packet := Packet{source, destination, timestamp}
	this.queue = append(this.queue, packet)
	this.packetSet[key] = struct{}{}
	this.destMap[destination] = append(this.destMap[destination], timestamp)

	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.queue) == 0 {
		return []int{}
	}

	packet := this.queue[0]
	this.queue = this.queue[1:]
	key := [3]int{packet.source, packet.destination, packet.timestamp}
	delete(this.packetSet, key)

	// Remove timestamp from destination map
	times := this.destMap[packet.destination]
	for i, t := range times {
		if t == packet.timestamp {
			this.destMap[packet.destination] = append(times[:i], times[i+1:]...)
			break
		}
	}

	return []int{packet.source, packet.destination, packet.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	tList := this.destMap[destination]
	if len(tList) == 0 {
		return 0
	}

	start := sort.Search(len(tList), func(i int) bool {
		return tList[i] >= startTime
	})

	end := sort.Search(len(tList), func(i int) bool {
		return tList[i] > endTime
	})

	return end - start
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
