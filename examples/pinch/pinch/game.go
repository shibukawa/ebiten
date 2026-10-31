// Copyright 2026 The Ebitengine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package pinch is a game that scales a square with pinch gestures.
// It is shared by the desktop and mobile builds of the example.
package pinch

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480

	minScale = 0.1
	maxScale = 10

	// gridPitch is the distance between the grid lines in the world.
	gridPitch = 80
)

var (
	gridColor   = color.RGBA{0x40, 0x40, 0x40, 0xff}
	axisColor   = color.RGBA{0x80, 0x80, 0x80, 0xff}
	squareColor = color.RGBA{0xff, 0, 0, 0xff}
)

// Game is a game that scales a square with pinch gestures.
type Game struct {
	// A point p in the world is drawn at p*scale + offset.
	scale   float64
	offsetX float64
	offsetY float64

	delta float64

	touchIDs []ebiten.TouchID
}

// NewGame returns a new game.
func NewGame() *Game {
	return &Game{
		scale:   1,
		offsetX: ScreenWidth / 2,
		offsetY: ScreenHeight / 2,
	}
}

// zoomCenter returns the point that a pinch zooms around: the middle of the fingers when two fingers touch
// the screen, and the cursor position otherwise.
func (g *Game) zoomCenter() (float64, float64) {
	g.touchIDs = ebiten.AppendTouchIDs(g.touchIDs[:0])
	if len(g.touchIDs) == 2 {
		x0, y0 := ebiten.TouchPositionF(g.touchIDs[0])
		x1, y1 := ebiten.TouchPositionF(g.touchIDs[1])
		return (x0 + x1) / 2, (y0 + y1) / 2
	}
	return ebiten.CursorPositionF()
}

func (g *Game) Update() error {
	// PinchDelta is the change of the scale the gesture asks for at this tick.
	g.delta = ebiten.PinchDelta()
	if g.delta != 0 {
		// Zoom around the center, so that the point under the center stays still.
		cx, cy := g.zoomCenter()
		scale := min(max(g.scale*(1+g.delta), minScale), maxScale)
		r := scale / g.scale
		g.offsetX = cx - (cx-g.offsetX)*r
		g.offsetY = cy - (cy-g.offsetY)*r
		g.scale = scale
	}

	// Pan with scrolling, such as with two fingers on a touchpad.
	dx, dy := ebiten.ScrollDelta()
	g.offsetX += dx
	g.offsetY += dy
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// The grid shows how the world is scaled and moved.
	step := gridPitch * g.scale
	for x := math.Mod(g.offsetX, step); x < ScreenWidth; x += step {
		vector.StrokeLine(screen, float32(x), 0, float32(x), ScreenHeight, 1, gridColor, false)
	}
	for y := math.Mod(g.offsetY, step); y < ScreenHeight; y += step {
		vector.StrokeLine(screen, 0, float32(y), ScreenWidth, float32(y), 1, gridColor, false)
	}
	vector.StrokeLine(screen, float32(g.offsetX), 0, float32(g.offsetX), ScreenHeight, 1, axisColor, false)
	vector.StrokeLine(screen, 0, float32(g.offsetY), ScreenWidth, float32(g.offsetY), 1, axisColor, false)

	// The square is centered at the origin of the world.
	s := float32(gridPitch * g.scale)
	vector.FillRect(screen, float32(g.offsetX)-s/2, float32(g.offsetY)-s/2, s, s, squareColor, true)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("Pinch on a touchpad or a touch screen to scale the red square\nScroll to move it\nScale: %0.3f\nPinchDelta: %+0.4f", g.scale, g.delta))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
