package data_structure

import (
	"fmt"
)

type HashTable struct {
	Values map[int]interface{}
}

func NewHash() *HashTable {
	return &HashTable{
		make(map[int]interface{})}
}

func (h *HashTable) Put(k, v interface{}) {
	hash := hash(k)
	h.Values[hash] = v
}

func (h *HashTable) Get(k interface{}) interface{} {
	hash := hash(k)

	return h.Values[hash]
}

func (h *HashTable) Delete(k interface{}) {
	hash := hash(k)
	delete(h.Values, hash)
}

func (h *HashTable) Size() int {
	return len(h.Values)
}

func (h *HashTable) Print() {
	for key, value := range h.Values {
		fmt.Printf("%d => %s\n", key, value)
	}
}

// 若 key 过长，hash 会溢出，这里暂不考虑
// Horner’s Method
func hash(key interface{}) int {
	k := fmt.Sprintf("%s", key)

	hash := 0
	for i := 0; i < len(k); i++ {
		hash = hash*31 + int(k[i])
	}
	return hash
}
