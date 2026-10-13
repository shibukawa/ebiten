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

package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480

	minScale = 0.1
	maxScale = 10
)

type Game struct {
	scale float64
	delta float64
}

func (g *Game) Update() error {
	// PinchDelta is the change of the scale the gesture asks for at this tick.
	g.delta = ebiten.PinchDelta()
	g.scale = min(max(g.scale*(1+g.delta), minScale), maxScale)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const size = 80
	s := float32(size * g.scale)
	vector.FillRect(screen, screenWidth/2-s/2, screenHeight/2-s/2, s, s, color.RGBA{0xff, 0, 0, 0xff}, false)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("Pinch on a touchpad to scale the red square\nScale: %0.3f\nPinchDelta: %+0.4f", g.scale, g.delta))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{scale: 1}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pinch (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
