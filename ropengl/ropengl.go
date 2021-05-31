// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// ropengl.go                                                                   \\
// manage opengl for the redacted engine.                                       \\
// ---------------------------------------------------------------------------- \\

package ropengl

// OpenGL : ROpenGL Instance.
var OpenGL ROpenGL

// ROpenGL : Redacted Engine OpenGL Information.
type ROpenGL struct {
	Version 	string 	// Version : OpenGL version used.
	Initialized bool	// initialized : OpenGL's initialization state.
}
