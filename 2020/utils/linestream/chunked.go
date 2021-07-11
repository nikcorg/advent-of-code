package linestream

type Chunk []*Line

type ChunkedLineChan = chan Chunk
type ReadOnlyChunkedLineChan = <-chan Chunk
type WriteOnlyChunkedLineChan = chan<- Chunk

func WithChunking(inp ReadOnlyLineChan) ReadOnlyChunkedLineChan {
	out := make(ChunkedLineChan, cap(inp))

	Chunked(inp, out)

	return out
}

func Chunked(inp ReadOnlyLineChan, out WriteOnlyChunkedLineChan) {
	var chunk Chunk

	go func() {
		defer close(out)

		for line := range inp {
			if line.Content() == "" {
				if len(chunk) > 0 {
					out <- chunk
				}
				chunk = Chunk{}
				continue
			}

			chunk = append(chunk, line)
		}

		if len(chunk) > 0 {
			out <- chunk
		}
	}()
}
