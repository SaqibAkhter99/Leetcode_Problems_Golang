//Problem: 1672. Richest Customer Wealth
/* You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.
   A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.
*/
//Problem Link: https://leetcode.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
	z:=0
	for _,val :=range accounts{
	  add:=0
	  for _,i :=range val{
		add+=i
	  }
	  if add > z{
		z = add
	  }
	}
	return z
	  
  }