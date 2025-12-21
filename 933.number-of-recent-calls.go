/*
 * @lc app=leetcode id=933 lang=golang
 *
 * [933] Number of Recent Calls
 */

// @lc code=start
type RecentCounter struct {
    requests []int
}


func Constructor() RecentCounter {
    return RecentCounter{
        requests: []int{},
    }
}


func (this *RecentCounter) Ping(t int) int {
    this.requests = append(this.requests, t)

    startIdx := 0
    for i, time := range this.requests {
        if time >= t-3000 {
            startIdx = i
            break
        }
    }
    this.requests = this.requests[startIdx:]
    return len(this.requests)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

