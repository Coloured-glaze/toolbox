// 算法
package algo

import "fmt"

// 冒泡排序， asc = true 为升序排列 false 为降序
func Mpp(arr *[]int, asc bool) *[]int {
	r := len(*arr) - 1
	if asc {
		for i := 0; i < r; i++ {
			for j := 0; j < r-i; j++ {
				if (*arr)[j] < (*arr)[j+1] {
					(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				}
			}
		}
	} else {
		for i := 0; i < r; i++ {
			for j := 0; j < r-i; j++ {
				if (*arr)[j] > (*arr)[j+1] {
					(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				}
			}
		}
	}
	return arr
}

// 二分查找
func BinaryFind(arr *[]int, findv int) (bool, int) {
	left, right, middle := 0, len(*arr)-1, 0
table:
	if left > right {
		return false, -1
	}
	middle = (left + right) / 2
	if (*arr)[middle] > findv {
		right = middle - 1
		middle--
		goto table
	} else if (*arr)[middle] < findv {
		left = middle + 1
		middle++
		goto table
	} else {
		return true, middle
	}
}

func BinaryFind2(arr *[]int, leftIndex int, rightIndex int, findVal int) {

	if leftIndex > rightIndex {
		fmt.Println(false, -1)
		return
	}
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal { // middle 左移
		BinaryFind2(arr, leftIndex, middle-1, findVal)

	} else if (*arr)[middle] < findVal { // middle 右移
		BinaryFind2(arr, middle+1, rightIndex, findVal)

	} else {
		fmt.Println(true, middle)
		return
	}
}
