package game

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

const (
	NOTHING     = 0
	WALL        = 1
	PLAYER      = 69
	BRICK       = "█"
	MAX_SAMPLES = 100
)

type game struct {
	isRunning    bool
	level        *level
	drawBuffer   bytes.Buffer
	stats        *stats
	maxFrameTime time.Duration
	input        *input
	currentMino  *Mino
	nextMino     *Mino
}

func NewGame(width, height, fps int) *game {
	lvl := newLevel(width, height)
	return &game{
		level:        lvl,
		isRunning:    false,
		stats:        newStats(),
		maxFrameTime: time.Second / time.Duration(fps),
		input:        newInput(),
		currentMino:  newMino(lvl),
		nextMino:     newMino(lvl),
	}
}

func (g *game) Start() {
	fmt.Fprint(os.Stdout, "\033]50;Free Mono\007")
	terminalRaw()
	defer terminalRestore()

	g.isRunning = true
	g.loop()
}

func (g *game) loop() {
	for g.isRunning {
		start := time.Now()
		event := g.input.read()
		g.stats.update()
		g.update(event)
		g.render()
		dt := g.maxFrameTime - time.Since(start)
		if dt > 0 {
			time.Sleep(dt)
		}
	}
}

func (g *game) renderLevel() {
	g.drawBuffer.Reset()
	for h := 0; h < g.level.height; h++ {
		for w := 0; w < g.level.width; w++ {
			if g.level.data[h][w] == WALL {
				g.drawBuffer.WriteString(BRICK)
				g.drawBuffer.WriteString(BRICK)
			} else if g.level.data[h][w] == NOTHING {
				g.drawBuffer.WriteString("  ")
			} else {
				g.drawBuffer.WriteString(string(g.level.data[h][w]))
				g.drawBuffer.WriteString(" ")
			}
		}
		g.drawBuffer.WriteString("\n")
	}
}

func (g *game) update(event event) {
	g.currentMino.clear()
	g.currentMino.update(event)
	if g.level.shouldSink() {
		g.currentMino.sink()
	}
}

func (g *game) render() {
	clearTerminal()
	g.renderMino()
	g.renderLevel()
	g.renderStats()
	fmt.Fprint(os.Stdout, g.drawBuffer.String())
}

func (g *game) renderMino() {
	g.currentMino.render()
}

func clearTerminal() {
	fmt.Fprint(os.Stdout, "\033[2J\033[1:1H")

}

func (g *game) renderStats() {
	g.drawBuffer.WriteString("-- STATS\n")
	g.drawBuffer.WriteString(fmt.Sprintf("-- FPS: %.2f\n", g.stats.fps))
}
