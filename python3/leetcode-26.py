# python3
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        n = 0
        for i in range(len(nums)):
            if i == 0:
                nums[n] = nums[i]
                n = n + 1
            if i != 0 and nums[i] != nums[i - 1]:
                nums[n] = nums[i]
                n = n + 1
        return n