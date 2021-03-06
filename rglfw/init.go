// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// init.go                                                                      \\
// initialize the rglfw instance.                                               \\
// ---------------------------------------------------------------------------- \\

package rglfw

import (
	"runtime"
)

// init : Initialize.
func init() {

	runtime.LockOSThread()	// Lock GLFW to the main thread.

	GLFW = New()			// Create the new RGLFW instance.
	
}