/*
Bắt đầu từ cuối dãy, các phần tử nhỏ sẽ được đẩy dần về phía trái dãy đến vị trí đúng của nó bằng cách thực
hiện các phép hoán vị 2 phần tử liên tiếp có tạo nghịch thế. Khi một phần tử nhỏ nhất được đưa về
đầu dãy thì chỉ xét việc hoán vị cho dãy mới gồm các phần tử còn lại
TH xau nhat : O(n^2)
*/
package bubblesort

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
