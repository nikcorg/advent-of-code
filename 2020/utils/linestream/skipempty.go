package linestream

func SkipEmpty(in LineChan, out LineChan) {
	go func() {
		defer close(out)

		for v := range in {
			if v == nil || v.Content() == "" {
				continue
			}

			out <- v
		}
	}()
}
