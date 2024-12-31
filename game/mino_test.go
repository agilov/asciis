package game_test

import (
	. "github.com/agilov/notagame/game"
	"github.com/stretchr/testify/assert"
	"testing"
)

type rotateMatrixTestCase []struct {
	right    bool
	input    Mino
	expected Mino
}

// go test -v -run TestRotate ./game/*.go
func TestRotate(t *testing.T) {
	css := []string{
		`l32
ab 
 cd
 d
bc
a `,
		`r23
 d
bc
a 
ab 
 cd`,
		`l41
abcd
d
c
b
a`,
		`l23
a 
bc
d 
 c 
abd`,
		`r32
 c 
abd
a 
bc
d `,
		`r32
 ab
cd 
c 
da
 b`,
		`r44
abcd
e  f
z  k
gh12
gzea
h  b
1  c
2kfd`,
		`l23
 a
bc
 d
acd
 b `,
	}

	createCase := func(c string) (right bool, mino, expected Mino) {
		right = c[0] == 'r'
		w, h := int(c[1]-'0'), int(c[2]-'0')
		dataLen := w*h + h
		minoData := c[4 : 4+dataLen]
		expectData := c[4+dataLen:]
		mino.Box.Width, mino.Box.Height = w, h
		for i := 0; i < len(minoData); i++ {
			if minoData[i] == '\n' {
				continue
			}
			mino.Tiles = append(mino.Tiles, Tile{Char: minoData[i]})
		}
		expected.Box.Width, expected.Box.Height = h, w
		for i := 0; i < len(expectData); i++ {
			if expectData[i] == '\n' {
				continue
			}
			expected.Tiles = append(expected.Tiles, Tile{Char: expectData[i]})
		}
		return
	}
	for _, tc := range css {
		right, input, expect := createCase(tc)
		input.Rotate(right)
		assert.Equal(t, expect.Position, input.Position, "Position should not change")
		assert.Equal(t, expect.Box.Height, input.Box.Height, "Should switch height")
		assert.Equal(t, expect.Box.Width, input.Box.Width, "Should switch width")
		assert.Equal(t, expect.Tiles, input.Tiles, "Should transform tiles matrix")
	}
}
