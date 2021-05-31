// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// main.go                                                                      \\
// run the application.                                                         \\
// ---------------------------------------------------------------------------- \\

package main

import (
	"os"
	"redacted/rwindow"
)

func main() {

	// err : Error variable.
	var err error

	// Open a new glfw window.
	err = rwindow.Window.Open()
	if err != nil {panic(err)}

	// Close the glfw window.
	err = rwindow.Window.Close()
	if err != nil {panic(err)}

	// Exit the application.
	os.Exit(0)

}
