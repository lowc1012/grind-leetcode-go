/*
 * @lc app=leetcode id=1396 lang=golang
 *
 * [1396] Design Underground System
 */

// @lc code=start
type Record struct {
	startStation string
	time         int
}

type UndergroundSystem struct {
	passenger map[int]Record

	// key: pair of (start station, end station)
	// value: [record of travel time, total travel counter]
	trafficRecord map[[2]string][2]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		make(map[int]Record),
		make(map[[2]string][2]int),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.passenger[id] = Record{
		stationName,
		t,
	}
	return
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	// update current traversal into corresponding traffic record
	latestRecord := this.passenger[id]
	startStation, startTime := latestRecord.startStation, latestRecord.time

	endStation, endTime := stationName, t

	trafficKey := [2]string{startStation, endStation}

	// get prev total travel time and prev total traverl counter
	prevTraffic := this.trafficRecord[trafficKey]
	prevTraverlTimeSum, prevTraverlCounter := prevTraffic[0], prevTraffic[1]

	this.trafficRecord[trafficKey] = [2]int{prevTraverlTimeSum + (endTime - startTime), prevTraverlCounter + 1}

	return
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	// get latest traffic info from station pair
	trafficInfo := this.trafficRecord[[2]string{startStation, endStation}]

	totalTraverlTime, totalTraverlCounter := trafficInfo[0], trafficInfo[1]

	return float64(totalTraverlTime) / float64(totalTraverlCounter)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
// @lc code=end

