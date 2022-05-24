# python3
# leetcode 88 思路：采用倒序的方式来比较，并且原地排序
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        i = m - 1
        j = n - 1
        for k in range(m + n - 1, -1, -1):
            if j < 0 or (i >= 0 and nums1[i] >= nums2[j]):
                nums1[k] = nums1[i]
                i -= 1
            else:
                nums1[k] = nums2[j]
                j -= 1
