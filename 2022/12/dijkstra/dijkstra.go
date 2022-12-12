package dijkstra

import (
	"container/heap"
	"fmt"
	"math"
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

type PointGenerator func() (Point, error)

func Dijkstra(gW, gH int, start, end Point, points <-chan Point, cost func(Point, Point) (int, error)) ([]Point, error) {
	pq := make(PriorityQueue, 0)

	heap.Init(&pq)

	visited := map[Point]struct{}{}
	dist := map[Point]int{}
	dist[start] = 0

	prev := map[Point]*Point{}

	vertices := map[Point]*Vertex{}
	for {
		p, ok := <-points
		if !ok {
			break
		}

		pDist, ok := dist[p]
		if !ok {
			pDist = math.MaxInt
		}
		v := &Vertex{
			value:    p,
			priority: pDist,
		}
		vertices[p] = v
		heap.Push(&pq, v)
	}

	for {
		if pq.Len() < 1 {
			break
		}

		u := heap.Pop(&pq).(*Vertex)
		visited[u.value] = struct{}{}

		ns := []Point{
			{u.value.X - 1, u.value.Y},
			{u.value.X, u.value.Y - 1},
			{u.value.X + 1, u.value.Y},
			{u.value.X, u.value.Y + 1},
		}

		for _, v := range ns {
			c, err := cost(u.value, v)
			if err != nil {
				continue
			} else if _, ok := visited[v]; ok {
				continue
			}

			uDist, ok := dist[u.value]
			alt := uDist + c
			if !ok {
				alt = c
			}

			if vDist, ok := dist[v]; !ok || alt < vDist {
				dist[v] = alt
				prev[v] = &u.value
				pq.update(vertices[v], v, alt)
			}
		}
	}

	path := []Point{}
	next := end

	for {
		if p, ok := prev[next]; !ok {
			return nil, fmt.Errorf("broken path to %+v from %+v", next, path[0])
		} else {
			path = append([]Point{next}, path...)
			next = *p
		}

		if next == start {
			break
		}
	}

	return path, nil
}
