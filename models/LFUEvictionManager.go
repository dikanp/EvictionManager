package models

// import "fmt"

type LFUEvictionManager struct {
	heap []MinHeap
}

func (lfuEvictionManager *LFUEvictionManager) push(key string) {
	checkData := lfuEvictionManager.checkDataHeap(key)
	if checkData == -1 {
		data := MinHeap{
			Frequent: 1,
			Key:      key,
		}
		lfuEvictionManager.heap = append(lfuEvictionManager.heap, data)
		lfuEvictionManager.minHeapifyUp(len(lfuEvictionManager.heap) - 1)
		return
	} else {
		lfuEvictionManager.minHeapifyUp(checkData)
	}
}

func (lfuEvictionManager *LFUEvictionManager) pop() string {
	deletedData := lfuEvictionManager.heap[0]
	lengthData := len(lfuEvictionManager.heap) - 1

	if len(lfuEvictionManager.heap) == 0 {

		return "no data"
	}

	lfuEvictionManager.heap[0] = lfuEvictionManager.heap[lengthData]
	lfuEvictionManager.heap = lfuEvictionManager.heap[:lengthData]
	lfuEvictionManager.minHeapifyDown(0)
	return deletedData.Key
}

func (lfuEvictionManager *LFUEvictionManager) clear() {
	lfuEvictionManager.heap = lfuEvictionManager.heap[:0]
}

func (lfuEvictionManager *LFUEvictionManager) minHeapifyUp(index int) {
	for lfuEvictionManager.heap[parent(index)].Frequent > lfuEvictionManager.heap[index].Frequent && index != 0 {
		lfuEvictionManager.swap(parent(index), index)
		index = parent(index)
	}
}

func (lfuEvictionManager *LFUEvictionManager) minHeapifyDown(index int) {
	lastIndex := len(lfuEvictionManager.heap) - 1
	left, right := lfuEvictionManager.left(index), lfuEvictionManager.right(index)
	childToCompare := 0

	for left >= lastIndex {
		if left == lastIndex {
			childToCompare = 1
		} else if lfuEvictionManager.heap[left].Frequent > lfuEvictionManager.heap[right].Frequent {
			childToCompare = right
		} else {
			childToCompare = left
		}

		if lfuEvictionManager.heap[index].Frequent < lfuEvictionManager.heap[childToCompare].Frequent {
			lfuEvictionManager.swap(index, childToCompare)
			index = childToCompare
			left, right = lfuEvictionManager.left(index), lfuEvictionManager.right(index)
		} else {
			return
		}
	}
}

func (lfuEvictionManager *LFUEvictionManager) checkDataHeap(key string) int {
	for index, value := range lfuEvictionManager.heap {
		if value.Key == key {
			lfuEvictionManager.heap[index].Frequent += 1
			return index
		}
	}
	return -1
}

func parent(i int) int {
	return (i - 1) / 2
}

func (lfuEvictionManager *LFUEvictionManager) left(i int) int {
	return 2*i + i
}

func (lfuEvictionManager *LFUEvictionManager) right(i int) int {
	return 2*i + 2
}

func (lfuEvictionManager *LFUEvictionManager) swap(i1, i2 int) {
	lfuEvictionManager.heap[i1], lfuEvictionManager.heap[i2] = lfuEvictionManager.heap[i2], lfuEvictionManager.heap[i1]
}