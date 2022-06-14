package data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue(3)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Print()

	q.Enqueue(4)
	q.Print()
}

func TestDequeue(t *testing.T) {
	q := NewQueue(3)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	first := q.Dequeue()
	assert.Equal(t, 1, first)
	q.Print()
}
