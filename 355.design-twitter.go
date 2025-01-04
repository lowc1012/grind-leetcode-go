/*
 * @lc app=leetcode id=355 lang=golang
 *
 * [355] Design Twitter
 */

// @lc code=start

type MaxHeap []Tweet

func (h MaxHeap) Less(i, j int) bool {
	return h[i].time > h[j].time
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Tweet))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type Tweet struct {
	id int
	time int
}

type Twitter struct {
	time int
	tweets map[int][]Tweet
	followees map[int]map[int]bool // true -> followed, false -> unfollowed
}


func Constructor() Twitter {
    return Twitter{
			tweets: make(map[int][]Tweet),
			followees: make(map[int]map[int]bool),
		}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.time++

	// create Tweet
	tweet := Tweet{
		tweetId,
		this.time,
	}
	this.tweets[userId] = append(this.tweets[userId], tweet)
}


func (this *Twitter) GetNewsFeed(userId int) []int {
  // use Heap to get feeds by ordered
	h := &MaxHeap{}
	heap.Init(h)

	// Add user's tweets to maximum heap
	for _, t := range this.tweets[userId] {
		heap.Push(h, t)
	}

	// Add followees' tweets to maximum heap
	for followeeId := range this.followees[userId] {
		for _, t := range this.tweets[followeeId] {
			heap.Push(h, t)
		}
	}

	// retrieve the most recent 10 tweets
	feed := []int{}
	for i := 10; i > 0 && h.Len() > 0; i-- {
		tweet := heap.Pop(h)
		feed = append(feed, tweet.(Tweet).id)
	}

	return feed
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
  // check if is existing in followees
	if this.followees[followerId] == nil {
		this.followees[followerId] = make(map[int]bool)
	}
	this.followees[followerId][followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
  // check if is existing in followees
	if this.followees[followerId] != nil {
		delete(this.followees[followerId], followeeId)
	}
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end

