// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// main.go                                                                      \\
// run the application.                                                         \\
// ---------------------------------------------------------------------------- \\

package main

import (
	"os"
	"fmt"
	
	"redacted/rglfw"
	"redacted/rwindow"
	"redacted/rengine"
	"redacted/ropengl"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.5-core/gl"
)

func main() {

	// err : Error variable.
	var err error

	// Initialize the Redacted Engine.
	err = rengine.Open()
	if err != nil {panic(err)}

	// Log version information.
	fmt.Println("Redacted Engine Version:", rengine.Engine.Version)
	fmt.Println("OpenGL Version:", ropengl.OpenGL.Version)
	fmt.Println("GLFW Version:", rglfw.GLFW.Version)

	// Application main loop.
	for !rwindow.Window.Window.ShouldClose() {

		gl.ClearColor(0, 0.5, 1.0, 1.0)

		// Clear the GL_COLOR_BUFFER_BIT & GL_DEPTH_BUFFER_BIT.
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update the window.
		rwindow.Window.Window.SwapBuffers()
		glfw.PollEvents()

	}

	// Terminate the Redacted Engine.
	err = rengine.Close()
	if err != nil {panic(err)}

	// Exit the application.
	os.Exit(0)

}
