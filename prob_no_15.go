// Problem: 15. 3Sum
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.
// Example 1:	Input: nums = [-1,0,1,2,-1,-4]	Output: [[-1,-1,2],[-1,0,1]]
// Example 2:	Input: nums = []	Output: []
// Example 3:	Input: nums = [0]	Output: []
// Constraints: 0 <= nums.length <= 3000 -105 <= nums[i] <= 105
// Problem Link: https://leetcode.com/problems/3sum/



func threeSum(nums []int) [][]int {
	//sort the array
    sort.Ints(nums)
    j:=0
	//create a slice to store the triplets
    var z [][]int
    for j < len(nums){
		// declare two pointers
        var l,r int = j+1,len(nums)-1
		// while the left pointer is less than the right pointer
        for l < r{
            var y[]int
			//calculate the sum of the three numbers
            add:=nums[j] + nums[l] + nums[r]
			//if the sum is less than zero, increment the left pointer
            if add < 0{
                l+=1
			//if the sum is greater than zero, decrement the right pointer	
            } else if add > 0{
                r-=1
			//if the sum is zero, append the triplet to the slice and increment the left pointer and decrement the right pointer	
            } else if add==0{
               y = append(y, nums[j])
               y = append(y, nums[l])
               y = append(y, nums[r])
               z = append(z, y)
               l+=1
               r-=1
            }
        }
            j+=1
      }
	//remove the duplicate triplets  
    seen := make(map[string]bool)
    output := make([][]int, 0, len(z))
    for _, val := range z {
        key := fmt.Sprintf("%v", val)
        if _, ok := seen[key]; !ok {
            seen[key] = true
            output = append(output, val)
        }
    }
	//return the slice
    return output  
}
  