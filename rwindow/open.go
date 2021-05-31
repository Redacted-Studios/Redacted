// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// open.go                                                                      \\
// open the rwindow instance.                                                   \\
// ---------------------------------------------------------------------------- \\

package rwindow

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Open : Initialize the RWindow Instance.
func Open() error {

	// err : Error variable.
	var err error

	// TODO : use rglfw.

	// If GLFW has not been initialized, initialize it.
	err = glfw.Init()

	// If error occurred, return error.
	if err != nil {return err}

	// Set the GLFW window flags.
	glfw.WindowHint(glfw.ContextVersionMajor, 4)         		// Maximum OpenGL version.
	glfw.WindowHint(glfw.ContextVersionMinor, 3)         		// Minimum OpenGL version.
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)	// Use the OpenGL core profile.

	// Open a new GLFW window.
	Window.Window, err = glfw.CreateWindow(1280, 720, "Redacted Engine", nil, nil)

	// If error occurred, return error.
	if err != nil {return err}

	// Make the Windows context current for OpenGL.
	Window.Window.MakeContextCurrent()

	// Set the window initialization state to true.
	Window.Initialized = true

	// No error occurred.
	return nil

}
