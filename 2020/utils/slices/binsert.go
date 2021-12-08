package slices

func binsert(ns []int, n int) []int {
	if len(ns) > 0 {
		// check for edge cases
		if n == ns[0] || n == ns[len(ns)-1] { // already in at the fringes
			return ns
		} else if n < ns[0] { // fits in at the head?
			return append([]int{n}, ns...)
		} else if n > ns[len(ns)-1] { // fits in at the tail?
			return append(ns, n)
		}

		// middle insertion
		lo := 0
		hi := len(ns) - 1
		c := hi / 2

		for lo < hi {
			if ns[c] == n || ns[c+1] == n { // is it already in there?
				return ns
			} else if ns[c] < n && ns[c+1] > n { // did we find our spot?
				break
			}

			if ns[c] < n {
				// discard lower half
				lo = c + 1
				c = lo + (hi-lo)/2
			} else if ns[c] > n {
				// discard upper half
				hi = c - 1
				c = hi - (hi-lo)/2
			}
		}

		head := ns[0 : c+1]
		tail := ns[c+1:]

		return append(head, append([]int{n}, tail...)...)
	}

	return append(ns, n)
}
