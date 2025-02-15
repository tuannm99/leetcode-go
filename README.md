# Dadastructure

## 1. Array/Slices

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(len(primes)) // 6 || cap(primes) -> 6
var s []int = primes[1:4]

names := [4]string{"John", "Paul", "George", "Ringo"}
fmt.Println(names) // [John Paul George Ringo]

a := names[0:2]
b := names[1:3]
fmt.Println(a, b) // [John Paul] [Paul George]

b[0] = "XXX"
fmt.Println(a, b) // [John XXX] [XXX George]
fmt.Println(names) // [John XXX George Ringo]


var s []int
fmt.Println(s, len(s), cap(s)) // [] 0 0
if s == nil {
    fmt.Println("nil!")
}
//print  nil!

a := make([]int, 5)
b := make([]int, 0, 5)

board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
}

```

## 2. Strings operation

```go

```

## 3. sort packages

```go

```

## 4. Map

## 5. Graph

### 1. Adjacency Matrix

```go
package main

import (
	"fmt"
)

func main() {
	// Directed graph with 4 nodes (0,1,2,3)
	graph := [][]int{
		{0, 1, 1, 0}, // Node 0 has edges to 1,2
		{0, 0, 1, 1}, // Node 1 has edges to 2,3
		{1, 0, 0, 1}, // Node 2 has edges to 0,3
		{0, 0, 0, 0}, // Node 3 has no outgoing edges
	}

	fmt.Println("Adjacency Matrix:")
	for _, row := range graph {
		fmt.Println(row)
	}
}
```

### 2. Adjacency List

```go
package main

import (
	"fmt"
)

func main() {
	// Directed graph using a map to store adjacency lists
	graph := map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {0, 3},
		3: {},
	}

	// Print adjacency list
	fmt.Println("Adjacency List:")
	for node, neighbors := range graph {
		fmt.Printf("%d -> %v\n", node, neighbors)
	}
}
```

### 3. Edge List

```go
package main

import (
	"fmt"
)

func main() {
	// Edge list (u, v, w)
	edges := []struct {
		u, v, w int
	}{
		{0, 1, 2}, // Edge from 0 to 1 with weight 2
		{0, 2, 4},
		{1, 2, 1},
		{1, 3, 7},
		{2, 3, 3},
	}

	// Print edge list
	fmt.Println("Edge List:")
	for _, edge := range edges {
		fmt.Printf("(%d -> %d, weight: %d)\n", edge.u, edge.v, edge.w)
	}
}
```

## 6. Linked List

```go
package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Print()
}
```

## 7. Tree

```go
package main

import (
	"fmt"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(data int) {
	if data < t.data {
		if t.left == nil {
			t.left = &TreeNode{data: data}
		} else {
			t.left.Insert(data)
		}
	} else {
		if t.right == nil {
			t.right = &TreeNode{data: data}
		} else {
			t.right.Insert(data)
		}
	}
}

func (t *TreeNode) InOrder() {
	if t.left != nil {
		t.left.InOrder()
	}
	fmt.Print(t.data, " ")
	if t.right != nil {
		t.right.InOrder()
	}
}

func main() {
	root := &TreeNode{data: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(13)
	root.Insert(17)
	fmt.Println("In-order traversal:")
	root.InOrder()
}
```

## 8. Heap

```go
package main

import (
	"errors"
	"fmt"
)

// MaxHeap structure
type MaxHeap struct {
	data []int
}

// NewMaxHeap initializes a new MaxHeap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{data: []int{}}
}

// Get parent index
func (h *MaxHeap) parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

// Get left child index
func (h *MaxHeap) left(i int) int {
	return 2*i + 1
}

// Get right child index
func (h *MaxHeap) right(i int) int {
	return 2*i + 2
}

// Insert a value into the heap
func (h *MaxHeap) insert(val int) {
	h.data = append(h.data, val)
	idx := len(h.data) - 1

	// Bubble-up
	for idx > 0 {
		parentIdx := h.parent(idx)
		if parentIdx == -1 || h.data[idx] <= h.data[parentIdx] {
			break
		}
		h.data[idx], h.data[parentIdx] = h.data[parentIdx], h.data[idx]
		idx = parentIdx
	}
}

// Remove an element at index i
func (h *MaxHeap) remove(i int) (int, error) {
	if i >= len(h.data) {
		return -1, errors.New("index out of range")
	}

	removedValue := h.data[i]
	h.data[i], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[i]
	h.data = h.data[:len(h.data)-1]

	h.heapify(i)
	return removedValue, nil
}

// Extract the maximum value (root)
func (h *MaxHeap) extractMax() (int, error) {
	if len(h.data) == 0 {
		return -1, errors.New("heap is empty")
	}
	max := h.data[0]
	h.remove(0)
	return max, nil
}

// Heapify at a given index
func (h *MaxHeap) heapify(i int) {
	largest := i
	left := h.left(i)
	right := h.right(i)

	if left < len(h.data) && h.data[left] > h.data[largest] {
		largest = left
	}

	if right < len(h.data) && h.data[right] > h.data[largest] {
		largest = right
	}

	if largest != i {
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		h.heapify(largest)
	}
}

// PriorityQueue based on MaxHeap
type PriorityQueue struct {
	heap *MaxHeap
}

// NewPriorityQueue initializes a new PriorityQueue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{heap: NewMaxHeap()}
}

// Enqueue adds an element to the priority queue
func (pq *PriorityQueue) Enqueue(val int) {
	pq.heap.insert(val)
}

// Dequeue removes and returns the highest priority element
func (pq *PriorityQueue) Dequeue() (int, error) {
	return pq.heap.extractMax()
}

// Peek returns the highest priority element without removing it
func (pq *PriorityQueue) Peek() (int, error) {
	if len(pq.heap.data) == 0 {
		return -1, errors.New("queue is empty")
	}
	return pq.heap.data[0], nil
}

// IsEmpty checks if the queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.heap.data) == 0
}

// Main function to test
func main() {
	pq := NewPriorityQueue()

	pq.Enqueue(10)
	pq.Enqueue(20)
	pq.Enqueue(5)
	pq.Enqueue(30)

	fmt.Println(pq.Dequeue()) // 30
	fmt.Println(pq.Dequeue()) // 20
	fmt.Println(pq.Dequeue()) // 10
	fmt.Println(pq.Dequeue()) // 5
	fmt.Println(pq.Dequeue()) // Error: queue is empty
}
```

# Common Patterns

## 1. Two Pointer

## 2. Sliding Window

## 3. BFS (Breadth-First Search)

## ... update

```go
package main

import (
	"fmt"
	"container/list"
)

func BFS(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value.(int)
		fmt.Print(node, " ")
		queue.Remove(element)
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}
}

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {0, 3},
		3: {},
	}
	fmt.Println("BFS Traversal:")
	BFS(graph, 0)
}
```

## 4. DFS (Depth-First Search)

```go
package main

import (
	"fmt"
)

func DFS(graph map[int][]int, node int, visited map[int]bool) {
	if visited[node] {
		return
	}
	fmt.Print(node, " ")
	visited[node] = true
	for _, neighbor := range graph[node] {
		DFS(graph, neighbor, visited)
	}
}

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {0, 3},
		3: {},
	}
	fmt.Println("DFS Traversal:")
	visited := make(map[int]bool)
	DFS(graph, 0, visited)
}
```
