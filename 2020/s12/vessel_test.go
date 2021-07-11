package s12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVesselRotate(t *testing.T) {
	tests := []struct {
		facing         CompassDirection
		cmds           []Instruction
		expectedFacing CompassDirection
	}{
		{east, []Instruction{{rotLeft, 180}}, west},
		{west, []Instruction{{rotLeft, 180}}, east},
		{south, []Instruction{{rotLeft, 180}, {rotRight, 270}}, west},
	}

	for _, test := range tests {
		v := &Vessel{facing: test.facing}

		for _, i := range test.cmds {
			v.Rotate(i.cmd, i.units)
		}

		assert.Equal(t, test.expectedFacing, v.facing)
	}
}
