package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
	_ "github.com/go-vgo/robotgo/base"
	_ "github.com/go-vgo/robotgo/key"
	_ "github.com/go-vgo/robotgo/mouse"
	_ "github.com/go-vgo/robotgo/screen"
	_ "github.com/go-vgo/robotgo/window"
)

var appVersion = "0.1.1" // see FyneApp.toml
var appAuthor = "Allan Marillier"
var appName = "jiggler"
var appCopyright = "Copyright (c) Allan Marillier, 2025-" + strconv.Itoa(time.Now().Year())

func main() {
	var jiggleInt int = 10 // default jiggle every 10 minutes
	var pixelMove int = 1  // default jiggle 1 pixel to be as fast as possible
	var jiggle int
	var pixels int
	var checkupdate bool

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("jiggler is a simple utility to prevent screensaver lockout by moving the mouse cursor")
		fmt.Println(appName, "version", appVersion, "-", appCopyright)
		fmt.Println("Example: jiggler -jiggle 10 -pixels 1")
		fmt.Println("Example: jiggler -j 10 -p 1")
		fmt.Println("To check for updates, use -checkupdate or -cu")
		fmt.Println("Example: jiggler -checkupdate")
		fmt.Println("Example: jiggler -cu")
		fmt.Println("NOTE: This will not auto update, this is a manual check, and a manual\nupdate of a portable app")
		fmt.Println("jiggle defaults are 10 minutes and 1 pixel")
		fmt.Println("For fastest jiggle performance, use -pixels 1")
		flag.PrintDefaults()
		fmt.Println("Use OS specific techniques to pass parameters and run as a background process")
	}

	// Define the command-line flags for jiggle interval and pixel movement
	// jiggleInterval := flag.Int("jiggle", 0, "Jiggle interval in minutes (default: 10 minutes)")
	// pixelMovement := flag.Int("pixels", pixelMove, "Number of pixels to move the mouse (default: 1 pixel)")
	flag.IntVar(&jiggle, "jiggle", 0, "Jiggle interval in minutes (default: 10 minutes)")
	flag.IntVar(&jiggle, "j", 0, "Jiggle interval in minutes (default: 10 minutes)")
	flag.IntVar(&pixels, "pixels", pixelMove, "Number of pixels to move the mouse (default: 1 pixel)")
	flag.IntVar(&pixels, "p", pixelMove, "Number of pixels to move the mouse (default: 1 pixel)")
	flag.BoolVar(&checkupdate, "checkupdate", false, "Check for updates (default: 0, no check)")
	flag.BoolVar(&checkupdate, "cu", false, "Check for updates (default: 0, no check)")
	flag.Parse()

	// Check if help flag is set
	if flag.NFlag() == 0 || flag.Lookup("help") != nil {
		flag.Usage()
		os.Exit(0)
	}

	if checkupdate {
		// Check for updates
		fmt.Println("Checking for updates...")
		updtmsg, updateAvail := updateChecker("amarillier", "jiggler", "jiggler", "https://github.com/amarillier/jiggler/releases/latest")
		fmt.Println(updtmsg)
		if updateAvail {
			fmt.Println("Update available:", updtmsg)
			fmt.Println("Please visit the GitHub releases page to download the latest version.")
		}
		os.Exit(0)
	}

	if jiggle == 0 {
		jiggle = jiggleInt
	}
	if jiggle < 1 {
		jiggle = 1
	} else if jiggle > 60 {
		jiggle = 60 // limit 60 minutes
	}
	if pixels < 1 {
		pixels = 1
	} else if pixels > 100 {
		pixels = 100 // limit 100 pixels
	}
	jiggleInt = jiggle
	pixelMove = pixels

	for {
		robotgo.MoveRelative(pixelMove, 0)  // MoveSmoothRelative(200, 0)
		robotgo.MoveRelative(0, pixelMove)  // MoveSmoothRelative(0, 200)
		robotgo.MoveRelative(-pixelMove, 0) // MoveSmoothRelative(-200, 0)
		robotgo.MoveRelative(0, -pixelMove) // MoveSmoothRelative(0, -200)
		// robotgo.Sleep(jiggleInt * 60 * 1000)
		time.Sleep(time.Duration(jiggleInt) * time.Minute)
	}
}

// "Now this is not the end. It is not even the beginning of the end. But it is, perhaps, the end of the beginning." Winston Churchill, November 10, 1942
