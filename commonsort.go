package commonsort

import (
	"algo/utility"
)

// 比较排序：
//		交换排序：冒泡排序、快排
//		插入排序：插入排序、Shell排序
//		选择排序：选择排序、堆排序
//		归并排序：二路归并排序、多路归并排序
// 非比较排序：
//		计数排序、桶排序、基数排序


// 下面为比较排序：
// 交换排序 - 冒泡排序
// 这个算法的名字由来是因为越小/大的元素会经由交换慢慢“浮”到数列的顶端，就如同水中气泡最终会上浮到顶端一样，故名“冒泡排序”
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; i < n-1-j; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 交换排序 - 快排
// 分治法：通过一趟排序将待排元素分隔成独立的两部分，其中一部分均比另一部分小。则可分别对这两部分元素继续进行排序，以达到整个序列有序。
func QuickSort(arr []int, left int, right int) {
	if left < right {
		i := left
		j := right
		key := arr[left]
		for i < j {
			for i < j && arr[j] >= key {
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}
			for i < j && arr[i] <= key {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		arr[i] = key

		QuickSort(arr, left, i-1)
		QuickSort(arr, i+1, right)
	}
}

// 插入排序
// 通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
// 插入排序在实现上采用in-place排序（即只需用到O(1)的额外空间的排序），因而在从后向前扫描过程中，需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。
func InsertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// 插入排序 - Shell排序
// 1959年Shell发明，第一个突破O(n2)的排序算法，是简单插入排序的改进版。它与插入排序的不同之处在于，它会优先比较距离较远的元素。希尔排序又叫“缩小增量排序”。
// 个人觉得：
// 1. 希尔排序比较鸡肋，没有快排快，稍微有点难懂。
// 2. 代码实现上，在简单插入排序的基础上，在外层嵌套一个gap for循环。
// 3. 当n>1w时（n^1.3 = 15.8w，nlogn = 13.3w），希尔排序就没有快排快了。
func ShellSort(arr []int) {
	n := len(arr)
	k := n / 2
	for k > 0 {
		for i := k; i < n; i++ {
			for j := i; j >= k && arr[j] < arr[j-k]; j -= k {
				arr[j], arr[j-k] = arr[j-k], arr[j]
			}
		}
		k /= 2
	}
}

// 选择排序
// 首先，在未排序序列中找到最小（大）元素，存放到排序序列的起始位置;
// 然后，再从剩余未排序元素中继续寻找最小（大）元素，放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。
func SelectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ { // 选择最小元素，插入到第i位置
		tmp := i
		for j := i; j < n; j++ {
			if arr[j] < arr[tmp] {
				tmp = j
			}
		}
		if tmp != i {
			arr[i], arr[tmp] = arr[tmp], arr[i]
		}
	}
}

// 选择排序 - 堆排序
// 利用堆这种数据结构所设计的一种排序算法。堆：子结点的键值或索引总是小于（或者大于）它的父节点。
// 算法描述：
// 1. 将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
// 2. 将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
// 3. 由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆。
// 4. 再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。
// 5. 不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。
func HeapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapAdjust(arr, i, n)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		heapAdjust(arr, 0, i)
	}
}

func heapAdjust(arr []int, jIdx int, len int) {
	child := 0
	for i := jIdx; i*2+1 < len; i = child {
		child = i*2 + 1
		if child+1 < len && arr[child] < arr[child+1] {
			child++
		}

		if arr[i] < arr[child] {
			arr[i], arr[child] = arr[child], arr[i]
		}
	}
}

// 归并排序
// 分治法（Divide and Conquer）:先将候选序列拆分成2个子序列，使每个子序列有序，然后将这2个子序列合并。 “使子序列有序”是一个递归的过程。
func MergeSort(arr []int, begin int, end int) {
	if begin < end {
		mid := (begin + end) / 2
		MergeSort(arr, begin, mid)
		MergeSort(arr, mid+1, end)
		merge(arr, begin, mid, end)
	}
}

