//Problem: 1929. Concatenation of Array
/* Given an integer array nums of length n, you want to create an array ans of length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).
   Specifically, ans is the concatenation of two nums arrays.
   Return the array ans.
*/
//Problem Link: https://leetcode.com/problems/concatenation-of-array/description/



func getConcatenation(nums []int) []int {
  
	//append the array to itself
	s := append(nums, nums...)
	return s
	  
  }