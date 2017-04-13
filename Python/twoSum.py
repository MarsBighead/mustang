#!/usr/bin/python

class Solution(object):
    def twoSum(sef, nums, target):
        for i in range(len(nums)):
            anotherLoc = nums.index(target - nums[i])
            if anotherLoc>=0:
                return [i,anotherLoc]
            else: 
                return []
        
        
        
nums=[2,11,15,7]
target=9
instance =Solution()
twoSum = instance.twoSum(nums, target)
print "Two nums' location:",twoSum
