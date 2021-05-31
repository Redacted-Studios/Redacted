// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// new.go                                                                       \\
// create a new rglfw instance.                                                 \\
// ---------------------------------------------------------------------------- \\

package rglfw

// New : Create a new RGLFW instance.
func New() RGLFW {

	return RGLFW{Initialized: false}

}
