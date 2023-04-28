//2549. Count Distinct Numbers on Board
// Problem Statement: You are given a positive integer n, that is initially placed on a board. Every day, for 109 days, you perform the following procedure:
// For each number x present on the board, find all numbers 1 <= i <= n such that x % i == 1.
// Then, place those numbers on the board.
// Return the number of distinct integers present on the board after 109 days have elapsed.

// Note:

// Once a number is placed on the board, it will remain on it until the end.
// % stands for the modulo operation. For example, 14 % 3 is 2.

// Example 1:
// Input: n = 1
// Output: 1
// Explanation: The numbers initially on the board are [1].
// After the first day, the numbers on the board are [1,2].
// After 109 days, the numbers on the board are [1,2,3,...,n].
// There are n distinct numbers present on the board.
// Example 2:
// Input: n = 2
// Output: 2
// Explanation: The numbers initially on the board are [1,2].
// After the first day, the numbers on the board are [1,2,3].
// After 109 days, the numbers on the board are [1,2,3,4,5,...,n].
// There are n distinct numbers present on the board.


func distinctIntegers(n int) int {
    if n==1{
      return 1
    }
    return n-1
}