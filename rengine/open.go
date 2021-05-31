// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// rengine.go                                                                   \\
// manage the redacted engine core.                                             \\
// ---------------------------------------------------------------------------- \\

package rengine

import (
	"fmt"

	"redacted/rglfw"
	"redacted/ropengl"
	"redacted/rwindow"
)

// Open : Initialize the REngine instance.
func Open() error {

	// err : Error variable.
	var err error

	// Initialize RGLFW.
	err = rglfw.Open()
	if err != nil {return err}

	// Initialize RWindow.
	err = rwindow.Open()
	if err != nil {return err}

	// Initialize ROpenGL.
	err = ropengl.Open()
	if err != nil {return err}

	// Set the Redacted Engine Version.
	Engine.Version = "Development Build, Pre-Alpha 0.0.0.1"

	// Set the REngine initialization state to true.
	Engine.Initialized = true

	// TODO : Remove.
	fmt.Println("Redacted Engine Successfully Initialized.")

	// No error occurred.
	return nil

}
