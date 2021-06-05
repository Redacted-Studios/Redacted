// ---------- Copyright 2021, Redacted Studios, All rights reserved. ---------- \\
// rpython.go                                                                   \\
// manage the redacted engines python 3 integration.                            \\
// ---------------------------------------------------------------------------- \\

package rpython

//
// #cgo windows CFLAGS: -I "python/windows/include"
// #cgo windows LDFLAGS: -L "python/windows/libs" -l python3 -l python39
// #cgo linux CFLAGS: -I "python/linux/include"
// #cgo linux LDFLAGS: -L "python/linux/libs" -l python3 -l python39
//
// #include "Python.h"
//
// typedef int (*intFunc) ();
//
// int bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int python() 
// {
//		Py_Initialize();
//		PyRun_SimpleString("import sys\nprint('Hello Python 3!')\nprint(sys.version)");
//		Py_FinalizeEx();
//		return 0;
// }
//
import "C"
import "fmt"

// Python : RPython Instance.
var Python RPython

// RPython : Redacted Engine Python 3 Instance.
type RPython struct {
	Version 	string 	// Version : RPython version used.
	Initialized bool	// Initialized : RPython's initialization state.
}

func Test() {
	f := C.intFunc(C.python)
	fmt.Println("C Print of 55 result:", int(C.bridge_int_func(f)))
}
