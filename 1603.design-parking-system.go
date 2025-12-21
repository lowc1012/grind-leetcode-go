/*
 * @lc app=leetcode id=1603 lang=golang
 *
 * [1603] Design Parking System
 */

// @lc code=start
type ParkingSystem struct {
    slots [3]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    ps := ParkingSystem{
        slots: [3]int{big, medium, small},
    }
    return ps
}


func (this *ParkingSystem) AddCar(carType int) bool {
    // carType can be of three kinds: big, medium, or small, which are represented by 1, 2, and 3
    num := this.slots[carType-1]
    if num > 0 {
        this.slots[carType-1] -= 1
        return true
    }
    return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// @lc code=end

