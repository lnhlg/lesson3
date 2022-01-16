package lesson3

import (
	"sort"
)

//PermuteUnique: 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列
/*parameter
nums: 整数数组
return: nums的不重复的全排列二维数组
*/
func PermuteUnique(nums []int) [][]int {

	//排序
	sort.Ints(nums)

	var result [][]int
	var path []int
	used := make([]bool, len(nums))

	//递归
	var backtracking func()
	backtracking = func() {

		if len(path) == len(nums) {

			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)

			return
		}

		for i := 0; i < len(nums); i++ {

			// 去重
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}

			if !used[i] {

				used[i] = true
				path = append(path, nums[i])
				backtracking()
				//回溯
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}

	backtracking()

	return result
}
