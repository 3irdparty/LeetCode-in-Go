package problem0033

func search(nums []int, target int) int {
	size := len(nums)

	left, right := 0, size-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	/* lo = hi，是最小值的索引值 */

	rotated := left /* 数组旋转了的距离 */

	left, right = 0, size-1

	for left <= right {
		mid := (left + right) / 2
		/* nums 是 rotated，所以需要使用 rotatedMid 来获取 mid 的值 */
		rotatedMid := (rotated + mid) % size
		switch {
		case nums[rotatedMid] > target:
			right = mid - 1
		case nums[rotatedMid] < target:
			left = mid + 1
		default:
			return rotatedMid
		}
	}

	return -1
}
