package packet

import "sort"

type PacketSlice []Packet

var _ sort.Interface = PacketSlice{}

func (p PacketSlice) Less(a, b int) bool {
	return p[a].Compare(p[b]) == GreaterThan
}

func (p PacketSlice) Len() int {
	return len(p)
}

func (p PacketSlice) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}
