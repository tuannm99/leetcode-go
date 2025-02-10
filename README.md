# Dadastructure

## 1. Array/Slices

## 2. HashMap

## 3. Graph

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

## 4. Linked List
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

## 5. Tree
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

