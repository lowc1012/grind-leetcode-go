/*
 * @lc app=leetcode id=1396 lang=golang
 *
 * [1396] Design Underground System
 */

// @lc code=start
type UndergroundSystem struct {
	checkIns map[int]checkInRecord
	routes   map[route]routeStats
}

type checkInRecord struct {
	station string
	time    int
}

type route struct {
	from string
	to   string
}

type routeStats struct {
	totalTime int
	count     int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkIns: make(map[int]checkInRecord),
		routes:   make(map[route]routeStats),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	// record check-in event
	this.checkIns[id] = checkInRecord{
		stationName,
		t,
	}
	return
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	// look up latest check-in for id
	c := this.checkIns[id]
	fromStation, startTime := c.station, c.time

	// record route data
	r := route{
		from: fromStation,
		to:   stationName,
	}

	// retrieve and update route stats data
	stats := this.routes[r]
	stats.totalTime += (t - startTime)
	stats.count += 1
	this.routes[r] = stats

	return
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    // retrieve route stats
    r := route{
        startStation,
        endStation,
    }
    stats := this.routes[r]
    return float64(stats.totalTime) / float64(stats.count)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
// @lc code=end

