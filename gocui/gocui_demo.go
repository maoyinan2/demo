/*
overlapping views.
The GUI can be modified at runtime (concurrent-safe).
Global and view-level keybindings.
Mouse support.
Colored text.
Customizable edition mode.
Easy to build reusable widgets, complex layouts...
*/

package main

import "demo/gocui/app"

var a *app.App
var cfg *app.Config

func main() {
	// s := initSearch()
	a = app.New()

	defer a.Close()

	a.Loop()
}
