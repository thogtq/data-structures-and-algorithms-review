/*
Build heap tu vi tri n-1/2 bang ham heaptify
Ham heaptify se kiem tra vi tri nut trai, phai tai vi tri k de dam bao tai k co gia tri lon nhat
Ham sort se chay tu n-1, sau moi loop se giam di 1 va sap xep lai heap
*/

package heapsort

//input: mang cac phan tu ngau nhien
//output
func buildHeap(array []int) {
	for k := (len(array) - 1) / 2; k >= 0; k-- {
		heaptify(array, k, len(array))
	}
}

//input: mang chua cac phan tu cay, vi tri can heaptify, kich co heap
//output: mang dai dien cho mot heap
func heaptify(array []int, k int, heapSize int) {
	leftNodeIdx := k*2 + 1
	highestNodeIdx := leftNodeIdx
	for highestNodeIdx < heapSize {
		rightNodeIndx := highestNodeIdx + 1
		if rightNodeIndx < heapSize {
			if array[highestNodeIdx] < array[rightNodeIndx] {
				highestNodeIdx = rightNodeIndx
			}
		}
		if array[k] > array[highestNodeIdx] {
			return
		}
		//swap k voi highestNode
		temp := array[k]
		array[k] = array[highestNodeIdx]
		array[highestNodeIdx] = temp
		k = highestNodeIdx
		highestNodeIdx = 2*k + 1
	}
}

//input: mang cac phan tu ngau nhien\n
//output: mang cac phan tu duoc sap xep tang dan
func HeapSort(array []int) {
	buildHeap(array)
	heapSize := len(array)
	for heapSize > 0 {
		heapSize--
		temp := array[0]
		array[0] = array[heapSize]
		array[heapSize] = temp
		heaptify(array, 0, heapSize)
	}
}
