package linestream

func SkipEmpty(in ReadOnlyLineChan) ReadOnlyLineChan {
	out := make(LineChan, cap(in))

	go func() {
		defer close(out)

		for v := range in {
			if v == nil || v.Content() == "" {
				continue
			}

			out <- v
		}
	}()

	return out
}
