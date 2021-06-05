// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// open.go                                                                      \\
// initialize the redacted engine python 3 instance.                            \\
// ---------------------------------------------------------------------------- \\

package rpython

// Open : Initialize the RPython instance.
func Open() error {

	// err : Error variable.
	// var err error

	// Initialize Python.
	// err = ropengl.Open()
	// if err != nil {return err}

	// Set the RPython Version.
	Python.Version = "Development Build, Pre-Alpha 0.0.0.1"

	// Set the RPython initialization state to true.
	Python.Initialized = true

	// No error occurred.
	return nil

}
