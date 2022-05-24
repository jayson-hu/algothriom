# python3
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = 0
        len_nums = len(nums)
        for i in range(len_nums):
            if nums[i] != 0:
                nums[n] = nums[i]
                n += 1
        while n < len_nums:
            nums[n] = 0
            n += 1
