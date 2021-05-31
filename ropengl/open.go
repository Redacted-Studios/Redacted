// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// open.go                                                                      \\
// initialize the ropengl instance.                                             \\
// ---------------------------------------------------------------------------- \\

package ropengl

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

// Open : Initialize the ROpenGL instance.
func Open() error {

	// err : Error variable.
	var err error

	// Initialize OpenGL.
	err = gl.Init()

	// If error occurred, return err.
	if err != nil {return err}

	// Set the ROpenGL's version.
	OpenGL.Version = gl.GoStr(gl.GetString(gl.VERSION))

	// Set the ROpenGL's initialization state to true.
	OpenGL.Initialized = true

	// No error occurred.
	return nil

}
