//Problem Statement: Remove All Adjacent Duplicates In String
//Given a string s of lowercase letters, a duplicate removal consists of choosing two adjacent and equal letters, and removing them.
//We repeatedly make duplicate removals on s until we no longer can.
//Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.
//Example 1:
//Input: s = "abbaca"
//Output: "ca"
//Explanation:
//For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move. The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
//Problem Link: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/




func removeDuplicates(s string) string {
	var stack[]string
	j:=1
	stack = append(stack,string(s[0]))
	for j < len(s){
		if len(stack) < 1 || string(s[j])!=string(stack[len(stack)-1]){
		  stack = append(stack,string(s[j]))
		}else{
		  stack = stack[:len(stack)-1]
	  }
	  j+=1
	}
	check := strings.Join(stack[:],"")
	return check
  }
  