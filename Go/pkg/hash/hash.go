package hash

const base = 1024

type MyHashSet struct {
	data [][]int
}

func Constructor() MyHashSet {
	return MyHashSet{make([][]int, base)}

}

func (this *MyHashSet) Add(key int) {
	idx := this.hash(key)
	if !this.Contains(key) {
		this.data[idx] = append(this.data[idx], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	idx := this.hash(key)
	for i := 0; i < len(this.data[idx]); i++ {
		if key == this.data[idx][i] {
			suffix := this.data[idx][i+1:]
			this.data[idx] = this.data[idx][:i]
			this.data[idx] = append(this.data[idx], suffix...)
		}
	}

}

func (this *MyHashSet) hash(key int) int {
	return key % base
}

func (this *MyHashSet) Contains(key int) bool {
	idx := this.hash(key)
	for i := 0; i < len(this.data[idx]); i++ {
		if key == this.data[idx][i] {
			return true
		}
	}
	return false
}
