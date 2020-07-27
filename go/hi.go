package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// main is the starting point for your Buffalo application.
// You can feel free and add to this `main` method, change
// what it does, etc...
// All we ask is that, at some point, you make sure to
// call `app.Serve()`, unless you don't want to start your
// application that is. :)
func main() {
	fmt.Println("asd")
	o := os.Args
	a := filepath.Dir(os.Args[0])
	fmt.Println(a, o)
}
