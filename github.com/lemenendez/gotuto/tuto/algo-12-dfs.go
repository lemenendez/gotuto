package main

import "fmt"

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
	fmt.Printf("%v ",v)
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

func main() {
	g :=  NewGraph()

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	// fmt.Print(g)

	g.Dfs(2)

}