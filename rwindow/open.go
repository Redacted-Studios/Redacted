// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// open.go                                                                      \\
// open the rwindow instance.                                                   \\
// ---------------------------------------------------------------------------- \\

package rwindow

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Open : Initialize the RWindow Instance.
func (window *RWindow) Open() error {

	// err : Error variable.
	var err error

	// If GLFW has not been initialized, initialize it.
	err = glfw.Init()

	// If error occurred, return error.
	if err != nil {
		return err
	}

	// Set the window initialization state to true.
	window.initialized = true

	// No error occurred.
	return nil

}
