//Problem Statement: Build Array from Permutation

/* Given a zero-based permutation nums (0-indexed), build an array ans of the same length where ans[i] = nums[nums[i]] for each 0 <= i < nums.length and return it.
   A zero-based permutation nums is an array of distinct integers from 0 to nums.length - 1 (inclusive).
 */

func buildArray(nums []int) []int {
    // Calculate the length of the array
	x:=len(nums)

	//iterate through the array
	for i:=0;i<x;i++{
	  nums = append(nums,nums[nums[i]])
	}
	//slice the array
	ans_slice := nums[x:]
	return ans_slice
	  
  }