// Problem Statement 1450. Number of Students Doing Homework at a Given Time
/* Problem Statement:
Given two integer arrays startTime and endTime and given an integer queryTime.

The ith student started doing their homework at the time startTime[i] and finished it at time endTime[i].

Return the number of students doing their homework at time queryTime. More formally, return the number of students where queryTime lays in the interval [startTime[i], endTime[i]] inclusive.
*/
// Problem Link: https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	counter:=0
	for i:=0;i<len(startTime);i++{
	  if startTime[i]<= queryTime && endTime[i] >= queryTime{
		counter+=1
	  }
	}
	return counter
	  
  }
