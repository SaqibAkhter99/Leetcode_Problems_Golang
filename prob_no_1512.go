//Problem Statement: Number of Good Pairs
//Given an array of integers nums return the number of good pairs.
//A pair (i,j) is called good if nums[i] == nums[j] and i < j.
//Example 1:
//Input: nums = [1,2,3,1,1,3]
//Output: 4
//Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
//Problem Link : https://leetcode.com/problems/number-of-good-pairs/


func numIdenticalPairs(nums []int) int {
	counter := 0
	for i:=0;i<len(nums);i++{
	  for j:=i+1;j<len(nums);j++{
		if nums[i] == nums[j]{
		  counter+=1
		}
	  }
	}
	return counter
	  
  }