//Problem 1827. Minimum Operations to Make the Array Increasing
// You are given an integer array nums (0-indexed). In one operation, you can choose an element of the array and increment it by 1.

// For example, if nums = [1,2,3], you can choose to increment nums[1] to make nums = [1,3,3].
// Return the minimum number of operations needed to make nums strictly increasing.

// An array nums is strictly increasing if nums[i] < nums[i+1] for all 0 <= i < nums.length - 1. An array of length 

//Problem link https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing/


func minOperations(nums []int) int {
	counter:=0
	iter:=0
	for iter < len(nums)-1{
	  curr:=nums[iter]
	  next:=nums[iter+1]
	  fmt.Println(curr,next)
	  if curr >= next{
		diff:= curr-next + 1
		next = curr+1
		nums[iter+1] = next
		counter+=diff
	  }
	  iter+=1
	}
	return counter
	  
  }