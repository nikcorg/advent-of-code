package dijkstra

import (
	"container/heap"
	"fmt"
)

// 1  function Dijkstra(Graph, source):
// 2
// 3      create vertex set Q
// 4
// 5      for each vertex v in Graph:
// 6          dist[v] ← INFINITY
// 7          prev[v] ← UNDEFINED
// 8          add v to Q
// 9      dist[source] ← 0
// 10
// 11      while Q is not empty:
// 12          u ← vertex in Q with min dist[u]
// 13
// 14          remove u from Q
// 15
// 16          for each neighbor v of u still in Q:
// 17              alt ← dist[u] + length(u, v)
// 18              if alt < dist[v]:
// 19                  dist[v] ← alt
// 20                  prev[v] ← u
// 21
// 22      return dist[], prev[]

func FindPath(start, end Point, cost func(Point, Point) (int, error)) ([]Point, map[Point]int, error) {
	visited := map[Point]struct{}{}
	prev := map[Point]*Point{}
	dist := map[Point]int{}
	vertices := map[Point]*Vertex{}
	pq := make(PriorityQueue, 0)

	dist[start] = 0

	heap.Init(&pq)
	heap.Push(&pq, &Vertex{value: start, priority: 0})

	for {
		if pq.Len() < 1 {
			break
		}

		currV := heap.Pop(&pq).(*Vertex)
		curr := currV.value
		currDist := dist[curr]

		visited[currV.value] = struct{}{}

		for next := range curr.Neighbours() {
			c, err := cost(curr, next)
			if err != nil {
				continue
			} else if _, ok := visited[next]; ok {
				continue
			}

			v := &Vertex{value: next, priority: c}
			heap.Push(&pq, v)
			vertices[next] = v

			distViaCurr := currDist + c

			if prevDistToNext, ok := dist[next]; !ok || distViaCurr < prevDistToNext {
				dist[next] = distViaCurr
				prev[next] = &currV.value
				pq.update(vertices[next], next, distViaCurr)
			}
		}
	}

	path := []Point{}
	next := end

	for {
		if p, ok := prev[next]; !ok {
			if len(path) == 0 {
				return nil, dist, fmt.Errorf("broken path to %+v", next)
			}
			return nil, dist, fmt.Errorf("broken path to %+v from %+v", next, path[0])
		} else {
			path = append([]Point{next}, path...)
			next = *p
		}

		if next == start {
			break
		}
	}

	return path, dist, nil
}
