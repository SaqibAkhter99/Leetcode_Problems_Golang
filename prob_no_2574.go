//Problem Statement: Left and Right Difference
/* Given a 0-indexed integer array nums, find a 0-indexed integer array answer where:
answer.length == nums.length.
answer[i] = |leftSum[i] - rightSum[i]|.
Where:
leftSum[i] is the sum of elements to the left of the index i in the array nums. If there is no such element, leftSum[i] = 0.
rightSum[i] is the sum of elements to the right of the index i in the array nums. If there is no such element, rightSum[i] = 0.
Return the array answer.
 */
//Problem Link: https://leetcode.com/problems/left-and-right-difference/

// driver code
package main

import "fmt"

//main function
func main(){
	//input array of your choice
	nums := []int{1,2,3,4}
	//call the function
	fmt.Println(leftRigthDifference(nums))
	  }

// solution function
func leftRigthDifference(nums []int) []int {
	//create a slice to store the net sum
	var net_sum[]int
	//create a variable to store the total sum
	total_sum :=0
	//iterate through the array and calculate the total sum
	for _,val := range nums{
	  total_sum+=val
	}
	c:=0
	//iterate through the array and calculate the net sum
	for _,val := range nums{
	  //right_sum
	  total_sum-=val
	  x:= c - total_sum
	  //check if the value is negative
	  if x < 0{
		x*=-1
	  }
	  //append the value to the slice
	  net_sum = append(net_sum,x) 
	  // left_sum
	  c+=val
	}
	return net_sum
  }


