//Problem :2652 Sum Multiples
//Statement: Given a positive integer n, return the sum of all the multiples of 3, 5, or 7.
//Example 1:
//Input: n = 20
//Output: 98
//Explanation: multiples of 3, 5 and 7 are 3, 5, 6, 7, 9, 10, 12, 14, 15, 18, so their sum is 98.

//Problem Link: https://leetcode.com/problems/sum-multiples/

func sumOfMultiples(n int) int {
    c:=0
    for i:=1;i<=n;i++{
        if i%3==0 || i%5==0 || i%7==0{
            c+=i
        }
    }
    return c
    
}