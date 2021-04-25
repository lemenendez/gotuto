package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enque(ele int) {
	// prepend https://stackoverflow.com/questions/53737435/how-to-prepend-int-to-slice
	q.items = append([]int{ele}, q.items...)
} 

func (q *Queue) Deque() int {
	index := len(q.items) - 1
	ele := q.items[index]
	q.items =  q.items[:index]
	return ele
}

func NewQueue() *Queue {
	return &Queue{}
}

type Stack struct {
	items []int
}

func (s *Stack) Push(ele int) {
	s.items = append(s.items, ele)
}

func (s *Stack) Pop() int {
	index := len(s.items) - 1
	ele := s.items[index]
	s.items =  s.items[:index]
	return ele
}

func NewStack() *Stack {
	s := &Stack{}
	return s
}

func (s Queue) Empty() bool {
	return len(s.items)==0
}

func (s Stack) Empty() bool {
	return len(s.items)==0
}

type Graph struct {
	visited map[int] bool
	adj map[int] []int
}

func (g *Graph) addEdge(v int, w int) {
	g.adj[v]= append(g.adj[v], w)
}

func NewGraph() *Graph  {
	var g *Graph
	g = &Graph{}
	g.adj = make(map[int] []int)
	return g
}

func (g *Graph ) dfs(v int) {
	g.visited[v] = true
	fmt.Printf("%v ", v)
	for _,n := range g.adj[v] {
		if !g.visited[n] {
			g.dfs(n)
		}
	}
}

func (g *Graph) Dfs(v int) {
	// reset the visited map
	g.visited = make(map[int] bool)	
	g.dfs(v)
}

func (g *Graph) StackDfs(v int) {
	g.visited = make(map[int] bool)	
	s:=NewStack()
	s.Push(v)
	g.visited[v] = true
	for !s.Empty() {
		newv := s.Pop()
		fmt.Printf("%v", newv)
		for _, n := range g.adj[newv] {
			if !g.visited[n] {
				s.Push(n)
				g.visited[n] = true
			}
		}
	}
}

func (g *Graph) Bfs(v int) {
	g.visited = make(map[int] bool)	
	q :=NewQueue()
	q.Enque(v)
	g.visited[v] = true
	for !q.Empty() {
		curv := q.Deque()
		fmt.Printf("%v ", curv)
		for _, n := range g.adj[curv] {
			if !g.visited[n] {
				q.Enque(n)
				g.visited[n] = true
			}
		}
	}
}

func main() {
	g :=  NewGraph()

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	// fmt.Print(g)

	// g.Dfs(2)
	// g.StackDfs(2)
	/*
	q :=NewQueue()

	q.Enque(5)
	q.Enque(5)
	fmt.Print(q)
	*/
	g.Bfs(2)
	
}