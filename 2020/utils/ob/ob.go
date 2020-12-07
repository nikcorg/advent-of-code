package ob

type Capture struct {
	b [][]byte
}

func (ob *Capture) Write(bytes []byte) (int, error) {
	ob.b = append(ob.b, bytes)
	return len(bytes), nil
}

func (ob *Capture) Reset() {
	ob.b = make([][]byte, 0)
}

func (ob *Capture) String() string {
	out := ""

	for _, chunk := range ob.b {
		out += string(chunk)
	}

	return out
}
