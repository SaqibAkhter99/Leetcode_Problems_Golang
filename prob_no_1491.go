// Problem 1491: Average Salary Excluding the Minimum and Maximum Salary
// Statement: Given an array of unique integers salary where salary[i] is the salary of the employee i.
// Return the average salary of employees excluding the minimum and maximum salary.
// Example 1:
// Input: salary = [4000,3000,1000,2000]
// Output: 2500.00000
// Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
// Average salary excluding minimum and maximum salary is (2000+3000)/2= 2500
//Problem Link: https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

func average(salary []int) float64 {
	sort.Ints(salary)
	add:=0
	for i:=1;i<len(salary)-1;i++{
	  add+=salary[i]
	}
	x:=len(salary)-2
	mean:=float64(add)/float64(x)
	return mean
	  
  }