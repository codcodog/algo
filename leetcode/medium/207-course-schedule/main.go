package main

import "fmt"

func main() {
	numCourses := 2
	prerequisites := [][]int{[]int{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisites)) // expected: true

	numCourses1 := 2
	prerequisites1 := [][]int{[]int{1, 0}, []int{0, 1}}
	fmt.Println(canFinish(numCourses1, prerequisites1)) // expected: false
}

type Queue struct {
	data []int
	top  int
}

func newQueue(size int) *Queue {
	queue := new(Queue)
	queue.data = make([]int, size)
	queue.top = -1
	return queue
}

func (q *Queue) empty() bool {
	return q.top < 0
}

func (q *Queue) push(n int) {
	q.top++
	q.data[q.top] = n
}

func (q *Queue) pop() int {
	data := q.data[q.top]
	q.top--
	return data
}

// 拓扑排序
// 查看有向图是否存在环
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegrees := make([]int, numCourses)
	adj := make([][]int, numCourses)

	for _, value := range prerequisites {
		first := value[1]
		second := value[0]
		inDegrees[second] += 1
		adj[first] = append(adj[first], second)
	}

	queue := newQueue(numCourses)
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue.push(i)
		}
	}

	counter := 0
	for !queue.empty() {
		top := queue.pop()
		counter++

		for _, successor := range adj[top] {
			inDegrees[successor] -= 1
			if inDegrees[successor] == 0 {
				queue.push(successor)
			}
		}
	}

	return counter == numCourses
}
