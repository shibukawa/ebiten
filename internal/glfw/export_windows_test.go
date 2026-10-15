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

package glfw

import (
	"unsafe"
)

func PackPoint64(x, y int32) uint64 {
	return (_POINT{x: x, y: y}).pack64()
}

const WheelPageScroll = _WHEEL_PAGESCROLL

func WheelScrollAmount(notches float64, setting uint32) (float64, ScrollUnit) {
	return wheelScrollAmount(notches, setting)
}

var SetFocusForTesting = _SetFocus

type DirectManipulationGesture struct {
	g directManipulationGesture
}

func NewDirectManipulationGesture() *DirectManipulationGesture {
	return &DirectManipulationGesture{
		g: newDirectManipulationGesture(),
	}
}

func (d *DirectManipulationGesture) Transform(m [6]float32) (panX, panY, pinch float64) {
	return d.g.transform(m)
}

// DirectManipulationVtblSlots returns the number of methods in each vtable, which must match
// directmanipulation.idl.
func DirectManipulationVtblSlots() map[string]uintptr {
	ptr := unsafe.Sizeof(uintptr(0))
	return map[string]uintptr{
		"IDirectManipulationManager":       unsafe.Sizeof(_IDirectManipulationManager_Vtbl{}) / ptr,
		"IDirectManipulationUpdateManager": unsafe.Sizeof(_IDirectManipulationUpdateManager_Vtbl{}) / ptr,
		"IDirectManipulationViewport":      unsafe.Sizeof(_IDirectManipulationViewport_Vtbl{}) / ptr,
		"IDirectManipulationContent":       unsafe.Sizeof(_IDirectManipulationContent_Vtbl{}) / ptr,
	}
}
