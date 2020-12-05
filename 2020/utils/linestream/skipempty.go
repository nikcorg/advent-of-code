package linestream

func SkipEmpty(in ReadOnlyLineChan, out WriteOnlyLineChan) {
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
