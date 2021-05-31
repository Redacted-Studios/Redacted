// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// main.go                                                                      \\
// run the application.                                                         \\
// ---------------------------------------------------------------------------- \\

package main

import (
	"os"
	"fmt"
	"redacted/rwindow"
	"redacted/ropengl"
	"redacted/rglfw"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.5-core/gl"
)

func main() {

	// err : Error variable.
	var err error

	// Init glfw.
	err = rglfw.Open()
	if err != nil {panic(err)}

	// Open a new glfw window.
	err = rwindow.Open()
	if err != nil {panic(err)}

	// Temp - Initialzize opengl.
	err = ropengl.Open()
	if err != nil {panic(err)}

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

	// Close the glfw window.
	err = rwindow.Close()
	if err != nil {panic(err)}

	// Exit the application.
	os.Exit(0)

}
