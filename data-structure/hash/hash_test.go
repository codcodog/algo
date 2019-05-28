package hashtable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	h := New()
	h.Put("hello", "world")
	h.Put(10, "nihao")
	assert.Equal(t, 2, h.Size())

	fmt.Println("TestPut: ")
	h.Print()
	fmt.Println("")
}

func TestGet(t *testing.T) {
	h := New()
	h.Put("hello", "world")
	h.Put("nihao", "shijie")
	h.Put(102, "102")

	hello := h.Get("hello")
	nihao := h.Get("nihao")
	num := h.Get(102)

	assert.Equal(t, "world", hello)
	assert.Equal(t, "shijie", nihao)
	assert.Equal(t, "102", num)
}

func TestDelete(t *testing.T) {
	h := New()
	h.Put("hello", "world")
	h.Put("nihao", "shijie")

	fmt.Println("TestDelete Before: ")
	h.Print()

	fmt.Println("TestDelete After: ")
	h.Delete("hello")
	h.Print()
	fmt.Println()

	assert.Equal(t, 1, h.Size())
}
