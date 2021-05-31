// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// open.go                                                                      \\
// initialize the rglfw instance.                                               \\
// ---------------------------------------------------------------------------- \\

package rglfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Open : Initialize the RGLFW instance.
func Open() error {

	// err : Error variable.
	var err error

	// Initialize GLFW.
	err = glfw.Init()

	// If error occurred, return err.
	if err != nil {return err}

	// Set the RGLFW's version.
	// GLFW.Version = gl.GoStr(gl.GetString(gl.VERSION))
	GLFW.Version = glfw.GetVersionString()

	// Set the RGLFW's initialization state to true.
	GLFW.Initialized = true

	// No error occurred.
	return nil

}
