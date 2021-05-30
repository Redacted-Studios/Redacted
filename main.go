// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// main.go                                                                      \\
// run the application.                                                         \\
// ---------------------------------------------------------------------------- \\

package main

import (
	"fmt"
	"github.com/redacted-studios/redacted/window"

	_ "github.com/go-gl/gl/v4.5-core/gl"
	_ "github.com/go-gl/glfw/v3.3/glfw"
)

func main() {

	win := window.RWindow{Caption: "Window - Modules"}

	fmt.Println("Hello Redacted Engine:", win.Caption)
}
