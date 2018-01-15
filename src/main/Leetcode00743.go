package main

func networkDelayTime(times [][]int, N int, K int) int {
	graph := make(map[int][][]int, len(times))
	for _, time := range times {
		graph[time[0]] = append(graph[time[0]], []int{time[1], time[2]})
	}
	v := make(map[int]bool, len(times))
	dist := make(map[int]int, len(times))
	dist[K] = 0
	i := K
	s := make([]int, 0, len(times))
	s = append(s, K)
	for len(s) > 0 {
		i, s = s[len(s)-1], s[:len(s)-1]
		v[i] = true
		edges := graph[i]
		for _, e := range edges {
			if v, ok := dist[e[0]]; !ok || v > dist[i]+e[1] {
				dist[e[0]] = dist[i] + e[1]
			} else {
				continue
			}
			s = append(s, e[0])
		}
	}
	if len(v) != N {
		return -1
	}
	max := -1
	for _, d := range dist {
		if d > max {
			max = d
		}
	}
	return max
}
