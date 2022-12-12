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

func Dijkstra(start, end Point, cost func(Point, Point) (int, error), dist map[Point]int) ([]Point, map[Point]int, error) {
	pq := make(PriorityQueue, 0)

	heap.Init(&pq)

	visited := map[Point]struct{}{}

	if dist == nil {
		dist = map[Point]int{}
	}

	dist[start] = 0

	prev := map[Point]*Point{}

	vertices := map[Point]*Vertex{}
	heap.Push(&pq, &Vertex{value: start, priority: 0})

	for {
		if pq.Len() < 1 {
			break
		}

		u := heap.Pop(&pq).(*Vertex)
		visited[u.value] = struct{}{}

		for n := range u.value.Neighbours() {
			c, err := cost(u.value, n)
			if err != nil {
				continue
			} else if _, ok := visited[n]; ok {
				continue
			}

			v := &Vertex{value: n, priority: c}
			heap.Push(&pq, v)
			vertices[n] = v

			uDist, ok := dist[u.value]
			alt := uDist + c
			if !ok {
				alt = c
			}

			if vDist, ok := dist[n]; !ok || alt < vDist {
				dist[n] = alt
				prev[n] = &u.value
				pq.update(vertices[n], n, alt)
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
