//Problem Statement: Final Value of Variable After Performing Operations
//There is a programming language with only four operations and one variable X:
//++X and X++ increments the value of the variable X by 1.
//--X and X-- decrements the value of the variable X by 1.
//Initially, the value of X is 0.
//Given an array of strings operations containing a list of operations, return the final value of X after performing all the operations.
//Example 1:
//Input: operations = ["--X","X++","X++"]
//Output: 1
//Explanation: The operations are performed as follows:

//Problem Link: https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

func finalValueAfterOperations(operations []string) int {
	value :=0
	for _,opt:= range operations{
	  if string(opt[len(opt)-1]) == "-" || string(opt[0]) == "-"{
		value = value - 1
	  }else{
		value = value + 1
	  }
	}
	return value
	  
  }