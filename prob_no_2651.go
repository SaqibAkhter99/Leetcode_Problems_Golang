//Problem 2651: Calculate Delayed Arrival Time
// Statement: Given the arrival time and the delayed time, find the delayed arrival time.
// Example 1:
// Input: arrivalTime = 11, delayedTime = 3
// Output: 14
// Explanation: The delayed arrival time is 14.
// Example 2:
// Input: arrivalTime = 15, delayedTime = 2
// Output: 17
//Problem Link: https://leetcode.com/problems/calculate-delayed-arrival-time/




func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
    add:=arrivalTime + delayedTime
    if add%24==0{
        return 0
    }
    if add > 24{
        val:=add - 24
        return val
    }
    return add
    
}