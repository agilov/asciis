package game

// mino is a thing that is moving around and combines
type mino struct {
	pos        position
	box        box
	level      *level
	data       []byte
	renderData []byte
	rotation   int
}

type box struct {
	width, height int
}
type position struct {
	x, y int
}

func newMino(lvl *level) *mino {
	m := &mino{
		rotation:   0,
		level:      lvl,
		pos:        position{4, 4},
		box:        box{3, 2},
		data:       []byte{'1', '2', '3', '4'},
		renderData: []byte{'a', 'b', 'C', 'd', ' ', ' '},
	}
	m.writeRenderData()
	go m.sink()
	return m
}

func (m *mino) sink() {
	m.pos.y++
	if m.pos.y+m.box.height >= m.level.height {
		m.pos.y--
	}
}

func (m *mino) rotate(right bool) {
	if right {
		m.rotation++
	} else {
		m.rotation--
	}
	if m.rotation > 3 {
		m.rotation = 0
	}
	if m.rotation < 0 {
		m.rotation = 3
	}
	m.writeRenderData()
	m.box = box{m.box.height, m.box.width}
	if m.pos.x+m.box.width >= m.level.width-1 {
		m.pos.x--
	}
}

func (m *mino) writeRenderData() {
	a, b, c, d := m.data[0], m.data[1], m.data[2], m.data[3]
	switch m.rotation {
	case 0:
		m.renderData[0] = a
		m.renderData[1] = b
		m.renderData[2] = ' '
		m.renderData[3] = ' '
		m.renderData[4] = c
		m.renderData[5] = d
	case 1:
		m.renderData[0] = ' '
		m.renderData[1] = a
		m.renderData[2] = c
		m.renderData[3] = b
		m.renderData[4] = d
		m.renderData[5] = ' '
	case 2:
		m.renderData[0] = d
		m.renderData[1] = c
		m.renderData[2] = ' '
		m.renderData[3] = ' '
		m.renderData[4] = b
		m.renderData[5] = a
	case 3:
		m.renderData[0] = ' '
		m.renderData[1] = d
		m.renderData[2] = b
		m.renderData[3] = c
		m.renderData[4] = a
		m.renderData[5] = ' '
	}
}

func (m *mino) render() {
	data := m.renderData
	for i, x, y := 0, m.pos.x, m.pos.y; ; i++ {
		m.level.data[y][x] = data[i]
		x++
		if x >= m.pos.x+m.box.width {
			x = m.pos.x
			y++
		}
		if y >= m.pos.y+m.box.height {
			return
		}
	}
}

func (m *mino) clear() {
	for x, y := m.pos.x, m.pos.y; ; {
		m.level.data[y][x] = NOTHING
		x++
		if x >= m.pos.x+m.box.width {
			x = m.pos.x
			y++
		}
		if y >= m.pos.y+m.box.height {
			return
		}
	}
}

func (m *mino) update(event event) {
	if event.Name() == "keyPressed" {
		if event.Key() == 'd' {
			m.pos.x++
		}
		if event.Key() == 'a' {
			m.pos.x--
		}
		if event.Key() == 'w' {
			m.pos.y--
		}
		if event.Key() == 's' {
			m.pos.y++
		}
		if event.Key() == arrowLeft {
			m.pos.x--
		}
		if event.Key() == arrowRight {
			m.pos.x++
		}
		if event.Key() == arrowUp {
			m.rotate(false)
		}
		if event.Key() == arrowDown {
			m.rotate(true)
		}
		if event.Key() == '\n' {
			m.pos.y = m.level.height - m.box.height - 1
		}
	}
	if m.pos.x <= 0 {
		m.pos.x++
	}
	if m.pos.y <= 0 {
		m.pos.y++
	}
	if m.pos.x+m.box.width >= m.level.width {
		m.pos.x--
	}
	if m.pos.y+m.box.height >= m.level.height {
		m.pos.y--
	}
}
