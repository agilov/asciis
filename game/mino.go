package game

type Tile struct {
	Char byte
}

// Mino is a thing that is moving around and combines
type Mino struct {
	Position Position
	Box      Box
	level    *level
	Tiles    []Tile
	Rotation int
}

type Box struct {
	Width, Height int
}
type Position struct {
	X, Y int
}

func newMino(lvl *level) *Mino {
	m := &Mino{
		Rotation: 0,
		level:    lvl,
		Position: Position{0, 0},
		Box:      Box{4, 4},
		Tiles: []Tile{
			{'1'}, {'2'}, {'3'}, {'4'},
			{'1'}, {' '}, {' '}, {'4'},
			{'1'}, {'2'}, {'3'}, {'4'},
			{'1'}, {'2'}, {'3'}, {'4'},
		},
	}
	go m.sink()
	return m
}

func (m *Mino) sink() {
	m.Position.Y++
	if m.Position.Y+m.Box.Height >= m.level.height {
		m.Position.Y--
	}
}

func (m *Mino) Rotate(right bool) {
	if right {
		m.Rotation++
	} else {
		m.Rotation--
	}
	if m.Rotation > 3 {
		m.Rotation = 0
	}
	if m.Rotation < 0 {
		m.Rotation = 3
	}
	if right {
		m.rotateRight()
	} else {
		m.rotateLeft()
	}
}

func (m *Mino) rotateLeft() {
	h := m.Box.Height
	w := m.Box.Width
	toH := w
	toW := h
	result := make([]Tile, len(m.Tiles))
	for i := 0; i < len(result); i++ {
		colPos := i / w
		colNum := i % w
		rowToPos := colPos
		rowToNum := toH - colNum - 1
		j := toW*rowToNum + rowToPos
		result[j] = m.Tiles[i]
	}
	m.Tiles = result
	m.Box.Height = toH
	m.Box.Width = toW
}

func (m *Mino) rotateRight() {
	h := m.Box.Height
	w := m.Box.Width
	toH := w
	toW := h
	//col, row := 1, h
	result := make([]Tile, len(m.Tiles))
	for i := 0; i < len(result); i++ {
		rowNum := i / w
		rowPos := i % w
		colToPos := rowPos
		colToNum := toW - rowNum - 1
		j := toW*colToPos + colToNum
		result[j] = m.Tiles[i]
	}
	m.Tiles = result
	m.Box.Height = toH
	m.Box.Width = toW
}

func (m *Mino) render() {
	data := m.Tiles
	for i, x, y := 0, m.Position.X, m.Position.Y; ; i++ {
		m.level.data[y][x] = data[i].Char
		x++
		if x >= m.Position.X+m.Box.Width {
			x = m.Position.X
			y++
		}
		if y >= m.Position.Y+m.Box.Height {
			return
		}
	}
}

func (m *Mino) clear() {
	for x, y := m.Position.X, m.Position.Y; ; {
		m.level.data[y][x] = NOTHING
		x++
		if x >= m.Position.X+m.Box.Width {
			x = m.Position.X
			y++
		}
		if y >= m.Position.Y+m.Box.Height {
			return
		}
	}
}

func (m *Mino) update(event event) {
	if event.Name() == "keyPressed" {
		if event.Key() == 'd' {
			m.Position.X++
		}
		if event.Key() == 'a' {
			m.Position.X--
		}
		if event.Key() == 'w' {
			m.Position.Y--
		}
		if event.Key() == 's' {
			m.Position.Y++
		}
		if event.Key() == arrowLeft {
			m.Position.X--
		}
		if event.Key() == arrowRight {
			m.Position.X++
		}
		if event.Key() == arrowUp {
			m.Rotate(false)
		}
		if event.Key() == arrowDown {
			m.Rotate(true)
		}
		if event.Key() == '\n' {
			m.Position.Y = m.level.height - m.Box.Height - 1
		}
	}
	if m.Position.X <= 0 {
		m.Position.X++
	}
	if m.Position.Y <= 0 {
		m.Position.Y++
	}
	if m.Position.X+m.Box.Width >= m.level.width {
		m.Position.X--
	}
	if m.Position.Y+m.Box.Height >= m.level.height {
		m.Position.Y--
	}
}
