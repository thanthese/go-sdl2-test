package main

// play with fonts: http://twinklebear.github.io/sdl2%20tutorials/2013/12/18/lesson-6-true-type-fonts-with-sdl_ttf/
// more with fonts: http://lazyfoo.net/SDL_tutorials/lesson07/

// further   fonts: http://gameprogrammingtutorials.blogspot.com/2010/02/sdl-tutorial-series-part-6-displaying.html
// crazy more onts: http://content.gpwiki.org/index.php/SDL_ttf:Tutorials:Basic_Font_Rendering
// colors fonts? http://www.gamedev.net/topic/618944-api-for-roguelike/
// sort: http://www.aaroncox.net/tutorials/2dtutorials/sdl_text.pdf

import (
	"fmt"
	"github.com/jackyb/go-sdl2/sdl"
	"github.com/jackyb/go-sdl2/sdl_ttf"
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

	if ttf.Init() != 0 {
		fmt.Fprintf(os.Stderr, "Failed to init ttf")
		os.Exit(1)
	}

	color := sdl.Color{255, 255, 255, 100}
	img := renderText("a", "/Users/thanthese/go/src/github.com/thanthese/sdltest/monaco.ttf",
		color, 32, renderer)
	if img == nil {
		return
	}

	// var h, w int
	// sdl.QueryTexture(img, nil, nil, &w, &h)

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

		renderTexture(img, renderer)

		renderer.SetDrawColor(0, 0, 0, 255)

		renderer.Present()

		sdl.Delay(5)
	}
}

func renderText(msg string, fontFile string, color sdl.Color,
	fontsize int, renderer *sdl.Renderer) *sdl.Texture {

	// open font
	font, e := ttf.OpenFont(fontFile, fontsize)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Failed openfont: %s\n", e)
		os.Exit(1)
	}
	defer font.Close()

	fg := sdl.Color{255, 255, 255, 50}
	bg := sdl.Color{100, 100, 100, 50}
	surf := font.RenderText_Shaded(msg, fg, bg)
	if surf == nil {
		fmt.Fprintf(os.Stderr, "Failed making a surface: %s\n", e)
		os.Exit(1)
	}
	defer surf.Free()

	texture := renderer.CreateTextureFromSurface(surf)
	if texture == nil {
		fmt.Fprintf(os.Stderr, "Failed making a texture: %s\n", e)
		os.Exit(1)
	}

	return texture
}

func renderTexture(tex *sdl.Texture, ren *sdl.Renderer) {
	rect := sdl.Rect{X: 50, Y: 50, W: 15, H: 20}

	ren.Copy(tex, nil, &rect)
}
