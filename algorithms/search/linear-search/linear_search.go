/*
Y tuong : duyet tuan tu cac phan tu mang
Do phuc tap : O(n) cho moi TH
*/
package linearsearch

func LinearSearch(arr []int, x int) int {
	for i := 0; i < len(arr); i++ {
		if x == arr[i] {
			return i
		}
	}
	return -1
}
