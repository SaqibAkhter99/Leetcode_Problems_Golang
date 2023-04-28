// Name: Add Digits
// Problem Statement: Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
// Example 1:
// Input: num = 38
// Output: 2
// Explanation: The process is
// 38 --> 3 + 8 --> 11
// 11 --> 1 + 1 --> 2
// Since 2 has only one digit, return it.

// Problem Link: https://leetcode.com/problems/add-digits/

func addDigits(num int) int {
	for true{
	  s:= strconv.Itoa(num)
	  add:=0
	  if num < 10 {
		break
	  }else{
		for _,val:= range s{
		  add+=int(val) - 48
		}
		num = add
	  }
  
	}
	return num	
  }
  