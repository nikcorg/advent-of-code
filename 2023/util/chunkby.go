package util

// source https://gist.github.com/mustafaturan/7a29e8251a7369645fb6c2965f8c2daf
func ChunkBy[X any](items []X, chunkSize int) (chunks [][]X) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
