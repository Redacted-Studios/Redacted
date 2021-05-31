// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// rengine.go                                                                   \\
// manage the redacted engine core.                                             \\
// ---------------------------------------------------------------------------- \\

package rengine

// Engine : REngine instance.
var Engine REngine

// REngine : Redacted Engine Core.
type REngine struct {
	Version 	string 	// Version : REngine version used.
	Initialized bool	// initialized : REngine's initialization state.
}
