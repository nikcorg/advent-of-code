package packet

import (
	"fmt"
	"strconv"
)

func ParseMessage(msg string) Packet {
	if msg[0] != '[' {
		panic(fmt.Errorf("malformed message: %s", msg))
	}

	p, _ := parsePacket(msg[1:], 0) // skip the opening bracket

	return p
}

func parsePacket(line string, offset int) (Packet, int) {
	var (
		packet = ListPacket{values: []Packet{}}
		token  string
	)

	for i := offset; i < len(line); i++ {
		switch line[i] {
		case '[':
			c, end := parsePacket(line, i+1)
			packet.Append(c)
			i = end

		case ']':
			if token != "" {
				packet.Append(NewIntPacket(mustParseInt(token)))
			}
			return packet, i

		case ',':
			if token != "" {
				packet.Append(NewIntPacket(mustParseInt(token)))
				token = ""
			}

		default:
			token += string(line[i])
		}
	}

	panic("unbalanced brackets")
}

func mustParseInt(x string) int {
	n, _ := strconv.Atoi(x)
	return n
}
