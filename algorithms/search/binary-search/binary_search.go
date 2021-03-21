/*
Y tuong : dua tren mang da sap xep, tim tu phan thu trung tam cua mang dem so sanh voi x de tim tiep o ben phai hoac ben trai
phan tu trung tam. Viec tim o ben phai/trai phan tu trung tam duoc dien ra tuong tu
Do phuc tap : O(n) cho moi TH
*/
package binarysearch

func main() {

}
func BinarySearch(arr []int, x int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		median := (left + right) / 2
		if arr[median] == x {
			return median
		} else if x > arr[median] {
			left = median + 1
		} else {
			right = median - 1
		}
	}
	return -1
}
