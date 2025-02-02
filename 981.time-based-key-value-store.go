/*
 * @lc app=leetcode id=981 lang=golang
 *
 * [981] Time Based Key-Value Store
 */

// @lc code=start
type TimeVal struct {
	val string
	ts  int
}


type TimeMap struct {
    store map[string][]TimeVal
}


func Constructor() TimeMap {
    return TimeMap{
		store: make(map[string][]TimeVal),
	}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	if _, exist := this.store[key]; !exist {
		this.store[key] = []TimeVal{}
	}
	this.store[key] = append(this.store[key], TimeVal{
		val: value,
		ts: timestamp,
	})
}


func (this *TimeMap) Get(key string, timestamp int) string {
	slice, exist := this.store[key]
	if !exist || len(slice) == 0 {
        return ""
    }

	left, right := 0, len(slice)-1

	// early termination
    if timestamp < slice[0].ts {
        return ""
    } else if timestamp >= slice[right].ts {
        return slice[right].val
    }

	// use binary search to find the 
	for left <= right {
		mid := (left+right) / 2
		if slice[mid].ts == timestamp {
			return slice[mid].val
		} else if slice[mid].ts < timestamp {
			if mid+1 < len(slice) && slice[mid+1].ts > timestamp {
                return slice[mid].val
            }
			left = mid+1
		} else {
			right = mid-1
		}
	}

	if right >= 0 {
		return slice[right].val
	}

	return ""
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end

