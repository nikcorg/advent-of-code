package linestream

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMux(t *testing.T) {
	const input = `line 1
line 2
line 3
line 4
line 5
line 6`

	output := make(LineChan)
	mux := NewMuxxer(output)
	wg := sync.WaitGroup{}

	// Setup listeners first
	for n := 0; n < 10; n++ {
		wg.Add(1)

		go func(recv ReadOnlyLineChan) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			expected := strings.Split(input, "\n")

			for {
				select {
				case <-ctx.Done():
					t.Error("test timed out")
					return

				case actual, ok := <-recv:
					if !ok {
						assert.Empty(t, expected)
						return
					}
					assert.Equal(t, expected[0], actual.Content())
					expected = expected[1:]
				}
			}
		}(mux.Recv())
	}

	// Start broadcast
	go func() {
		defer close(output)
		for n, line := range strings.Split(input, "\n") {
			output <- &Line{n, line}
		}
	}()

	wg.Wait()
}
