package lld

import "container/heap"

type Twitter struct {
	time       int
	userTweets map[int][]Tweet          // userId -> list of tweets
	follows    map[int]map[int]struct{} // followerId -> set of followeeIds
}
type Tweet struct {
	tweetId int
	time    int
}

func Constructor() Twitter {
	return Twitter{
		time:       0,
		userTweets: make(map[int][]Tweet),
		follows:    make(map[int]map[int]struct{}),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.userTweets[userId] = append(this.userTweets[userId], Tweet{
		tweetId: tweetId,
		time:    this.time,
	})
	this.time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &TweetHeap{}
	heap.Init(h)

	// Ensure the user follows themselves
	this.Follow(userId, userId)

	for followee := range this.follows[userId] {
		tweets := this.userTweets[followee]
		for i := len(tweets) - 1; i >= 0 && i >= len(tweets)-10; i-- {
			heap.Push(h, tweets[i])
			if h.Len() > 10 {
				heap.Pop(h)
			}
		}
	}

	result := make([]int, h.Len())
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(Tweet).tweetId
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.follows[followerId] == nil {
		this.follows[followerId] = make(map[int]struct{})
	}
	this.follows[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followeeId != followerId { // cannot unfollow self
		delete(this.follows[followerId], followeeId)
	}
}

// Priority Queue to keep most recent tweets (min-heap)
type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TweetHeap) Push(x interface{}) {
	*h = append(*h, x.(Tweet))
}

func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	t := old[n-1]
	*h = old[0 : n-1]
	return t
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
