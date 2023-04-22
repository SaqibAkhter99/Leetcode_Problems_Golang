//Problem Statement: Kids With the Greatest Number of Candies
//Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.
//For each kid check if there is a way to distribute extraCandies among the kids such that he or she can have the greatest number of candies among them.
//Notice that multiple kids can have the greatest number of candies.
//Example 1:
//Input: candies = [2,3,5,1,3], extraCandies = 3
//Output: [true,true,true,false,true]
//Explanation: If you give all extraCandies to:
//- Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
//- Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
//- Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
//- Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
//- Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.

//Problem Link: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/



func kidsWithCandies(candies []int, extraCandies int) []bool {
	x:=true
	y:=false
	z :=0
	var c[]bool
	for i:=0;i<len(candies);i++{
	  if candies[i] > z{
		z = candies[i]
	  }
	}
	for _,val := range candies{
	  if val+extraCandies >=z{
		c = append(c,x)
	  } else{
		c = append(c,y)
	  }
	}
	return c
	  
  }