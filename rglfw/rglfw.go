// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// rglfw.go                                                                     \\
// manage glfw for the redacted engine.                                         \\
// ---------------------------------------------------------------------------- \\

package rglfw

// OpenGL : ROpenGL Instance.
var GLFW RGLFW

// RGLFW : Redacted Engine GLFW Information.
type RGLFW struct {
	Version 	string 	// Version : GLFW version used.
	Initialized bool	// initialized : GLFW's initialization state.
}
