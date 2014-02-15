package main

// play with fonts: http://twinklebear.github.io/sdl2%20tutorials/2013/12/18/lesson-6-true-type-fonts-with-sdl_ttf/

import (
	"fmt"
	"github.com/jackyb/go-sdl2/sdl"
	"os"
)

var winTitle string = "Go-SDL2 Test"
var winWidth, winHeight int = 800, 600

const UPPERCASE = 1
const LOWERCASE = 0

func main() {

	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		fmt.Fprintf(os.Stderr, "Failed to init: %s\n", sdl.GetError())
		os.Exit(1)
	}
	defer sdl.Quit()

	window := sdl.CreateWindow(winTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight,
		sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}
	defer window.Destroy()

	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(1)
	}
	defer renderer.Destroy()

	var x int32 = 300
	var y int32 = 0

	for quit := false; !quit; {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyDownEvent:

				fmt.Printf(
					"[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)

				if sdl.GetKeyName(t.Keysym.Sym) == "Q" && t.Keysym.Mod == LOWERCASE {
					quit = true
				}

				if sdl.GetKeyName(t.Keysym.Sym) == "H" && t.Keysym.Mod == LOWERCASE {
					x -= 10
				}
				if sdl.GetKeyName(t.Keysym.Sym) == "L" && t.Keysym.Mod == LOWERCASE {
					x += 10
				}
				if sdl.GetKeyName(t.Keysym.Sym) == "J" && t.Keysym.Mod == LOWERCASE {
					y += 10
				}
				if sdl.GetKeyName(t.Keysym.Sym) == "K" && t.Keysym.Mod == LOWERCASE {
					y -= 10
				}
			case *sdl.QuitEvent:
				quit = true
			}
		}

		renderer.Clear()
		rect := sdl.Rect{x, y, 200, 200}
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.DrawRect(&rect)
		renderer.SetDrawColor(0, 0, 0, 255)

		renderer.Present()
	}
}
