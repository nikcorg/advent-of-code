package s5

func binsert(ns []int, n int) []int {
	if len(ns) > 0 {
		// check for edge cases
		if n < ns[0] {
			return append([]int{n}, ns...)
		} else if n > ns[len(ns)-1] {
			return append(ns, n)
		}

		// middle insertion
		lo := 0
		hi := len(ns) - 1
		c := hi / 2

		for lo < hi {
			// did we find our spot?
			if ns[c] < n && ns[c+1] > n {
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
