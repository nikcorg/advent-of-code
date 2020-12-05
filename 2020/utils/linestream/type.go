package linestream

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
)

const newline = '\n'

type Line struct {
	num     int
	content string
}

func (l Line) Num() int {
	return l.num
}

func (l Line) RawContent() string {
	return l.content
}

func (l Line) Content() string {
	return strings.TrimSpace(l.content)
}

type LineChan = chan *Line

// New creates a new LineStream reader
func New(ctx context.Context, reader *bufio.Reader) LineChan {
	out := make(chan *Line)
	linesRead := 0

	go func() {
		defer close(out)

		var (
			text string
			err  error
		)

		for {
			text, err = reader.ReadString(newline)
			linesRead++

			if err != nil && err != io.EOF {
				panic(fmt.Errorf("error: %w", err))
			}

			select {
			case <-ctx.Done():
				return
			case out <- &Line{linesRead, text}:
				if err != nil && err == io.EOF {
					return
				}
			}
		}
	}()

	return out
}

func SkipEmpty(in LineChan) LineChan {
	out := make(chan *Line)

	go func() {
		defer close(out)

		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				if v == nil || v.Content() == "" {
					continue
				}

				out <- v
			}
		}
	}()

	return out
}

func WithDoneSignalling(done context.CancelFunc, in LineChan) LineChan {
	out := make(chan *Line)

	go func() {
		defer close(out)

		for {
			select {
			case v, ok := <-in:
				if !ok {
					done()
					return
				}
				out <- v
			}
		}
	}()

	return out
}
