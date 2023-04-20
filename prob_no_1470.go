//Problem: 1470. Shuffle the Array
/* Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
Return the array in the form [x1,y1,x2,y2,...,xn,yn]. */

//Problem Link: https://leetcode.com/problems/shuffle-the-array/description/


func shuffle(nums []int, n int) []int {

	var x[]int
	
	for i:=0;i<n;i++{
		 x = append(x,nums[i])
		 x = append(x,nums[n+i])
	}
	return x
	  
  }