func merge(arr []int, begin int, mid int, end int) {
	result := make([]int, 0, end-begin+1)
	left, right := begin, mid+1

	for left <= mid && right <= end {
		if arr[left] < arr[right] {
			result = append(result, arr[left])
			left++
		} else {
			result = append(result, arr[right])
			right++
		}
	}

	if left <= mid {
		result = append(result, arr[left:mid+1]...)
	}
	if right <= end {
		result = append(result, arr[right:end+1]...)
	}

	for i := 0; i <= end-begin; i++ {
		arr[begin+i] = result[i]
	}
}

// 下面为非比较排序：
// 计数排序（Counting Sort），不是基于比较的排序算法，其核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。
// 作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
// 算法描述：
// 找出待排序的数组中最大和最小的元素；--下面代码简化了，只用了最小值用的是0.
// 统计数组中每个值为i的元素出现的次数，存入临时数组的第i项；
// 将临时数组里的数据灌入目标数组（也可以是源数组）：计数大于0的，顺序灌入。计数为几就灌几次。
func CountingSort(arr []int) {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	buff := make([]int, max+1)
	for _, v := range arr {
		buff[v] ++
	}

	jIdx := 0
	for i, v := range buff {
		for j := 0; j < v; j++ {
			if v > 0 {
				arr[jIdx] = i
				jIdx++
			}
		}
	}
}

// 桶排序(Bucket Sort)是计数排序的升级版（如果桶的数量足够多就会变成计数排序）。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。
// 工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。
// 算法描述
// 设置一个定量的数组当作空桶；
// 遍历输入数据，并且把数据一个一个放到对应的桶里去；
// 对每个不是空的桶进行排序；
// 从不是空的桶里把排好序的数据拼接起来。
func BucketSort(arr []int) {
	// 初始化桶
	bucketCount := 100 // 假设有100个桶
	bucket := [][]int{}
	for i := 0; i < bucketCount; i++ {
		tmp := make([]int, 1)
		bucket = append(bucket, tmp)
	}

	// 将数据分配到桶中
	for _, v := range arr {
		bucket[v/bucketCount] = append(bucket[v/bucketCount], v)
	}

	// 循环所有的桶进行排序
	jIdx := 0
	for i := 0; i < bucketCount; i++ {
		if len(bucket[i]) > 1 {
			// 对每个桶内元素进行排序，可以使用任何排序算法。
			QuickSort(bucket[i], 0, len(bucket[i])-1)

			// 重新灌入源数组
			for j := 1; j < len(bucket[i]); j++ {
				arr[jIdx] = bucket[i][j]
				jIdx++
			}
		}
	}
}

// 基数排序（Radix Sort）是桶排序(Bucket Sort)的一种。
// 这里的桶为每一位上数字的范围，为【0-9】十个桶。分别对个位数、十位数、百位数…上面的数字分级排序。
func RadixSort(arr []int) []int {
	// 最大值，确定数量级
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	// 按照数量级分段
	for bit := 1; max/bit > 0; bit *= 10 {
		// 每次处理一个级别的排序
		arr = bitSort(arr, bit)
	}

	utility.PrintArray(arr)

	return arr
}


// 按照位数排序
func bitSort(arr []int, bit int) []int {
	// 数组长度
	length := len(arr)

	// 统计长度 0~9
	bitCounts := make([]int, 10)

	for i := 0; i < length; i++ {
		// 余数
		num := (arr[i] / bit) % 10

		// 统计余数相等的个数，进行递增
		bitCounts[num]++
	}

	// 叠加（目的是增加稳定性）
	// 原来是0~9，现在变成描述0~9的累加数量，这样就可以很方便的知道最终的坐标
	for j := 1; j < len(bitCounts); j++ {
		bitCounts[j] += bitCounts[j-1]
	}

	// 临时数组
	res := make([]int, length)

	// 将原来的数组倒序遍历一遍
	for i := length - 1; i >= 0; i-- {
		// 余数
		num := (arr[i] / bit) % 10

		res[bitCounts[num]-1] = arr[i]

		// 统计余数相等的个数，进行递减
		bitCounts[num]--
	}

	return res
}
