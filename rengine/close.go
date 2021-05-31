// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// close.go                                                                     \\
// close the rengine instance.                                                  \\
// ---------------------------------------------------------------------------- \\

package rengine

import (
	"redacted/rglfw"
	"redacted/ropengl"
	"redacted/rwindow"
)

// Close : Terminate the REngine instance.
func Close() error {

	// err : Error variable.
	var err error

	// Close the RWindow instance.
	err = rwindow.Close()
	if err != nil {return err}

	// Close the ROpenGL instance.
	err = ropengl.Close()
	if err != nil {return err}

	// Close the RGLFW instance.
	err = rglfw.Close()
	if err != nil {return err}

	// Set the REngine initialization state to false.
	Engine.Initialized = false
	
	// No error occurred.
	return nil

}