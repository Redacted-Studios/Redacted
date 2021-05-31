// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// new.go                                                                       \\
// create a new rwindow instance.                                               \\
// ---------------------------------------------------------------------------- \\

package rwindow

// New : Create a new RWindow instance.
func New() RWindow {

	return RWindow{initialized: false}

}
