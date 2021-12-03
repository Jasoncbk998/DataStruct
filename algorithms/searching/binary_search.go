/**
 * @Author liuxu22
 * @Description //TODO
 * @Data 2021/9/19 11:56 上午
 **/
package searching

/**
二分查找
传入一个有序数组,输入目标值,返回索引
核心思想就是通过移位,找到middle
然后通过arr[middle] 是否等于目标值,判断是在middle的左边还是右边
从而不断定位
*/

// Search 传入一个有序数组
func Search(sortedArray []int, el int) int {
	//一个数组分为两半
	init, end := 0, len(sortedArray)-1

	for init <= end {
		//找到中间索引
		middle := ((end - init) >> 1) + init
		//判断目标值是否为arr[middle]
		if sortedArray[middle] == el {
			return middle
		}
		if sortedArray[middle] < el {
			init = middle + 1
		} else {
			end = middle + 1
		}

	}
	return -1
}

func Search1(arr []int, ele int) int {
	init, end := 0, len(arr)-1

	for init <= end {
		middle := (end-init)>>1 + init

		if arr[middle] == ele {
			return middle
		}
		if arr[middle] < ele {
			init = middle + 1
		} else {
			end = middle - 1
		}
	}
	//找不到
	return -1
}
