// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// rwindow.go                                                                   \\
// manage the redacted engines glfw window.                                     \\
// ---------------------------------------------------------------------------- \\

package rwindow

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Window : RWindow Instance.
var Window RWindow

// RWindow : Redacted Engine Window instance.
type RWindow struct {
	width 		int 			// width : Window's width in pixels.
	height 		int 			// height : Window's height in pixels.
	caption 	string 			// caption : Window's text caption.
	Initialized bool			// initialized : Window's initialization state.
	Window      *glfw.Window 	// window : glfw window instance.
}
