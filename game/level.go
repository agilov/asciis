package game

import (
	"time"
)

type level struct {
	width, height int
	data          [][]byte
	lastSink      time.Time
}

func newLevel(width, height int) *level {
	data := make([][]byte, height)
	for h := 0; h < height; h++ {
		data[h] = make([]byte, width)
		for w := 0; w < width; w++ {
			if w == width-1 || w == 0 || h == 0 || h == height-1 {
				data[h][w] = WALL
			} else {
				data[h][w] = NOTHING
			}
		}
	}

	return &level{width: width, height: height, data: data, lastSink: time.Now()}
}

func (l *level) shouldSink() bool {
	if time.Since(l.lastSink).Seconds() > 1 {
		l.lastSink = time.Now()
		return true
	}
	return false
}

func (l *level) x() {

}

func newGame(width, height, fps int) *game {
	return &game{
		level:        newLevel(width, height),
		isRunning:    false,
		stats:        newStats(),
		maxFrameTime: time.Second / time.Duration(fps),
	}
}
