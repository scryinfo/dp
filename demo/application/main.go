package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilog"
)

// Constants
const ()

// Vars
var ()

func main() {
	// Initialize astilectron
	a, err := astilectron.New(astilectron.Options{
		AppName:            "<your app name>",
		AppIconDefaultPath: "<your .png icon>",  // If path is relative, it must be relative to the data directory
		AppIconDarwinPath:  "<your .icns icon>", // Same here
		BaseDirectoryPath:  "<where you want the provisioner to install the dependencies>",
	})
	defer a.Close()
	if err != nil {
		astilog.Fatal(err)
	}

	// Start astilectron
	err = a.Start()
	if err != nil {
		astilog.Fatal(err)
	}

	// Blocking pattern
	a.Wait()


	// Create a new window
	w, err := a.NewWindow("http://127.0.0.1:4000", &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(600),
	})
	if err != nil {
		astilog.Fatal(err)
	}
	err = w.Create()
	if err != nil {
		astilog.Fatal(err)
	}

	// Open dev tools
	err = w.OpenDevTools()
	if err != nil {
		astilog.Fatal(err)
	}

	// Close dev tools
	err = w.CloseDevTools()
	if err != nil {
		astilog.Fatal(err)
	}
}
