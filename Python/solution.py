#!/usr/bin/python

class Solution(object):
    def twoSum(sef, nums, target):
        for i in range(len(nums)):
            anotherLoc = nums.index(target - nums[i])
            if anotherLoc>=0:
                return [i,anotherLoc]
            else: 
                return []
        
    def isSelfCrossing(sef, x):
         """
         :type x: List[int] 
         :rtype : bool 
         """
         targetLoc = [x[0]-x[2],x[3]-x[1]]
         if targetLoc[0]>=0 and targetLoc[1]>=0:
             return True
         else:
             return False
    def combinationSum(sef, nums, target):
        """
        : type nums: List[int]
        : type target: int 
        : rtype: int 
        """
        cnt= 0
        nums= sorted(nums)
        if len(nums)<=0 or nums[0]>target: 
            return 0 
        else:
            cntSolution=[0]*target
            print cntSolution
            for i in range(target):
                sum=0
                for j in range(len(nums)):
                    if nums[j]<target:
                        sum+=target-nums[j]
                    else:
                        break
                cntSolution[i]+=sum
            print "Count Solution",cntSolution
        return cnt

nums=[2,11,15,7]
target=9
instance =Solution()
twoSum = instance.twoSum(nums, target)
print "Two nums' location:",twoSum

target=4
nums=[2,1,3]
print "Sum4 nums target" ,target, nums
sum4 = instance.combinationSum(nums, target)

steps =[
  [2,1,1,2],
  [1,2,3,4],
  [1,1,1,1],
]
for i in range(len(steps)):
    isCross = instance.isSelfCrossing(steps[i])
    print "Cross Y/N:",isCross
isCrossing = instance.isSelfCrossing([2,1,1,2])
