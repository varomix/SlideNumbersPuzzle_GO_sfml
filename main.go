package main

import (
	"runtime"
	"time"

	sf "github.com/manyminds/gosfml"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	ticker := time.NewTicker(time.Second / 60)

	var (
		gameWidth  uint = 544
		gameHeight uint = 650
	)

	// main window
	renderWindow := sf.NewRenderWindow(sf.VideoMode{gameWidth, gameHeight, 32}, "Slide Numbers Puzzle", sf.StyleClose, sf.DefaultContextSettings())

	// bg texture
	bgTex, _ := sf.NewTextureFromFile("assets/images/bg.png", nil)
	bg, _ := sf.NewSprite(bgTex)

	// Load font
	font, _ := sf.NewFontFromFile("assets/fonts/LuckiestGuy.ttf")

	// Title
	testTxt, _ := sf.NewText(font)
	testTxt.SetString("Test text")
	testTxt.SetCharacterSize(36)

	for renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			// poll events
			for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
				switch ev := event.(type) {
				case sf.EventKeyPressed:
					if ev.Code == sf.KeyEscape {
						renderWindow.Close()
					}
				case sf.EventClosed:
					renderWindow.Close()

				}
			}

		}

		// clear window
		renderWindow.Clear(sf.ColorCyan())

		// Draw here
		renderWindow.Draw(bg, sf.DefaultRenderStates())
		renderWindow.Draw(testTxt, sf.DefaultRenderStates())

		// and display it
		renderWindow.Display()

	}

}
