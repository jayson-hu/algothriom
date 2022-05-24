package main


//88. 合并两个有序数组 思路，采取倒序的方法，进行对比
func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		i int = m - 1
		j int = n - 1
	)
	for k := m + n - 1; k >= 0; k-- {
		if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}

	}

}
func main() {
	//nums1 = [1,2,3,0,0,0] m = 3, nums2 = [2,5,6], n = 3
}
