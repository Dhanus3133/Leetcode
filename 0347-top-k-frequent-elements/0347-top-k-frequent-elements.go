func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		if _, pres := m[num]; pres {
			m[num]++
		} else {
			m[num] = 1
		}
	}
	ans := []int{}

	h := &KVHeap{}
	heap.Init(h)

	for c, d := range m {
		heap.Push(h, KV{k: c, v: d})
	}
	for i := 0; i < k; i++ {
		temp := heap.Pop(h)
		ans = append(ans, temp.(KV).k)
	}
	return ans
}

// Heap Implementation

type KV struct {
	k, v int
}

type KVHeap []KV

func (kv KVHeap) Less(i int, j int) bool {
	return kv[i].v > kv[j].v
}

func (kv KVHeap) Swap(i int, j int) {
	kv[i], kv[j] = kv[j], kv[i]
}

func (kv KVHeap) Len() int {
	return len(kv)
}

func (kv *KVHeap) Push(v interface{}) {
	*kv = append(*kv, v.(KV))
}

func (kv *KVHeap) Pop() interface{} {
	old := *kv
	n := len(old)
	x := old[n-1]
	*kv = old[0 : n-1]
	return x
}
