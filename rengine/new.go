// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// new.go                                                                       \\
// create a new rengine instance.                                               \\
// ---------------------------------------------------------------------------- \\

package rengine

// New : Create a new REngine instance.
func New() REngine {
	return REngine{Initialized: false}
}